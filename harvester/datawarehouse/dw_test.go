package datawarehouse

import (
	"log"
	"fmt"
	"context"
	"os"
	"testing"

	. "github.com/onsi/gomega"
  "go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/bson/primitive"

)

type MyType struct{ S string }

//Name and case is important, must match those of WarehouseObject
type MyType2 struct {
	Path string
	Data MyType
}

type ParentType struct {
	Parent  string
	Message string
}

var dbh Datawarehouse

const DB = "dwTest"
const COL = "dwTest"

func TestMain(m *testing.M) {
	dbh = Datawarehouse{}
	dbh.Open()

	defer dbh.Close()
	code := m.Run()
	db := dbh.Client.Database(DB)
	db.Drop(context.Background())
	os.Exit(code)
}

func TestAdd(t *testing.T) {
	RegisterTestingT(t)

	obj := MyType{"abc"}
	_ = dbh.Add(DB, COL, "garden/addpath", obj)

	result := MyType2{}

  coll := dbh.Client.Database(DB).Collection(COL)

  res := coll.FindOne(context.Background(), bson.D{{"path", "garden/addpath"}})
	Expect(res).ShouldNot(BeNil())
  err := res.Decode(&result)
	Expect(err).Should(BeNil())

  log.Println(result)

	Expect(result.Data.S).To(Equal("abc"))
}

func TestUpsert(t *testing.T) {
	RegisterTestingT(t)
  coll := dbh.Client.Database(DB).Collection(COL)

	obj := MyType{"abc"}
	_ = dbh.Add(DB, COL, "garden/upsert", obj)

	obj2 := MyType{"xyz"}
	dbh.Upsert(DB, COL, "garden/upsert", obj2)
	result := MyType2{}
  res := coll.FindOne(context.Background(), bson.D{{"path", "garden/upsert"}})
  err := res.Decode(&result)
	Expect(err).Should(BeNil())
	Expect(result.Data.S).To(Equal("xyz"))

  cur, err := coll.Find(context.Background(), bson.D{{"path", "garden/upsert"}})
	Expect(cur.Next(context.TODO())).To(Equal(true))
	Expect(cur.Next(context.TODO())).To(Equal(false))
}

func TestDelPath(t *testing.T) {
	RegisterTestingT(t)
  coll := dbh.Client.Database(DB).Collection(COL)

	obj := MyType{"abc"}
	_ = dbh.Add(DB, COL, "garden/delete/3", obj)
	_ = dbh.Add(DB, COL, "garden/delete/4", obj)
	_ = dbh.Add(DB, COL, "garden/delete/5", obj)
	_ = dbh.Add(DB, COL, "garden/delete/6", obj)
  n, err := coll.CountDocuments(context.Background(), bson.D{{"path", primitive.Regex{Pattern: "garden/delete", Options: ""}}})
	Expect(err).Should(BeNil())
	Expect(n).To(Equal(int64(4)))

	result1 := dbh.DelPath(DB, COL, "garden/delete/6")
	Expect(err).Should(BeNil())
	Expect(result1).To(Equal(1))

  n, err = coll.CountDocuments(context.Background(), bson.D{{"path", primitive.Regex{Pattern: "garden/delete", Options: ""}}})
	Expect(err).Should(BeNil())
	Expect(n).To(Equal(int64(3)))

	result2 := dbh.DelPath(DB, COL, "garden/delete/")
	Expect(err).Should(BeNil())
	Expect(result2).To(Equal(3))
}

func TestAddLots(t *testing.T) {
	RegisterTestingT(t)
  coll := dbh.Client.Database(DB).Collection(COL)

	pathRoot := "/stony/path/"
	obj := MyType{"abc"}
	var wObjs []interface{}

	for i := 0; i < 100; i++ {
		path := fmt.Sprintf("%s%d", pathRoot, i)
		wO := MyType2{Path: path, Data: obj}
		wObjs = append(wObjs, wO)
	}

	err := dbh.AddLots(DB, COL, wObjs)
  n, err := coll.CountDocuments(context.Background(), bson.D{{"path", primitive.Regex{Pattern: "/stony/path/", Options: ""}}})
	Expect(err).Should(BeNil())
	Expect(n).To(Equal(int64(100)))

	result := MyType2{}
  res := coll.FindOne(context.Background(), bson.D{{"path", "/stony/path/66"}})
  err = res.Decode(&result)
	Expect(err).Should(BeNil())
	Expect(result.Path).To(Equal("/stony/path/66"))
	Expect(result.Data.S).To(Equal("abc"))

}

