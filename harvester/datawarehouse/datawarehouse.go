package datawarehouse

import (
	"context"
	//"errors"

	"log"
	"os"
	"strings"
	"time"

  "go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/bson/primitive"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
)

//Holds a Mongo Session
type Datawarehouse struct {
	Client *mongo.Client
}

//Open a Mongo Session
func (dw *Datawarehouse) Open() {
	var err error

  mongo_uri := os.Getenv("MONGOHOST")
  if mongo_uri == "" {
    mongo_uri = "mongodb://localhost:27017"
  }
  log.Printf("Opening connection to : %s\n", mongo_uri)
  dw.Client, err = mongo.NewClient(options.Client().ApplyURI(mongo_uri))
  if err != nil {
    panic(err)
  }
  ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
  err = dw.Client.Connect(ctx)
  if err != nil {
    panic(err)
  }
}

//Close a Mongo Session
func (dw *Datawarehouse) Close() {
	dw.Client.Disconnect(context.Background())
}

//Create an index on the path, if one doesn't exist
func (dw *Datawarehouse) EnsurePathIndex(db string, coll string) {
  c := dw.Client.Database(db).Collection(coll)
	mod := mongo.IndexModel{
		Keys: bson.M{
				"path": 1,
		},
		// create UniqueIndex option 
		Options: options.Index().SetUnique(true),
	}
	_, err := c.Indexes().CreateOne(context.Background(), mod)
  if err != nil {
    panic(err)
  }
}

/*
//Create a GeoSpatial index there must be a geo field of the approved type in data
func (dw *Datawarehouse) EnsureGeoIndex(db string, coll string) {
	session := dw.Session.Clone()
	defer session.Close()

	c := session.DB(db).C(coll)

	index := mgo.Index{
		Key:        []string{"$2dsphere:data.geo"},
		Background: true,
	}
	err := c.EnsureIndex(index)
	if err != nil {
		panic(err)
	}
}

/*
	Create an index on parent, if one doesn't exist
* /
func (dw *Datawarehouse) EnsureParentIndex(db string, coll string) {
	session := dw.Session.Clone()
	defer session.Close()
	c := session.DB(db).C(coll)
	index := mgo.Index{
		Key: []string{"parent"},
	}
	err := c.EnsureIndex(index)
	if err != nil {
		panic(err)
	}
}
*/

//Generic type to add to Mongo
//When retrieving from the database the target type must have path and data e.g.
// FindType struct{ Path string; Data MyType}
type WarehouseObject struct {
	Path string
	Data interface{}
}

//Add a record, cloning the Session should make use of existing or pooled connections
func (dw *Datawarehouse) Add(db string, coll string, path string, obj interface{}) error {
  c := dw.Client.Database(db).Collection(coll)

	wObj := WarehouseObject{path, obj}
  _, err := c.InsertOne(context.Background(), wObj)

	// Check err - panic if not nil and not duplicate
	if err != nil && !strings.Contains(err.Error(), "duplicate key error") {
		log.Panic(err)
	}
	return err
}

func (dw *Datawarehouse) AddLots(db string, coll string, objs []interface{}) error {
  c := dw.Client.Database(db).Collection(coll)

  _, err := c.InsertMany(context.Background(), objs)

	return err
}

//Upsert a record, should only be used for unique paths whose data may chnage
//It would probably be more valid to have a set of records with a 'latest' pointer
func (dw *Datawarehouse) Upsert(db string, coll string, path string, obj interface{}) int {
  c := dw.Client.Database(db).Collection(coll)
	wObj := WarehouseObject{path, obj}

	opts := options.Replace()
	up := true
	opts.Upsert = &up

	changeinfo, err := c.ReplaceOne(context.Background(), bson.M{"path": path}, &wObj, opts)
	if err != nil {
		log.Panic(err)
	}
	return int(changeinfo.UpsertedCount + changeinfo.ModifiedCount)
}

//Delete documents based on path
func (dw *Datawarehouse) DelPath(db string, coll string, path string) int {
  c := dw.Client.Database(db).Collection(coll)
	filter :=bson.D{{"path", primitive.Regex{Pattern: path, Options: ""}}}

	changeInfo, err := c.DeleteMany(context.Background(), filter)

	if err != nil {
		log.Panic(err)
	}
	return int(changeInfo.DeletedCount)
}

//Want to return into a hash
type RetWarehouseObject struct {
	Path string
	Data map[string]interface{}
}

//I would like to be able to get a generic type from the db and cast it when I get back to
//the calling context, not sure it's possible.
func (dw *Datawarehouse) Get(db string, coll string, path string) map[string]interface{} {
  c := dw.Client.Database(db).Collection(coll)

	result := RetWarehouseObject{} //"", data}
	res := c.FindOne(context.Background(), bson.D{{"path",path}})
	if res == nil {
		return nil
	}
	err := res.Decode(&result)
	if err != nil {
		log.Print(err)
		return nil
	}

	return result.Data
}

//Fetch only active record from database
func (dw *Datawarehouse) GetActive(db string, coll string, path string) map[string]interface{} {
  c := dw.Client.Database(db).Collection(coll)

	result := RetWarehouseObject{} //"", data}

	res := c.FindOne(context.Background(), bson.M{"path": path, "data.deactivated": bson.M{"$ne" : true}})
	if res == nil {
		return nil
	}
	err := res.Decode(&result)
	if err != nil {
		log.Print(err)
		return nil
	}

	return result.Data
}

/*
	Generic retrieve for a specific path
*/
func (dw *Datawarehouse) GetAll(db string, coll string, parent string) []RetWarehouseObject {
  c := dw.Client.Database(db).Collection(coll)
	// Get results for specified parent
	results := []RetWarehouseObject{} //"", data}
	cur, err := c.Find(context.Background(), bson.M{"data.parent": parent})
	if err != nil {
		log.Print(err)
		return nil
	}
	err = cur.All(context.Background(), &results)
	if err != nil {
		log.Print(err)
		return nil
	}
	// Return results
	return results
}

// Returns all results that have a path like that provided
func (dw *Datawarehouse) GetLike(db string, coll string, path string, activeOnly ...string) []RetWarehouseObject {
  c := dw.Client.Database(db).Collection(coll)

	results := []RetWarehouseObject{}
	query := path + "*"
	filter :=bson.M{"path": primitive.Regex{Pattern: query, Options: ""}}
	if len(activeOnly) > 0 {
		filter =bson.M{"path": primitive.Regex{Pattern: query, Options: ""}, "data.deactivated": false}
	}

	cur, err := c.Find(context.Background(), filter)
	if err != nil {
		log.Print(err)
		return nil
	}
	err = cur.All(context.Background(), &results)
	if err != nil {
		log.Print(err)
		return nil
	}
	return results
}
