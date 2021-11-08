package fetchers

import (
	. "github.com/onsi/gomega"
	"testing"

	"github.com/derbexuk/wurzel/harvester/feeds/config"
)

var data = `
{
  "range": "Sheet1!A1:E6",
  "majorDimension": "ROWS",
  "values": [
    [
      "name",
      "col1",
      "col2",
      "col3",
      "col4"
    ],
    [
      "aaa",
      "11",
      "22",
      "33",
      "44"
    ],
    [
      "bb",
      "11",
      "22",
      "33",
      "44"
    ],
    [
      "cc",
      "11",
      "22",
      "33",
      "44"
    ],
    [
      "dd",
      "11",
      "22",
      "33",
      "44"
    ],
    [
      "ww",
      "11",
      "22",
      "33",
      "44"
    ]
  ]
}
`

func TestGsheets(t *testing.T) {
	RegisterTestingT(t)

	yaml := `
---
format : GSHEET
frequency : 60m
feed source : https://sheets.googleapis.com/v4/spreadsheets/1h2jihswhVlUpysistWq0GWSshXxUyQcy2Wu1_Gpc1zQ/values/Sheet1!A1:E6
credentials :
    service file : jhtest-b1e83fac8378.json
paths :
  -
    source : .
    destination : /SiteMatch/people
`
	fc := config.FeedConfig{}
	fc.Populate([]byte(yaml))
	feed := &GSheetsFetcher{Config: fc}
	res := feed.process([]byte(data), nil)
	matchStr := `["aaa","11","22","33","44"]`
	Expect(res["/SiteMatch/people"]).Should(ContainSubstring(matchStr))
}

func TestGsheetsHdr(t *testing.T) {
	RegisterTestingT(t)

	yaml := `
---
format : GSHEET
frequency : 60m
feed source : https://sheets.googleapis.com/v4/spreadsheets/1h2jihswhVlUpysistWq0GWSshXxUyQcy2Wu1_Gpc1zQ/values/Sheet1!A1:E6
credentials :
    service file : jhtest-b1e83fac8378.json
params :
  headerrow : 0
paths :
  -
    source : .
    destination : /SiteMatch/people
`
	fc := config.FeedConfig{}
	fc.Populate([]byte(yaml))
	feed := &GSheetsFetcher{Config: fc}
	res := feed.process([]byte(data), nil)
	matchStr := `[{"col1":"11","col2":"22","col3":"33","col4":"44","name":"aaa"}`
	Expect(res["/SiteMatch/people"]).Should(ContainSubstring(matchStr))
}

func TestGsheetsHdrOffset(t *testing.T) {
	RegisterTestingT(t)

	yaml := `
---
format : GSHEET
frequency : 60m
feed source : https://sheets.googleapis.com/v4/spreadsheets/1h2jihswhVlUpysistWq0GWSshXxUyQcy2Wu1_Gpc1zQ/values/Sheet1!A1:E6
credentials :
    service file : jhtest-b1e83fac8378.json
params :
    #This tells us where to take the column names from (we map these to keys)
    headerrow : 0
    #Tells us where to start reading data from, default is next row after headerrow
    datastartrow : 2
paths :
  -
    source : .
    destination : /SiteMatch/people
`
	fc := config.FeedConfig{}
	fc.Populate([]byte(yaml))
	feed := &GSheetsFetcher{Config: fc}
	res := feed.process([]byte(data), nil)
	matchStr := `[{"col1":"11","col2":"22","col3":"33","col4":"44","name":"aaa"}`
	Expect(res["/SiteMatch/people"]).ShouldNot(ContainSubstring(matchStr))
	matchStr = `[{"col1":"11","col2":"22","col3":"33","col4":"44","name":"bb"}`
	Expect(res["/SiteMatch/people"]).Should(ContainSubstring(matchStr))
}

func TestGsheetsHdrMap(t *testing.T) {
	RegisterTestingT(t)

	yaml := `
---
format : GSHEET
frequency : 60m
feed source : https://sheets.googleapis.com/v4/spreadsheets/1h2jihswhVlUpysistWq0GWSshXxUyQcy2Wu1_Gpc1zQ/values/Sheet1!A1:E6
credentials :
    service file : jhtest-b1e83fac8378.json
params :
  headerrow : 0
  headermap :
    col1 : ht
    col3 : length
paths :
  -
    source : .
    destination : /SiteMatch/people
`
	fc := config.FeedConfig{}
	fc.Populate([]byte(yaml))
	feed := &GSheetsFetcher{Config: fc}
	res := feed.process([]byte(data), nil)
	matchStr := `{"col2":"22","col4":"44","ht":"11","length":"33","name":"dd"}`
	Expect(res["/SiteMatch/people"]).Should(ContainSubstring(matchStr))
}
