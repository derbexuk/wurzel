package pois

import (
	"context"
	"fmt"
	"github.com/derbexuk/wurzel/harvester/datawarehouse"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	//"go.mongodb.org/mongo-driver/mongo/options"
)

var DB = "poievt"
var COLLECTION = "pois"

type Poi struct {
	ID          string
	Title       string
	Description *string
	Deactivated *bool
	Location    *string
	GeoJSON     string         `json:geojson"`
	Geotype     string         `json:",omitempty" bson:",omitempty"`
	Geolocation []*Geolocation `json:",omitempty" bson:",omitempty"`
	Properties  map[string]string
	Refs        []string
	Path        *string `json:"path,omitempty" bson:",omitempty"`
}

// geolocation user type.
type Geolocation struct {
	Latitude  *float64
	Longitude *float64
}

func Create(poi *Poi, path string, dw *datawarehouse.Datawarehouse) error {

	geos := make([]*Geolocation, 0)
	for _, geo := range poi.Geolocation {
		geos = append(geos, geo)
	}

	path = fmt.Sprintf("%s/%s", path, poi.ID)
	err := dw.Add(DB, COLLECTION, path, poi)
	return err
}

func Update(path string, poi *Poi, dw *datawarehouse.Datawarehouse) int {
	updateCount := 0
	var old *Poi
	if nil != poi.Deactivated {
		res := dw.Get(DB, COLLECTION, path)
		if nil != res {
			old = new(Poi)
			MapToPoi(res, old, path)
		}
	} else {
		obj := GetByPath(path, dw)
		old = obj
	}
	if nil != old {
		//Have to overwrite properties in a lump
		geos := make([]*Geolocation, 0)
		if "" == poi.Title {
			poi.Title = old.Title
		}
		if "" == poi.ID {
			poi.ID = old.ID
		}
		if nil == poi.Geolocation {
			poi.Geolocation = old.Geolocation
		} else {
			for _, geo := range poi.Geolocation {
				geos = append(geos, geo)
			}
		}

		if nil == poi.Location {
			poi.Location = old.Location
		}
		if nil == poi.Properties {
			poi.Properties = old.Properties
		}
		if "" == poi.GeoJSON {
			poi.GeoJSON = old.GeoJSON
		}
		if "" == poi.Geotype {
			poi.Geotype = old.Geotype
		}
		if nil == poi.Description {
			poi.Description = old.Description
		}
		if nil == poi.Deactivated {
			poi.Deactivated = old.Deactivated
		}
		if nil == poi.Refs {
			poi.Refs = old.Refs
		}
		updateCount = dw.Upsert("poievt", "pois", path, poi)
	}
	return updateCount
}

func Deactivate(path string, dw *datawarehouse.Datawarehouse) int {
	old := GetByPath(path, dw)
	var deactivateFlag bool
	deactivateFlag = true
	if nil != old {
		old.Deactivated = &deactivateFlag
		updateCount := dw.Upsert("poievt", "pois", "/" + path, old)
		return updateCount
	} else {
		return 0
	}
}

func Delete(path string, dw *datawarehouse.Datawarehouse) int {
	noRemoved := dw.DelPath(DB, COLLECTION, path)
	return noRemoved

}

func GetByPath(path string, dw *datawarehouse.Datawarehouse) *Poi {
	var poi Poi

	fullpath := fmt.Sprintf("/%s", path)
	res := dw.GetActive(DB, COLLECTION, fullpath)

	if res != nil {
		MapToPoi(res, &poi, fullpath)
		return &poi
	} else {
		return nil
	}
}

func ListByPath(path string, dw *datawarehouse.Datawarehouse) (pois []*Poi) {
	var results []datawarehouse.RetWarehouseObject

	c := dw.Client.Database(DB).Collection(COLLECTION)
	query := path + "*"
	log.Println("Query : " + query)
	filter := bson.M{"path": primitive.Regex{Pattern: query, Options: ""}, "data.deactivated": false}
	cur, err := c.Find(context.Background(), filter)
	if err == nil {
		err = cur.All(context.Background(), &results)
	}

	if err != nil {
		log.Println(err)
		return
	}

	if len(results) >= 1 {
		for _, result := range results {
			poi := Poi{}
			println("path****", result.Path)
			MapToPoi(result.Data, &poi, result.Path)
			pois = append(pois, &poi)
		}
	}
	return pois
}

func ListByRefs(input string, dw *datawarehouse.Datawarehouse) (pois []*Poi) {
	var results []datawarehouse.RetWarehouseObject

	c := dw.Client.Database(DB).Collection(COLLECTION)
	filter := bson.M{"data.refs": primitive.Regex{Pattern: input, Options: ""}, "data.deactivated": false}
	cur, err := c.Find(context.Background(), filter)
	if err == nil {
		err = cur.All(context.Background(), &results)
	}

	if err != nil {
		log.Println(err)
		return
	}

	if len(results) >= 1 {
		for _, result := range results {
			poi := Poi{}
			MapToPoi(result.Data, &poi, result.Path)
			pois = append(pois, &poi)
		}
	}
	return pois
}

func MapToPoi(r map[string]interface{}, poi *Poi, path string) {

	id := r["id"].(string)
	poi.ID = id
	title := r["title"].(string)
	poi.Title = title
	if r["geojson"] != nil {
		geojson := r["geojson"].(string)
		poi.GeoJSON = geojson
	}
	geotype := r["geotype"].(string)
	poi.Geotype = geotype
	if r["location"] != nil {
		location := r["location"].(string)
		poi.Location = &location
	}
	if r["description"] != nil {
		description := r["description"].(string)
		poi.Description = &description
	}
	if r["deactivated"] != nil {
		deactivated := r["deactivated"].(bool)
		poi.Deactivated = &deactivated
	}
	if r["path"] != nil {
		path := r["path"].(string)
		poi.Path = &path
	}

	refs := make([]string, 0)
	for _, ref := range r["refs"].(primitive.A) {
		refs = append(refs, ref.(string))
	}
	poi.Refs = refs

	properties := make(map[string]string)
	for k, v := range r["properties"].(map[string]interface{}) {
		properties[k] = v.(string)
	}
	poi.Properties = properties

	if r["geolocation"] != nil {
		geos := make([]*Geolocation, 0)
		for _, geo := range r["geolocation"].(primitive.A) {
			geoObj := geo.(map[string]interface{})
			lat := geoObj["latitude"].(float64)
			lon := geoObj["longitude"].(float64)
			geos = append(geos, &Geolocation{&lat, &lon})
		}
		poi.Geolocation = geos
	}

}
