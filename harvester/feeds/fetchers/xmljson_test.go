package fetchers

import (
	"encoding/json"

	//"bytes"
	"log"
	//xj "github.com/basgys/goxml2json"
	. "github.com/onsi/gomega"
	"testing"

	"github.com/derbexuk/wurzel/harvester/feeds/config"
)

func TestXmlJsonCnv(t *testing.T) {
	RegisterTestingT(t)

	genYaml := `
---
format : JSON
frequency : 60m
feed source : http://3foxinternational.com/endpoints/get_news.php?key=433114771&s=0&n=10
paths :
  -
    source : .
    destination : /croydon/biz/test
params :
  generic :
    poi :
      iterator : "nodes.node"
      Id : "iterator.Nid"
      Title : "iterator.Title"
      Geotype : "=point"
      Lat : "iterator.Latitude"
      Long : "iterator.Longitude"
      Properties :
        Category : "iterator.Category"
`

	fc := config.FeedConfig{}
	fc.Populate([]byte(genYaml))
	feed := &XmlJsonFetcher{Config: fc}
	jsbuf := readData("appfeedbiz.xml")
	res := feed.process(jsbuf, nil)
	log.Println(res)
	gResult := GenericResult{}
	log.Println(res["/croydon/biz/test"])
	err := json.Unmarshal([]byte(res["/croydon/biz/test"]), &gResult)
	Expect(err).To(BeNil())

	Expect(len(gResult.Pois)).To(Equal(5))

}
