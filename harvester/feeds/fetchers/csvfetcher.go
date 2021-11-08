package fetchers

import (
	"github.com/derbexuk/wurzel/harvester/feeds/config"
)

//JSON Fetcher a REST JSON feed
type CSVFetcher struct {
	Config config.FeedConfig
	rx     RESTFetcher
}

func (fx *CSVFetcher) process(contents []byte, subsites map[string]string) map[string]string {
	feed := &SpreadSheetParser{Config: fx.Config}
	return feed.ParseCSV(contents)
}

func (fx *CSVFetcher) Fetch(signals chan bool) {
	fx.rx.Fetch(signals, fx.Config, fx.process)
}
