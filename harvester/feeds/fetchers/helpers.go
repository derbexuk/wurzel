package fetchers

import (
	"crypto/sha1"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/derbexuk/wurzel/combiner/pois"
	"github.com/teris-io/shortid"
)

//This has to be called after the Geolocations are ste into the PoI object
func SetGeoID(poi *pois.Poi) {
	path := ""
	for _, loc := range poi.Geolocation {
		path += fmt.Sprintf("%f%f", *loc.Latitude, *loc.Longitude)
	}
	poi.ID = fmt.Sprintf("%x", sha1.Sum([]byte(path)))
}

//func getVal(src string, data map[string]interface{}) (string, error) {
func getVal(src string, data map[string]interface{}) (interface{}, error) {
	if strings.HasPrefix(src, "=") {
		return src[1:], nil
	}
	dPath := strings.SplitN(src, ".", 2)
	if len(dPath) == 1 {
		val, ok := data[dPath[0]]
		if ok {
			return val, nil
			//return val.(string), nil
		}
		return "", fmt.Errorf("Value not found for %s", dPath[0])
	} else {
		//Is it an array element we are looking for?
		idx := strings.IndexRune(dPath[0], '[')
		if idx > 0 {
			elIdx, _ := strconv.Atoi(dPath[0][idx+1 : len(dPath[0])-1])
			val := data[dPath[0][:idx]].([]interface{})[elIdx]
			return getVal(dPath[1], val.(map[string]interface{}))
		} else {
			val, ok := data[dPath[0]]
			if !ok {
				return "", fmt.Errorf("Value not found for %s", dPath[0])
			}
			return getVal(dPath[1], val.(map[string]interface{}))
		}
	}
}

type FeedIterator struct {
	Iterator  string
	Length    int
	Iteration int
}

func (it *FeedIterator) New(def map[interface{}]interface{}, data map[string]interface{}) error {
	var err error
	iterator, ok := def["iterator"].(string)
	if ok {
		it.Iterator = iterator
		it.Length, err = getLen(iterator, data)
		if err != nil {
			log.Println(err)
			return err
		}
		log.Printf("%s : %d\n", it.Iterator, it.Length)
	}
	return nil
}

func (it *FeedIterator) Next() bool {
	it.Iteration += 1
	if it.Iteration >= it.Length {
		return false
	}
	return true
}

func (it *FeedIterator) GetProps(propDef map[interface{}]interface{}, data map[string]interface{}) map[string]string {
	properties := make(map[string]string)
	for target, src := range propDef {
		log.Printf("%s : %s\n", src.(string), target.(string))
		value, err := it.GetItVal(src.(string), data)
		if err != nil {
			log.Printf("GetProps : %v", err)
			continue
		}
		properties[target.(string)] = value
		log.Println(value)
	}
	return properties
}

func (it *FeedIterator) GetItValRaw(src string, data map[string]interface{}) (interface{}, error) {
	//Not an iterating record set
	if it.Length == 0 {
		return getVal(src, data)
	}
	//Not an iterating field
	if !strings.HasPrefix(src, "iterator") {
		return getVal(src, data)
	}
	itBits := strings.SplitN(src, ".", 2)
	selector := fmt.Sprintf("%s[%d].%s", it.Iterator, it.Iteration, itBits[1])
	log.Println(selector)
	return getVal(selector, data)
}

func (it *FeedIterator) GetItVal(src string, data map[string]interface{}) (string, error) {
	val, err := it.GetItValRaw(src, data)
	return val.(string), err
}

func getLen(src string, data map[string]interface{}) (int, error) {
	dPath := strings.SplitN(src, ".", 2)
	if len(dPath) == 1 {
		val, ok := data[dPath[0]]
		if ok {
			return len(val.([]interface{})), nil
		}
		return -1, fmt.Errorf("Value not found for %s", dPath[0])
	} else {
		val, ok := data[dPath[0]]
		if !ok {
			return -1, fmt.Errorf("Value not found for %s", dPath[0])
		}
		return getLen(dPath[1], val.(map[string]interface{}))
	}
}

func (it *FeedIterator) GetTimes(def map[interface{}]interface{}, data map[string]interface{}) (start string, end string) {
	//We want to have a standard time format
	start, err := it.GetItVal(def["Start"].(string), data)
	if err != nil {
		log.Panic("Missing start time : "+ err.Error())
	}
	//So we have to specify the input format (in Go style)
	timeFormat, err := it.GetItVal(def["timeFormat"].(string), data)
	if err != nil {
		log.Panic(err)
	}

	//We have to specify an end but, it can be set to the start or forever
	end, err = it.GetItVal(def["End"].(string), data)
	if err != nil {
		log.Panic(err)
	}
	return GetKwdTimes(start, end, timeFormat)
}

