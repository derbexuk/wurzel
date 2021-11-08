package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strings"
	"time"

	"gopkg.in/yaml.v2"

	ds "github.com/derbexuk/wurzel/harvester/datastore"
	"github.com/derbexuk/wurzel/harvester/datawarehouse"

	"github.com/derbexuk/wurzel/harvester/feeds/fetchers"
	"github.com/derbexuk/wurzel/combiner/events"
	"github.com/derbexuk/wurzel/combiner/organisms"
	"github.com/derbexuk/wurzel/combiner/pois"
)

//TODO Read from env
var DB = "poievt"

type ConsumerConfigs struct {
	Sources map[string]map[string]string
	//Sources map[string]ConsumerConfig
}

func (ccs *ConsumerConfigs) Read(config_file string) error {
	config, err := ioutil.ReadFile(config_file)
	if err != nil {
		log.Printf("Failed to read config file : error: %v", err)
		return err
	}
	err = yaml.Unmarshal([]byte(config), ccs)
	if err != nil {
		log.Printf("Failed to load config file : error: %v", err)
		return err
	}
	return nil
}

func (ccs *ConsumerConfigs) Consume() {
	ds.Init()

	updatePubSub := ds.SubscribeUpdates()
	log.Println(updatePubSub)
	for {
		msg, err := updatePubSub.ReceiveMessage()
		if err != nil {
			panic(err)
		}

		log.Println(msg.Channel, msg.Payload)
		_, ok := ccs.Sources[msg.Payload]
		if ok {
			ccs.Process(msg.Payload)
		}
	}
}

