//Repeatedly fetch, parse and store the feeds defined in feeds.d
//The directory $CSPACE_FEED_DIR (e.g. /etc/cspace/feeds.d) is read at startup
//and periodically thereafter changes are detected and sub processes stopped, started and restarted as appropriate.
//The period can be set in $CSPACE_FEED_FREQ as a go duration (e.g. 5s)
//
//As currently envisioned there should only be one scheduler per host and one processor per feed

package scheduler

import (
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/derbexuk/wurzel/harvester/feeds/config"
	"github.com/derbexuk/wurzel/harvester/feeds/fetchers"
)

type Scheduler struct {
	feedDir      string
	processChans map[string]chan bool
	freq         time.Duration
}

func (s *Scheduler) runFile(fileName string) {
	log.Println("Running " + fileName)
	c := config.FeedConfig{}
	c.Read(s.feedDir + fileName)

	var feed fetchers.Fetcher
	switch f := c.Format; f {
	case "CSV":
		feed = &fetchers.CSVFetcher{Config: c}
	case "GSHEET":
		feed = &fetchers.GSheetsFetcher{Config: c, FeedDir: s.feedDir}
	case "JSON":
		feed = &fetchers.JSONFetcher{Config: c}
	case "XML":
		feed = &fetchers.XMLFetcher{Config: c}
	case "XMLJSON":
		feed = &fetchers.XmlJsonFetcher{Config: c}
	default:
		log.Println("Unrecognised feed format ", f)
		return
	}

	signals := make(chan bool, 1)
	s.processChans[fileName] = signals
	log.Println("type:", reflect.TypeOf(feed))
	feed.Fetch(signals)
}

func (s *Scheduler) killFile(fileName string) {
	log.Println("KILLING  " + fileName)
	close(s.processChans[fileName])
	delete(s.processChans, fileName)
}

func (s *Scheduler) Setup() {

	s.feedDir = os.Getenv("CSPACE_FEED_DIR")
	if s.feedDir == "" {
		s.feedDir = "/etc/cspace/feeds.d/"
	}

	freq, err := time.ParseDuration(os.Getenv("CSPACE_FEED_FREQ"))
	if err != nil {
		log.Println(err)
		freq = 30 * time.Second
	}

	s.freq = freq
	s.processChans = make(map[string]chan bool)
}

func (s *Scheduler) Loop() {
	var lastReadFiles []os.FileInfo
	for {
		var newFiles []os.FileInfo
		allFiles, err := ioutil.ReadDir(s.feedDir)
		if err != nil {
			log.Panic(err)
		}

		for _, file := range allFiles {
			if strings.HasSuffix(file.Name(), ".yaml") {
				newFiles = append(newFiles, file)
			}
		}

		if len(lastReadFiles) == 0 {
			lastReadFiles = newFiles
			for _, file := range newFiles {
				s.runFile(file.Name())
			}
		} else {
			i := 0
			for _, file := range lastReadFiles {
				if i > len(newFiles)-1 {
					s.killFile(file.Name())
					continue
				}

				newFile := newFiles[i]
				if (file.Size() != newFile.Size()) || (file.ModTime() != newFile.ModTime()) {
					//Config file has changed
					if file.Name() == newFile.Name() {
						s.killFile(file.Name())
						s.runFile(file.Name())
					} else if newFile.Name() < file.Name() {
						for newFile.Name() < file.Name() {
							s.runFile(newFile.Name())
							i++
							newFile = newFiles[i]
						}
						i--
					} else {
						s.killFile(file.Name())
						continue
					}
				}
				i++

			}
			for i < len(newFiles) {
				newFile := newFiles[i]
				s.runFile(newFile.Name())
				i++
			}
		}

		lastReadFiles = newFiles
		time.Sleep(s.freq)
	}
}
