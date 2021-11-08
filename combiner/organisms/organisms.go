package organisms

import (
	"context"
	"github.com/derbexuk/wurzel/harvester/datawarehouse"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const DB = "poievt"
const COLLECTION = "organisms"

type Organism struct {
	ID          string
	Title       string
	Description *string
	Deactivated *bool
	Properties  map[string]string
	Refs        []string
	Path        *string `json:"path,omitempty" bson:",omitempty"`
}

type WareHouseMessage struct {
	path string
	org  Organism
}

// Creates a new Organism database entry
func Create(org *Organism, path string, dw *datawarehouse.Datawarehouse) error {
	path = fmt.Sprintf("%s/%s", path, org.ID)
	err := dw.Add(DB, COLLECTION, path, org)
	return err
}

func Update(path string, org *Organism, dw *datawarehouse.Datawarehouse) int {
	updateCount := 0
	var old *Organism
	if nil != org.Deactivated {
		res := dw.Get(DB, COLLECTION, path)
		if nil != res {
			old = new(Organism)
			MapToOrganism(res, old, path)
		}
	} else {
		obj := GetByPath(path, dw)
		old = obj
	}
	if nil != old {
		var id, title string
		//Have to overwrite properties in a lump
		if "" == org.Title {
			title = old.Title
		} else {
			title = org.Title
		}
		if nil == org.Description {
			org.Description = old.Description
		}
		if nil == org.Properties {
			org.Properties = old.Properties
		}
		if nil == org.Deactivated {
			org.Deactivated = old.Deactivated
		}
		if nil == org.Refs {
			org.Refs = old.Refs
		}
		if "" == org.ID {
			id = old.ID
		} else {
			id = org.ID
		}
		p := Organism{ID: id, Title: title, Description: org.Description, Refs: org.Refs, Properties: org.Properties, Deactivated: org.Deactivated}

		updateCount = dw.Upsert("poievt", "organisms", path, p)
	}
	return updateCount
}

func Delete(path string, dw *datawarehouse.Datawarehouse) int {
	noRemoved := dw.DelPath(DB, COLLECTION, path)
	return noRemoved

}

func Deactivate(path string, dw *datawarehouse.Datawarehouse) int {
	old := GetByPath(path, dw)
	var deactivateFlag bool
	deactivateFlag = true
	if nil != old {
		old.Deactivated = &deactivateFlag
		updateCount := dw.Upsert("poievt", "organisms", path, old)
		return updateCount
	} else {
		return 0
	}
}

func GetByPath(path string, dw *datawarehouse.Datawarehouse) *Organism {
	var org Organism

	fullpath := fmt.Sprintf("/%s", path)
	res := dw.GetActive(DB, COLLECTION, fullpath)

	if res != nil {
		MapToOrganism(res, &org, fullpath)
		return &org
	} else {
		return nil
	}

}

func ListByPath(path string, dw *datawarehouse.Datawarehouse) (orgs []*Organism){
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
			org := Organism{}
			MapToOrganism(result.Data, &org, result.Path)
			orgs = append(orgs, &org)
		}
	}
	return orgs
}

func ListByRefs(input string, dw *datawarehouse.Datawarehouse) (orgs []*Organism) {
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
			org := Organism{}
			MapToOrganism(result.Data, &org, result.Path)
			orgs = append(orgs, &org)
		}
	}
	return orgs
}

func MapToOrganism(r map[string]interface{}, organism *Organism, path string) {

	id := r["id"].(string)
	organism.ID = id
	title := r["title"].(string)
	organism.Title = title
	if r["description"] != nil {
		description := r["description"].(string)
		organism.Description = &description
	}
	deactivated := r["deactivated"].(bool)
	organism.Deactivated = &deactivated
	organism.Path = &path

	refs := make([]string, 0)
	for _, ref := range r["refs"].(primitive.A) {
		refs = append(refs, ref.(string))
	}
	organism.Refs = refs

	properties := make(map[string]string)
	for k, v := range r["properties"].(map[string]interface{}) {
		properties[k] = v.(string)
	}
	organism.Properties = properties

}