func GetKwdTimes(startkwd, endkwd, timeFormat string) (start string, end string) {
	var tStart time.Time
	var err error

	if startkwd == "now" {
		tStart = time.Now()
	} else {
		tStart, err = time.Parse(timeFormat, startkwd)
		if err != nil {
			log.Panic(err)
		}
	}
	start = tStart.Format(time.RFC3339)

	//We have to specify an end but, it can be set to the start or forever
	if endkwd == "start" {
		end = tStart.Format(time.RFC3339)
	} else if endkwd == "-1" { /* Make generic */
		end = tStart.AddDate(0, 0, -1).Format(time.RFC3339)
	} else {
		tEnd, err := time.Parse(timeFormat, endkwd)
		if err != nil {
			log.Panic(err)
		}
		end = tEnd.Format(time.RFC3339)
	}
	return start, end
}

func (it *FeedIterator) GetId(def map[interface{}]interface{}, data map[string]interface{}) string {
	id, err := it.GetItVal(def["Id"].(string), data)
	if err != nil {
		log.Panic(err)
	}
	if id == "auto" {
		id, _ = shortid.Generate()
	}
	//Handle this back in poi creation
	if id == "geo" {
		id = "geo"
	}
	return id
}

func (it *FeedIterator) GetRefs(rf *References, def map[interface{}]interface{}, data map[string]interface{}) []string {
	_, ok := def["Refs"]
	if ok {
		ref, err := it.GetItVal(def["Refs"].(string), data)
		if err != nil {
			log.Panic(err)
		}
		if ref == "poi" {
			return []string{rf.Poi}
		} else if ref == "event" {
			return []string{rf.Event}
		} else if ref == "organism" {
			return []string{rf.Organism}
		} else {
			return []string{ref}
		}
	} else {
		return []string{}
	}
}

func (it *FeedIterator) GetGeo(def map[interface{}]interface{}, data map[string]interface{}) (pois.Geolocation, error) {
	latStr, err := it.GetItVal(def["Lat"].(string), data)
	lat, err := strconv.ParseFloat(strings.TrimSpace(latStr), 64)
	if err != nil {
		return pois.Geolocation{}, err
	}
	longStr, err := it.GetItVal(def["Long"].(string), data)
	long, err := strconv.ParseFloat(strings.TrimSpace(longStr), 64)
	if err != nil {
		return pois.Geolocation{}, err
	}
	gL := pois.Geolocation{Latitude: &lat, Longitude: &long}
	return gL, nil
}

//Partially supports the Geometry part of RFC 7946 'GeoJson Format'
func (it *FeedIterator) GetGeoJson(def map[interface{}]interface{}, data map[string]interface{}) (string, []*pois.Geolocation, error) {
	log.Println("GEO : " + def["Geometry"].(string))
	geom, err := it.GetItValRaw(def["Geometry"].(string), data)
	if err != nil {
		log.Panic(err)
	}

	var locs []*pois.Geolocation
	log.Println(geom.(map[string]interface{})["type"].(string))
	geoType := ""
	switch geom.(map[string]interface{})["type"].(string) {
	case "Point":
		geoType = "point"
		coords := geom.(map[string]interface{})["coordinates"].([]interface{})
		long := coords[0].(float64)
		lat := coords[1].(float64)
		locs = make([]*pois.Geolocation, 1)
		locs[0].Latitude = &lat
		locs[0].Longitude = &long
	case "LineString":
		geoType = "line"
		locs = makeLocArray(geom.(map[string]interface{})["coordinates"].([]interface{}))
	case "Polygon":
		geoType = "area"
		locs = makeLocArray(geom.(map[string]interface{})["coordinates"].([]interface{}))
	}
	return geoType, locs, nil
}

func makeLocArray(geos []interface{}) []*pois.Geolocation {
	var locs []*pois.Geolocation
	locs = make([]*pois.Geolocation, len(geos))
	for i, geo := range geos {
		coords := geo.([]interface{})
		long := coords[0].(float64)
		lat := coords[1].(float64)
		loc := pois.Geolocation{Latitude: &lat, Longitude: &long}
		locs[i] = &loc
	}
	return locs
}
