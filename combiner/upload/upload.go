package upload

import (
	"bytes"
	ds "github.com/derbexuk/wurzel/harvester/datastore"
	"github.com/derbexuk/wurzel/harvester/feeds/config"
	"github.com/derbexuk/wurzel/harvester/feeds/fetchers"
	"io"
	"mime/multipart"
	"os"
)

func UploadCsv(file multipart.File, configFile string) error {

	ds.Init()
	var Buf bytes.Buffer
	feedDir := os.Getenv("CSPACE_COLLECTIONS_FEED_DIR")
	if feedDir == "" {
		feedDir = "../upload/feeds.d/"
	}
	fc := config.FeedConfig{}
	fc.Read(feedDir + configFile)

	io.Copy(&Buf, file)
	contents := Buf.String()
	Buf.Reset()

	parser := &fetchers.SpreadSheetParser{Config: fc}
	result := parser.ParseCSV([]byte(contents))

	for path, data := range result {
		ds.Set(path, data)
		ds.PublishUpdate(path)
	}

	return nil
}

func FetchGsheet(configFile string) error {

	feedDir := os.Getenv("CSPACE_COLLECTIONS_FEED_DIR")
	if feedDir == "" {
		feedDir = "../upload/feeds.d/"
	}

	con := config.FeedConfig{}
	con.Read(feedDir + configFile)
	feed := &fetchers.GSheetsFetcher{Config: con, FeedDir: feedDir}
	feed.Fetch(nil)
	return nil
}