func (ccs *ConsumerConfigs) Process(path string) error {
	dJson, err := ds.Get(path)
	if err != nil {
		log.Println(err)
		return err
	}
	incoming := fetchers.GenericResult{}
	err = json.Unmarshal([]byte(dJson), &incoming)
	if err != nil {
		log.Println(err)
		return err
	}

	dw := datawarehouse.Datawarehouse{}
	dw.Open()
	defer dw.Close()

	if len(incoming.Events) > 0 {
		err = ccs.ProcessEvents(path, incoming.Events, &dw)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	if len(incoming.Pois) > 0 {
		err = ccs.ProcessPois(path, incoming.Pois, &dw)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	if len(incoming.Organisms) > 0 {
		err = ccs.ProcessOrganisms(path, incoming.Organisms, &dw)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}

func (ccs *ConsumerConfigs) ProcessPois(path string, pois []*pois.Poi, dw *datawarehouse.Datawarehouse) error {
	log.Println(path)
	var wos []interface{}
	var err error
	for i, poi := range pois {
		wo := datawarehouse.WarehouseObject{}
		//Create path from config
		wo.Path, err = makePath(ccs.Sources[path]["poiDestPath"], poi)
		if err != nil {
			log.Println(err)
			continue
		}
		refPath, ok := ccs.Sources[path]["poiRefPath"]
		if ok {
			newRefs := []string{}
			for _, ref := range poi.Refs {
				newRefs = append(newRefs, refPath+"/"+ref)
			}
			poi.Refs = newRefs
		}
		wo.Data = pois[i]
		wos = append(wos, wo)
	}

	dropData, ok := ccs.Sources[path]["dropData"]
	if ok && "true" == dropData {
		nos := dw.DelPath(DB, "pois", ccs.Sources[path]["poiDestPath"]+"*")
		log.Println("Deleted poi ", nos)
	}

	err = dw.AddLots(DB, "pois", wos)
	return err
}

/*
allow substitutions from object fiels into the path
e.g.
  /pois/businesses/=Properties.Category
would subsitute in the value for Properties.Category in the passed in object
so if Properties.Category was 'Financial Services' the path for object id 984
would be
  /pois/businesses/Financial Services/984
*/
func makePath(path string, obj interface{}) (string, error) {
	//If there is an '=' in the field then we need to perform substitution
	pAry := strings.SplitN(path, "=", 2)
	p := reflect.ValueOf(obj)
	if p == reflect.ValueOf(nil) || reflect.Indirect(p) == reflect.ValueOf(nil) {
		return "", errors.New("empty object")
	}
	log.Println(p)
	idField := reflect.Indirect(p).FieldByName("ID")
	if idField == reflect.ValueOf(nil) {
		return "", errors.New("Missing ID field")
	}
	id := idField.String()
	if id == "" {
		return "", errors.New("Missing ID field")
	}
	//No substitution
	if len(pAry) == 1 {
		return path + "/" + id, nil
	}
	sAry := strings.SplitN(pAry[1], ".", 2)
	//substitution from main attribute
	if len(sAry) == 1 {
		f := reflect.Indirect(p).FieldByName(sAry[0])
		return pAry[0] + f.String() + "/" + id, nil
	}

	//substitution from Properties attribute
	f := reflect.Indirect(p).FieldByName("Properties")
	if f == reflect.ValueOf(nil) {
		return "", errors.New("Missing Properties field")
	}
	attr := f.MapIndex(reflect.ValueOf(sAry[1]))
	return pAry[0] + attr.String() + "/" + id, nil
}

func (ccs *ConsumerConfigs) ProcessEvents(path string, events []*events.Event, dw *datawarehouse.Datawarehouse) error {
	var wos []interface{}
	var err error
	for i, event := range events {
		wo := datawarehouse.WarehouseObject{}
		//Create path from config
		wo.Path, err = makePath(ccs.Sources[path]["eventDestPath"], event)
		if err != nil {
			log.Println(err)
			continue
		}
		refPath, ok := ccs.Sources[path]["eventRefPath"]
		if ok {
			newRefs := []string{}
			for _, ref := range event.Refs {
				newRefs = append(newRefs, refPath+"/"+ref)
			}
			event.Refs = newRefs
		}
		wo.Data = events[i]
		wos = append(wos, wo)
	}

	dropData, ok := ccs.Sources[path]["dropData"]
	if ok && "true" == dropData {
		nos := dw.DelPath(DB, "events", ccs.Sources[path]["eventDestPath"]+"*")
		log.Println("Deleted events ", nos)
	}
	err = dw.AddLots(DB, "events", wos)
	return err
}

func (ccs *ConsumerConfigs) ProcessOrganisms(path string, organisms []*organisms.Organism, dw *datawarehouse.Datawarehouse) error {
	var err error
	var wos []interface{}
	for i, organism := range organisms {
		wo := datawarehouse.WarehouseObject{}
		//Create path from config
		wo.Path, err = makePath(ccs.Sources[path]["organismDestPath"], organism)
		if err != nil {
			log.Println(err)
			continue
		}
		refPath, ok := ccs.Sources[path]["organismRefPath"]
		if ok {
			newRefs := []string{}
			for _, ref := range organism.Refs {
				newRefs = append(newRefs, refPath+"/"+ref)
			}
			organism.Refs = newRefs
		}
		wo.Data = organisms[i]
		wos = append(wos, wo)
	}

	dropData, ok := ccs.Sources[path]["dropData"]
	if ok && "true" == dropData {
		nos := dw.DelPath(DB, "organisms", ccs.Sources[path]["organismDestPath"]+"*")
		log.Println("Deleted orgs ", nos)
	}
	err = dw.AddLots(DB, "organisms", wos)
	return err
}

func main() {

	if value, ok := os.LookupEnv("COLL_DB"); ok {
		DB = value
	}
	var config_file string
	if value, ok := os.LookupEnv("COLL_CONFIG_FILE"); ok {
		config_file = value
	} else {
		log.Fatal("No config file Env Var")
	}

	sleepTime := 1
	for i := 0; i < 10; i++ {
		if _, err := os.Stat(config_file); err == nil {
			break
		}
		time.Sleep(time.Duration(sleepTime) * time.Second)
		sleepTime = sleepTime * i
	}

	ccs := ConsumerConfigs{}
	err := ccs.Read(config_file)
	if err != nil {
		log.Fatal(err)
	}
	ccs.Consume()

}
