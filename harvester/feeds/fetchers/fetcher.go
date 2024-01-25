package fetchers

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"time"
	"strings"

	ds "github.com/derbexuk/wurzel/harvester/datastore"
	"github.com/derbexuk/wurzel/harvester/feeds/config"
)

//Fetcher : Get some data
//Do something with it
type Fetcher interface {
	Fetch(chan bool)
	process([]byte, map[string]string) map[string]string
}

//RESTFetcher : here for type composition
type RESTFetcher struct {
	feedClient *http.Client
}

//Pull data from REST server
func (h RESTFetcher) pull(url string) []byte {
	if h.feedClient == nil {
		h.feedClient = &http.Client{
			Timeout: time.Second * 60,
		}
	}
	resp, err := h.feedClient.Get(url)
	if err != nil {
		log.Println(err)
		return nil
	}

	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil
	}

	return contents
}

//Regularly pull and process
//Updated paths are published on the feedupdates Redis channel
//Killed off by closing the signal channel
func (h *RESTFetcher) poll(signals chan bool, url string, freq string, process func([]byte, map[string]string) map[string]string, subs map[string]string) {
	//Sleep time
	d, err := time.ParseDuration(freq)
	if err != nil {
		log.Println(err)
	}

	//Only publish updated content
	var oldContents []byte
	for {
		ds.Init()
		contents := h.pull(url)
		if contents != nil && !bytes.Equal(contents, oldContents) {
			oldContents = contents
			result := process(contents, subs)

			for path, data := range result {
				ds.Set(path, data)
				ds.PublishUpdate(path)
			}
		}

		//Using time.After on a channel makes the sleep interruptable
		select {
		case <-time.After(d):
			log.Println("no activity on " + url)
			break
		case sig := <-signals:
			log.Println("received kill", sig)
			return
		}

	}
}

//Only supports one level of sub feed, should probably generate an array of URLS
func (fx *RESTFetcher) Fetch(signals chan bool, conf config.FeedConfig, process func([]byte, map[string]string) map[string]string) {
	log.Println("TYPE:", reflect.TypeOf(fx))

	//Pull only once
	if conf.Frequency == "Once" {
		ds.Init()
		contents := fx.pull(conf.Source)
		result := process(contents, nil)

		for path, data := range result {
			ds.Set(path, data)
			ds.PublishUpdate(path)
		}
		return
	}

	//Bare feed
	if len(conf.Subs) == 0 {
		log.Printf("Fetching URL : %s\n", conf.Source)
		go fx.poll(signals, conf.Source, conf.Frequency, process, nil)
	} else {
		for key, val := range conf.Subs {
			for _, param := range val {
				sourcePath := strings.Split(conf.Source,key)
				url := fmt.Sprintf("%s%s%s", sourcePath[0], param, sourcePath[1])
				log.Printf("Fetching SUBS URL : %s\n", url)
				go fx.poll(signals, url, conf.Frequency, process, map[string]string{key: param})
			}
		}
	}
}
