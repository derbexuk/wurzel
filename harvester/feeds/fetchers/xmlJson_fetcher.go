package fetchers

import (
	"bytes"
	"github.com/derbexuk/wurzel/harvester/feeds/config"
	"log"

	xj "github.com/basgys/goxml2json"
)

//JSON Fetcher a REST JSON feed
type XmlJsonFetcher struct {
	Config config.FeedConfig
	rx     RESTFetcher
}

func (fx *XmlJsonFetcher) process(contents []byte, subsites map[string]string) map[string]string {
	jx := JSONFetcher{Config: fx.Config, rx: fx.rx}
	xml := bytes.NewReader(contents)
	jsbuf, err := xj.Convert(xml)
	if err != nil {
		log.Panic(err)
	}

	return jx.process(jsbuf.Bytes(), subsites)
}

func (fx *XmlJsonFetcher) Fetch(signals chan bool) {
	fx.rx.Fetch(signals, fx.Config, fx.process)
}
