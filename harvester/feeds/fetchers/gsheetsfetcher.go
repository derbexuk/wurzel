package fetchers

import (
	//"encoding/json"
	"io/ioutil"
	//"fmt"
	"log"
	//"strconv"
	//"strings"

	"github.com/derbexuk/wurzel/harvester/feeds/config"
	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
)

//GSheetsFetcher a REST Google Sheets feed
type GSheetsFetcher struct {
	FeedDir string
	Config  config.FeedConfig
	rx      RESTFetcher
}

//Process a sheet
//The output comes back as a map of paths to  data, either a string or strigified JSON
//
//There is probably only one path
//If there is a header row then we convert the strings to maps/objects
//If startrow is present in Params the we read the data from that row
//
//Some spreadsheet layout issues can also be fixed by just returning areas of the sheet, see
//example config file
func (fx *GSheetsFetcher) process(contents []byte, subsites map[string]string) map[string]string {
	feed := &SpreadSheetParser{Config: fx.Config}
	return feed.ParseGsheet(contents)
}

//Sheets uses a Google Service account and needs an OAuth connection created
//from the JWToken in service file
func (fx *GSheetsFetcher) Fetch(signals chan bool) {
	credFile := fx.FeedDir + "credentials/" + fx.Config.Credentials["service file"]
	log.Println(credFile)
	ctx := context.Background()

	b, err := ioutil.ReadFile(credFile)
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	gConfig, err := google.JWTConfigFromJSON([]byte(b), "https://www.googleapis.com/auth/spreadsheets.readonly")
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	fx.rx.feedClient = gConfig.Client(ctx)
	fx.rx.Fetch(signals, fx.Config, fx.process)
}
