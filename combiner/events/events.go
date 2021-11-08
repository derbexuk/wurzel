package events

import (
	"context"
	"fmt"
	"log"

	"github.com/derbexuk/wurzel/harvester/datawarehouse"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB = "poievt"
var COLLECTION = "events"

type Event struct {
	ID          string
	Title       string
	Description *string
	Deactivated *bool
	Start       string
	End         *string
	Properties  map[string]string
	Refs        []string
	Path        *string `json:"path,omitempty" bson:",omitempty"`
}

// Create an Event

func Create(path string, event *Event, dw *datawarehouse.Datawarehouse) error {
	fullPath := fmt.Sprintf("/%s/%s", path, event.ID)
	err := dw.Add(DB, COLLECTION, fullPath, event)
	return err
}

// Update an Event
func Update(path string, event *Event, dw *datawarehouse.Datawarehouse) int {
	updateCount := 0
	var oldEvent *Event
	path = "/" + path

	if event.Deactivated != nil {
		res := dw.Get(DB, COLLECTION, path)
		if nil != res {
			oldEvent = new(Event)
			MapToEvent(res, oldEvent, path)
		}
	} else {
		obj := GetByPath(path, dw)
		oldEvent = obj
	}

	if oldEvent != nil {
		var id, title, start string
		if "" == event.Title {
			title = oldEvent.Title
		} else {
			title = event.Title
		}
		if "" == event.ID {
			id = oldEvent.ID
		} else {
			id = event.ID
		}
		if nil == event.Description {
			event.Description = oldEvent.Description
		}
		if nil == event.Deactivated {
			event.Deactivated = oldEvent.Deactivated
		}
		if "" == event.Start {
			start = oldEvent.Start
		} else {
			start = event.Start
		}
		if nil == event.End {
			event.End = oldEvent.End
		}
		if nil == event.Properties {
			event.Properties = oldEvent.Properties
		}
		if nil == event.Refs {
			event.Refs = oldEvent.Refs
		}
		p := Event{ID: id, Title: title, Start: start, End: event.End, Description: event.Description, Refs: event.Refs, Properties: event.Properties, Deactivated: event.Deactivated}
		updateCount = dw.Upsert(DB, COLLECTION, path, p)
	}
	return updateCount
}

// Delete an Event
func Delete(path string, dw *datawarehouse.Datawarehouse) int {
	fullpath := fmt.Sprintf("/%s", path)
	updateCount := dw.DelPath(DB, COLLECTION, fullpath)
	return updateCount
}

// Deactivate an Event
func Deactivate(path string, dw *datawarehouse.Datawarehouse) int {
	fullpath := fmt.Sprintf("/%s", path)
	oldData := dw.Get(DB, COLLECTION, fullpath)
	log.Printf("Finding Event %s\n", path)
	if oldData != nil {
	log.Printf("Deact Event %s\n", path)
		oldEvent := Event{}
		MapToEvent(oldData, &oldEvent, path)
		var deactivateFlag bool
		deactivateFlag = true
		oldEvent.Deactivated = &deactivateFlag
		updateCount := dw.Upsert(DB, COLLECTION, fullpath, oldEvent)
		return updateCount
	} else {
		return 0
	}
}

// Get an Event
func GetByPath(path string, dw *datawarehouse.Datawarehouse) *Event {
	var event Event
	fullpath := fmt.Sprintf("/%s", path)
	result := dw.GetActive(DB, COLLECTION, fullpath)
	if result == nil || len(result) == 0 {
		return nil
	}
	MapToEvent(result, &event, fullpath)
	return &event
}

// Get an Event by path
func ListByPath(path string, dw *datawarehouse.Datawarehouse) (events []*Event, err error) {
	c := dw.Client.Database(DB).Collection(COLLECTION)
	var results []datawarehouse.RetWarehouseObject
	query := path + "*"
	log.Println("Query : " + query)
	filter := bson.M{"path": primitive.Regex{Pattern: query, Options: ""}, "data.deactivated": bson.M{"$ne" : true}}
	cur, err := c.Find(context.Background(), filter)
	if err == nil {
		err = cur.All(context.Background(), &results)
	}

	if err != nil {
		log.Println(err)
		return
	}
	if results != nil {
		if len(results) >= 1 {
			for _, result := range results {
				event := Event{}
				MapToEvent(result.Data, &event, result.Path)
				events = append(events, &event)
			}
		}
	}
	return events, nil
}

// Get Events for a specified path within a specified date-time range
func ListByTimeAndPath(path, start, end string, dw *datawarehouse.Datawarehouse) (events []*Event, err error) {
	// Create DB session
	c := dw.Client.Database(DB).Collection(COLLECTION)
	var results []datawarehouse.RetWarehouseObject
	// Build query
	fullpath := fmt.Sprintf("/%s", path)
	query := bson.M{"path": primitive.Regex{Pattern: fullpath, Options: ""}, "data.deactivated": bson.M{"$ne" : true}}
	if start != "" {
		if end != "" {
			query["data.start"] = bson.M{"$gte": start, "$lte": end}
		} else {
			query["data.start"] = bson.M{"$gte": start}
		}
	}
	// Retrieve data
	cur, err := c.Find(context.Background(), query)
	if err == nil {
		err = cur.All(context.Background(), &results)
	}

	if err != nil {
		log.Println(err)
		return
	}
	// Map data
	for _, result := range results {
		event := Event{}
		MapToEvent(result.Data, &event, result.Path)
		events = append(events, &event)
	}
	return
}

// Get latest Event for a specified path
func TimeSearch(path string, start string, end string, order string, limit int64, dw *datawarehouse.Datawarehouse) (events []*Event, err error) {
	c := dw.Client.Database(DB).Collection(COLLECTION)
	var results []datawarehouse.RetWarehouseObject
	// Build query
	fullpath := fmt.Sprintf("/%s", path)
	query := bson.M{"path": primitive.Regex{Pattern: fullpath, Options: ""}, "data.deactivated": bson.M{"$ne" : true}}
	query["data.start"] = bson.M{"$gte": start, "$lte": end}
	opts := options.Find().SetSort(bson.D{{"data.start", 1}})
	if order == "d" {
		opts.SetSort(bson.D{{"data.start", -1}})
	}
	if limit > 0 {
		opts.SetLimit(limit)
	}
	// Retrieve data
	cur, err := c.Find(context.Background(), query, opts)
	if err == nil {
		err = cur.All(context.Background(), &results)
	}

	if err != nil {
		log.Println(err)
		return
	}
	// Map data
	for _, result := range results {
		event := Event{}
		MapToEvent(result.Data, &event, result.Path)
		events = append(events, &event)
	}
	return
}

// Map hash to CspaceEvent
func MapToEvent(r map[string]interface{}, event *Event, path string) {
	id := r["id"].(string)
	event.ID = id
	title := r["title"].(string)
	event.Title = title
	if r["description"] != nil {
		description := r["description"].(string)
		event.Description = &description
	}
	if r["deactivated"] != nil {
		deactivated := r["deactivated"].(bool)
		event.Deactivated = &deactivated
	}
	start := r["start"].(string)
	event.Start = start
	end := r["end"].(string)
	event.End = &end
	event.Path = &path
	properties := make(map[string]string)
	for k, v := range r["properties"].(map[string]interface{}) {
		properties[k] = v.(string)
	}
	event.Properties = properties
	refs := make([]string, 0)
	for _, ref := range r["refs"].(primitive.A) {
		refs = append(refs, ref.(string))
	}
	event.Refs = refs
}

/*
func MapToEventSummary(r map[string]interface{}, event *EventSummary, path string) {
	id := r["id"].(string)
	event.ID = &id
	title := r["title"].(string)
	event.Title = &title
	start := r["start"].(string)
	event.Start = &start
	end := r["end"].(string)
	event.End = &end
	event.Path = &path
}
*/
