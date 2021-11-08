package fetchers

import (
	"fmt"
	"log"
	"strings"

	xj "github.com/basgys/goxml2json"
	"github.com/beevik/etree"

	"github.com/derbexuk/wurzel/harvester/feeds/config"
)

//XMLFetcher a REST XML feed
type XMLFetcher struct {
	Config config.FeedConfig
	rx     RESTFetcher
}

//Process a document using XPATH
//The output comes back as a map of paths to  data, either a string or strigified JSON
//
//If the source path has /attrs/ in it then extract the named attribute.
//If the target of the source path is a bear element then get it's content
//If the target of the source path has children then extract and JSONify the subtree
//
//If the destination path has |XXX| XXX will be looked up in subsites and substitued for
//
//goxml2json prepends a - to attribute names and #content to element text may need to fork
//it and change the behaviour
func (fx *XMLFetcher) process(contents []byte, subsites map[string]string) map[string]string {
	results := make(map[string]string)
	doc := etree.NewDocument()
	if err := doc.ReadFromString(string(contents)); err != nil {
		log.Println(err)
	}

	for _, path := range fx.Config.Paths {
		data := ""
		if strings.Contains(path["source"], "/attrs/") {
			attrPath := strings.Split(path["source"], "/attrs/")
			if (attrPath != nil) && (len(attrPath) >= 2) {
				el := doc.FindElement(attrPath[0])
				if el == nil {
					log.Printf("Attribute element path not found : %s\n", attrPath[0])
				} else {
					data = el.SelectAttrValue(attrPath[1], "")
				}
			}
		} else {
			el := doc.FindElement(path["source"])
			if el == nil {
				log.Printf("Source path not found : %s\n", path["source"])
				continue
			}
			if el.ChildElements() == nil {
				data = el.Text()
			} else {
				//I imagine that this is horribly inefficient
				newDoc := doc.Copy()
				newDoc.SetRoot(el)
				str, err := newDoc.WriteToString()
				if err != nil {
					log.Printf("Failed to write doc for : %s\n", path["source"])
				}
				json, err := xj.Convert(strings.NewReader(str))
				if err != nil {
					log.Printf("JSON conversion failed for : %s\n", path["source"])
				}

				data = json.String()
			}
		}
		destination := path["destination"]
		//Almost supports multiple subsitution
		if strings.Contains(destination, "|") {
			destPath := strings.Split(destination, "|")
			destination = fmt.Sprintf("%s%s%s", destPath[0], subsites[destPath[1]], destPath[2])
		}
		results[destination] = data
	}
	//log.Println(results)
	return results
}

func (fx *XMLFetcher) Fetch(signals chan bool) {
	fx.rx.Fetch(signals, fx.Config, fx.process)
}
