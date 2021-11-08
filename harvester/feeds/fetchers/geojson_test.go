package fetchers

import (
	"encoding/json"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"log"
	"testing"

	"github.com/derbexuk/wurzel/harvester/feeds/config"
)

func TestGeoJson(t *testing.T) {
	RegisterTestingT(t)

	testdata := "./testdata/Masterplan_boundaries.geojson"
	contents, err := ioutil.ReadFile(testdata) // just pass the file name
	if err != nil {
		log.Print(err)
	}

	genYaml := `
---
format : JSON
frequency : 60m
feed source : http://3foxinternational.com/endpoints/get_news.php?key=433114771&s=0&n=10
paths :
  -
    source : .
    destination : /airQuality/test
params :
  generic :
    poi :
      iterator : "features"
      Id : "iterator.properties.EntityHandle"
      Title : "iterator.properties.Layer"
      Geotype : "=geojson"
      Geometry : "iterator.geometry"
`

	fc := config.FeedConfig{}
	fc.Populate([]byte(genYaml))
	feed := &JSONFetcher{Config: fc}
	res := feed.process([]byte(contents), nil)
	gResult := GenericResult{}
	log.Println(res["/airQuality/test"])
	err = json.Unmarshal([]byte(res["/airQuality/test"]), &gResult)
	Expect(err).To(BeNil())
}