func TestGet(t *testing.T) {
	RegisterTestingT(t)
	// add an object
	obj := MyType{"01"}
	_ = dbh.Add(DB, COL, "garden/get/item01", obj)
	// retrieve the object
	result := dbh.Get(DB, COL, "garden/get/item01")
	var value = result["s"].(string)
	Expect(value).To(Equal("01"))
}

// Test GetAll function
func TestGetAll(t *testing.T) {
	RegisterTestingT(t)
	// add some objects
	obj1 := ParentType{"001", "Parent 001 message 1"}
	obj2 := ParentType{"002", "Parent 002 message 1"}
	obj3 := ParentType{"001", "Parent 001 message 2"}
	obj4 := ParentType{"002", "Parent 002 message 2"}
	_ = dbh.Add(DB, COL, "garden/getallpath/parent", obj1)
	_ = dbh.Add(DB, COL, "garden/getallpath/parent", obj2)
	_ = dbh.Add(DB, COL, "garden/getallpath/parent", obj3)
	_ = dbh.Add(DB, COL, "garden/getallpath/parent", obj4)
	// Retrieve the objects - should have 2 of them
	result := dbh.GetAll(DB, COL, "001")
	Expect(len(result)).To(Equal(2))
	// Get the contents of the 1st object
	data := result[0].Data
	// Check the parent data - should be "001"
	parent := data["parent"].(string)
	Expect(parent).To(Equal("001"))
	// Check the message data - should be "Parent 001 message 1"
	message := data["message"].(string)
	Expect(message).To(Equal("Parent 001 message 1"))
}

// Test GetLike function
func TestGetLike(t *testing.T) {
	RegisterTestingT(t)
	// add some objects
	obj1 := ParentType{"001", "Parent 001 message 1"}
	obj2 := ParentType{"002", "Parent 002 message 1"}
	_ = dbh.Add(DB, COL, "farm/drive/parent/001", obj1)
	_ = dbh.Add(DB, COL, "farm/drive/parent/002", obj2)
	// Retrieve the objects - should have 2 of them
	results := dbh.GetLike(DB, COL, "farm/drive")
	Expect(len(results)).To(Equal(2))
	// Get the contents of the 1st object
	data := results[0].Data
	// Check the parent data - should be "001"
	parent := data["parent"].(string)
	Expect(parent).To(Equal("001"))
	// Check the message data - should be "Parent 001 message 2"
	message := data["message"].(string)
	Expect(message).To(Equal("Parent 001 message 1"))
	// Get the contents of the 2nd object
	data = results[1].Data
	// Check the parent data - should be "002"
	parent = data["parent"].(string)
	Expect(parent).To(Equal("002"))
	// Check the message data - should be "Parent 002 message 2"
	message = data["message"].(string)
	Expect(message).To(Equal("Parent 002 message 1"))
}

// Test Add - attempt to add a duplicate record
func TestAddDuplicate(t *testing.T) {
	RegisterTestingT(t)
	// Needs to be indexed to catch duplicates hence different collection
	collection := "dups"
	dbh.EnsurePathIndex(DB, collection)
	// Object to insert
	obj := MyType{"abc"}
	// Insert first object
	err := dbh.Add(DB, collection, "duplicate/object", obj)
	Expect(err).Should(BeNil())
	// Insert second (duplicate) object
	err = dbh.Add(DB, collection, "duplicate/object", obj)
	Expect(err).ShouldNot(BeNil())
	Expect(err.Error()).Should(ContainSubstring("duplicate key error"))
}
