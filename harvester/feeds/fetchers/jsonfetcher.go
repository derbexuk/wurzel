package fetchers

import (
	"bytes"
	"github.com/derbexuk/wurzel/harvester/feeds/config"
	"github.com/derbexuk/wurzel/combiner/events"
	"github.com/derbexuk/wurzel/combiner/organisms"
	"github.com/derbexuk/wurzel/combiner/pois"
	//"bitbucket.org/skunk/wurzel/combiner/server/app"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

//JSON Fetcher a REST JSON feed
type JSONFetcher struct {
	Config config.FeedConfig
	rx     RESTFetcher
}

type References struct {
	Poi      string
	Event    string
	Organism string
}

func (fx *JSONFetcher) Parse(contents []byte) map[string]string {
	return fx.process([]byte(contents), nil)
}

func (fx *JSONFetcher) process(contents []byte, extras map[string]string) map[string]string {
	results := make(map[string]string)

	path := fx.Config.Paths[0]
	destination := path["destination"]
	sourcePath := strings.Split(destination, "|")
        if len(sourcePath) == 3 {
		ss, ok := extras[sourcePath[1]]
		if ok {
        		destination = fmt.Sprintf("%s%s%s", sourcePath[0], ss, sourcePath[2])
        	}
	}
	log.Println("destination : ", destination)

	//Is there a Generic definition in the config file
	gD, ok := fx.Config.Params["generic"]
	if ok {
		var data map[string]interface{}
		var dataAry []map[string]interface{}

		//Array of objects
		if bytes.HasPrefix(bytes.TrimSpace(contents), []byte("[")) {
			if err := json.Unmarshal(contents, &dataAry); err != nil {
				log.Printf(">>>%s", string(contents))
				log.Panic(err)
			}
		} else {
			if err := json.Unmarshal(contents, &data); err != nil {
				log.Printf(">>>%s", string(contents))
				log.Panic(err)
			}
			dataAry = append(dataAry, data)
		}

		results[destination] = makeGenerics(gD.(map[interface{}]interface{}), extras, dataAry)
	} else { /* Return raw json */
		results[destination] = string(contents)
	}
	log.Println("PROCESSED")
	return results
}

func makeGenerics(genericDef map[interface{}]interface{}, extras map[string]string, dataAry []map[string]interface{}) string {

	gResult := GenericResult{}
	rf := References{}

	var order []string
	od, ok := genericDef["order"]
	if ok {
		for _, e := range od.([]interface{}) {
			order = append(order, e.(string))
		}
	} else {
		//Default order of entity processing
		order = []string{"event", "organism", "poi"}
	}
	for i := 0; i < len(dataAry); i++ {
		for _, entityType := range order {
			if entityType == "poi" {
				//Doesn't have to be a defintion for any entity type
				_, ok := genericDef["poi"]
				if ok {
					gResult.Pois = append(gResult.Pois, rf.makePoi(genericDef["poi"].(map[interface{}]interface{}), dataAry[i])...)
				}
			} else if entityType == "event" {
				_, ok := genericDef["event"]
				if ok {
					gResult.Events = append(gResult.Events, rf.makeEvent(genericDef["event"].(map[interface{}]interface{}), extras, dataAry[i])...)
				}
			} else if entityType == "organism" {
				_, ok := genericDef["organism"]
				if ok {
					gResult.Organisms = append(gResult.Organisms, rf.makeOrganism(genericDef["organism"].(map[interface{}]interface{}), dataAry[i])...)
				}
			}
		}
	}
	jstr, err := json.Marshal(gResult)
	if err != nil {
		log.Printf("%v\n", gResult)
		log.Panic(err)
	}
	return string(jstr)
}

func (rf *References) makeEvent(eventDef map[interface{}]interface{}, extras map[string]string, data map[string]interface{}) (evs []*events.Event) {
	var err error

	eventIt := FeedIterator{}
	eventIt.New(eventDef, data)
	if eventIt.Length == 0 {
		evs = make([]*events.Event, 1)
	} else {
		evs = make([]*events.Event, eventIt.Length)
	}

	for {
		event := events.Event{}
		event.Title, err = eventIt.GetTitle(eventDef["Title"].(string), extras, data)
		if err != nil {
			log.Panic(err)
		}

		evtEnd := ""
		event.Start, evtEnd = eventIt.GetTimes(eventDef, data)
		if evtEnd != "" {
			event.End = &evtEnd
		}

		//Optional
		event.Refs = eventIt.GetRefs(rf, eventDef, data)

		//Optional
		_, ok := eventDef["Description"]
		if ok {
			desc, _ := eventIt.GetItVal(eventDef["Description"].(string), data)
			event.Description = &desc
		}
		event.ID = eventIt.GetId(eventDef, data)
		rf.Event = event.ID

		_, ok = eventDef["Properties"]
		if ok {
			event.Properties = eventIt.GetProps(eventDef["Properties"].(map[interface{}]interface{}), data)
		}
		evs[eventIt.Iteration] = &event
		if !eventIt.Next() {
			break
		}
	}
	return evs
}

func (rf *References) makeOrganism(organismDef map[interface{}]interface{}, data map[string]interface{}) (orgs []*organisms.Organism) {
	var err error

	organismIt := FeedIterator{}
	organismIt.New(organismDef, data)
	if organismIt.Length == 0 {
		orgs = make([]*organisms.Organism, 1)
	} else {
		orgs = make([]*organisms.Organism, organismIt.Length)
	}

	for {
		organism := organisms.Organism{}
		organism.Title, err = organismIt.GetItVal(organismDef["Title"].(string), data)
		if err != nil {
			log.Panic(err)
		}

		//Optional
		organism.Refs = organismIt.GetRefs(rf, organismDef, data)

		//Optional
		_, ok := organismDef["Description"]
		if ok {
			d, _ := organismIt.GetItVal(organismDef["Description"].(string), data)
			organism.Description = &d
		}
		organism.ID = organismIt.GetId(organismDef, data)
		rf.Organism = organism.ID

		_, ok = organismDef["Properties"]
		if ok {
			organism.Properties = organismIt.GetProps(organismDef["Properties"].(map[interface{}]interface{}), data)
		}
		orgs[organismIt.Iteration] = &organism
		if !organismIt.Next() {
			break
		}
	}
	return orgs
}

func (rf *References) makePoi(poiDef map[interface{}]interface{}, data map[string]interface{}) (poiAry []*pois.Poi) {
	var err error

	poiIt := FeedIterator{}
	poiIt.New(poiDef, data)
	if poiIt.Length == 0 {
		poiAry = make([]*pois.Poi, 1)
	} else {
		poiAry = make([]*pois.Poi, poiIt.Length)
	}

	for {
		poi := pois.Poi{}
		poi.Title, err = poiIt.GetItVal(poiDef["Title"].(string), data)
		if err != nil {
			log.Println("Can't get PoI Title")
			log.Println(err)
			if !poiIt.Next() {
				break
			} else {
				continue
			}
		}
		//Optional
		_, ok := poiDef["Description"]
		if ok {
			d, _ := poiIt.GetItVal(poiDef["Description"].(string), data)
			poi.Description = &d
		}
		poi.Geotype, err = poiIt.GetItVal(poiDef["Geotype"].(string), data)
		if err != nil {
			log.Panic(err)
		}
		if poi.Geotype == "point" {
			gL, err := poiIt.GetGeo(poiDef, data)
			if err != nil {
				log.Println("Can't get PoI Location")
				log.Println(err)
				if !poiIt.Next() {
					break
				} else {
					continue
				}
			}
			gAry := make([]*pois.Geolocation, 1)
			gAry[0] = &gL
			poi.Geolocation = gAry
		} else if poi.Geotype == "geojson" {
			poi.Geotype, poi.Geolocation, err = poiIt.GetGeoJson(poiDef, data)
			if err != nil {
				log.Panic(err)
			}
		}
		poi.ID = poiIt.GetId(poiDef, data)
		if poi.ID == "geo" {
			SetGeoID(&poi)
		}
		rf.Poi = poi.ID

		//Optional
		poi.Refs = poiIt.GetRefs(rf, poiDef, data)

		_, ok = poiDef["Properties"]
		if ok {
			poi.Properties = poiIt.GetProps(poiDef["Properties"].(map[interface{}]interface{}), data)
		}
		poiAry[poiIt.Iteration] = &poi
		if !poiIt.Next() {
			break
		}
	}
	return poiAry
}

func (fx *JSONFetcher) Fetch(signals chan bool) {
	fx.rx.Fetch(signals, fx.Config, fx.process)
}
