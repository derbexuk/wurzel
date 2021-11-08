package fetchers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/derbexuk/wurzel/harvester/feeds/config"
)

/*
	Read test data file
*/
func readData(testname string) []byte {
	testdata := "./testdata/" + testname + ".txt"
	bites, err := ioutil.ReadFile(testdata) // just pass the file name
	if err != nil {
		log.Print(err)
	}
	// Uncomment to debug...
	// fmt.Println(b) 	  		// print the content as 'bytes'
	// strng := string(bites) 	// convert content to a 'string'
	// fmt.Println(strng)     	// print the content as a 'string'
	return bites
}

// ================== GSheet Tests ===================
var gsheetdata = readData("gsheetdata")

func TestParseGsheet(t *testing.T) {
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
	feed := &SpreadSheetParser{Config: fc}
	res := feed.ParseGsheet([]byte(gsheetdata))
	//	fmt.Println("res:", res) // ####
	matchStr := `["aaa","11","22","33","44"]`
	Expect(res["/SiteMatch/people"]).Should(ContainSubstring(matchStr))
}

func TestParseGsheetHdr(t *testing.T) {
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
	feed := &SpreadSheetParser{Config: fc}
	res := feed.ParseGsheet([]byte(gsheetdata))
	//	fmt.Println("res:", res) // ####
	matchStr := `[{"col1":"11","col2":"22","col3":"33","col4":"44","name":"aaa"}`
	Expect(res["/SiteMatch/people"]).Should(ContainSubstring(matchStr))
}

func TestParseGsheetHdrOffset(t *testing.T) {
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
	feed := &SpreadSheetParser{Config: fc}
	res := feed.ParseGsheet([]byte(gsheetdata))
	//	fmt.Println("res:", res) // ####
	matchStr := `[{"col1":"11","col2":"22","col3":"33","col4":"44","name":"aaa"}`
	Expect(res["/SiteMatch/people"]).ShouldNot(ContainSubstring(matchStr))
	matchStr = `[{"col1":"11","col2":"22","col3":"33","col4":"44","name":"bb"}`
	Expect(res["/SiteMatch/people"]).Should(ContainSubstring(matchStr))
}

func TestParseGsheetHdrMap(t *testing.T) {
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
	feed := &SpreadSheetParser{Config: fc}
	res := feed.ParseGsheet([]byte(gsheetdata))
	//	fmt.Println("res:", res) // ####
	matchStr := `{"col2":"22","col4":"44","ht":"11","length":"33","name":"dd"}`
	Expect(res["/SiteMatch/people"]).Should(ContainSubstring(matchStr))
}

// ================== CSV Tests ===================
var csvdata = readData("csvdata")

func TestParseCsv(t *testing.T) {
	RegisterTestingT(t)

	yaml := `
---
format : CSV
frequency : Once
feed source : https://sheets.googleapis.com/v4/spreadsheets/1h2jihswhVlUpysistWq0GWSshXxUyQcy2Wu1_Gpc1zQ/values/Sheet1!A1:E6
paths :
  -
    source : .
    destination : /SiteMatch/people
`
	fc := config.FeedConfig{}
	fc.Populate([]byte(yaml))
	feed := &SpreadSheetParser{Config: fc}
	res := feed.ParseCSV([]byte(csvdata))
	//	fmt.Println("res:", res) // ####
	matchStr := `["aaa","11","22","33","44"]`
	Expect(res["/SiteMatch/people"]).Should(ContainSubstring(matchStr))
}

func TestParseCsvHdr(t *testing.T) {
	RegisterTestingT(t)

	yaml := `
---
format : CSV
frequency : Once
feed source : https://sheets.googleapis.com/v4/spreadsheets/1h2jihswhVlUpysistWq0GWSshXxUyQcy2Wu1_Gpc1zQ/values/Sheet1!A1:E6
params :
  headerrow : 0
paths :
  -
    source : .
    destination : /SiteMatch/people
`
	fc := config.FeedConfig{}
	fc.Populate([]byte(yaml))
	feed := &SpreadSheetParser{Config: fc}
	res := feed.ParseCSV([]byte(csvdata))
	//	fmt.Println("res:", res) // ####
	matchStr := `[{"col1":"11","col2":"22","col3":"33","col4":"44","name":"aaa"}`
	Expect(res["/SiteMatch/people"]).Should(ContainSubstring(matchStr))
}

func TestParseCsvHdrOffset(t *testing.T) {
	RegisterTestingT(t)

	yaml := `
---
format : CSV
frequency : Once
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
	feed := &SpreadSheetParser{Config: fc}
	res := feed.ParseCSV([]byte(csvdata))
	//	fmt.Println("res:", res) // ####
	matchStr := `[{"col1":"11","col2":"22","col3":"33","col4":"44","name":"aaa"}`
	Expect(res["/SiteMatch/people"]).ShouldNot(ContainSubstring(matchStr))
	matchStr = `[{"col1":"11","col2":"22","col3":"33","col4":"44","name":"bb"}`
	Expect(res["/SiteMatch/people"]).Should(ContainSubstring(matchStr))
}

func TestParseCsvHdrMap(t *testing.T) {
	RegisterTestingT(t)

	yaml := `
---
format : CSV
frequency : Once
feed source : https://sheets.googleapis.com/v4/spreadsheets/1h2jihswhVlUpysistWq0GWSshXxUyQcy2Wu1_Gpc1zQ/values/Sheet1!A1:E6
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
	feed := &SpreadSheetParser{Config: fc}
	res := feed.ParseCSV([]byte(csvdata))
	//	fmt.Println("res:", res) // ####
	matchStr := `{"col2":"22","col4":"44","ht":"11","length":"33","name":"dd"}`
	Expect(res["/SiteMatch/people"]).Should(ContainSubstring(matchStr))
}

// ================== CSV Tests with realistic data ===================
// The input test data...
var csv01input = readData("csv01input")

// The expected result...
var csv01output = readData("csv01output")

func TestParseCsv01(t *testing.T) {
	RegisterTestingT(t)

	yaml := `
---
format : CSV
frequency : Once
feed source : https://sheets.googleapis.com/v4/spreadsheets/1h2jihswhVlUpysistWq0GWSshXxUyQcy2Wu1_Gpc1zQ/values/Sheet1!A1:E6
paths :
  -
    source : .
    destination : /SiteMatch/people
`
	fc := config.FeedConfig{}
	fc.Populate([]byte(yaml))
	feed := &SpreadSheetParser{Config: fc}
	res := feed.ParseCSV([]byte(csv01input))
	fmt.Println("res:", res) // ####
	matchStr := string(csv01output)
	Expect(res["/SiteMatch/people"]).Should(ContainSubstring(matchStr))
}

// The input test data...
var csv02input = readData("csv02input")

// The expected result...
var csv02output = readData("csv02output")

func TestParseCsv02(t *testing.T) {
	RegisterTestingT(t)

	yaml := `
---
format : CSV
frequency : Once
feed source : https://sheets.googleapis.com/v4/spreadsheets/1h2jihswhVlUpysistWq0GWSshXxUyQcy2Wu1_Gpc1zQ/values/Sheet1!A1:E6
paths :
  -
    source : .
    destination : /SiteMatch/people
`
	fc := config.FeedConfig{}
	fc.Populate([]byte(yaml))
	feed := &SpreadSheetParser{Config: fc}
	res := feed.ParseCSV([]byte(csv02input))
	fmt.Println("res:", res) // ####
	matchStr := string(csv02output)
	Expect(res["/SiteMatch/people"]).Should(ContainSubstring(matchStr))
}

// The input test data...
var csv03input = readData("csv03input")

// The expected result...
var csv03output = readData("csv03output")

func TestParseCsv03(t *testing.T) {
	RegisterTestingT(t)

	yaml := `
---
format : CSV
frequency : Once
feed source : https://sheets.googleapis.com/v4/spreadsheets/1h2jihswhVlUpysistWq0GWSshXxUyQcy2Wu1_Gpc1zQ/values/Sheet1!A1:E6
paths :
  -
    source : .
    destination : /SiteMatch/people
`
	fc := config.FeedConfig{}
	fc.Populate([]byte(yaml))
	feed := &SpreadSheetParser{Config: fc}
	res := feed.ParseCSV([]byte(csv03input))
	fmt.Println("res:", res) // ####
	matchStr := string(csv03output)
	Expect(res["/SiteMatch/people"]).Should(ContainSubstring(matchStr))
}

// The input test data...
var csv04input = readData("csv04input")

// The expected result...
var csv04output = readData("csv04output")

func TestParseCsv04(t *testing.T) {
	RegisterTestingT(t)

	yaml := `
---
format : CSV
frequency : Once
feed source : https://sheets.googleapis.com/v4/spreadsheets/1h2jihswhVlUpysistWq0GWSshXxUyQcy2Wu1_Gpc1zQ/values/Sheet1!A1:E6
paths :
  -
    source : .
    destination : /SiteMatch/people
`
	fc := config.FeedConfig{}
	fc.Populate([]byte(yaml))
	feed := &SpreadSheetParser{Config: fc}
	res := feed.ParseCSV([]byte(csv04input))
	fmt.Println("res:", res) // ####
	matchStr := string(csv04output)
	Expect(res["/SiteMatch/people"]).Should(ContainSubstring(matchStr))
	matchStr2 := "TEST"
	Expect(res["/SiteMatch/people"]).Should(ContainSubstring(matchStr2))
}

// The input test data...
var csv05input = readData("csv05input")

// The expected result...
var csv05output = readData("csv05output")

func TestParseCsv05(t *testing.T) {
	RegisterTestingT(t)

	yaml := `
---
format : CSV
frequency : Once
feed source : https://sheets.googleapis.com/v4/spreadsheets/1h2jihswhVlUpysistWq0GWSshXxUyQcy2Wu1_Gpc1zQ/values/Sheet1!A1:E6
paths :
  -
    source : .
    destination : /SiteMatch/people
`
	fc := config.FeedConfig{}
	fc.Populate([]byte(yaml))
	feed := &SpreadSheetParser{Config: fc}
	res := feed.ParseCSV([]byte(csv05input))
	fmt.Println("res:", res) // ####
	matchStr := string(csv05output)
	Expect(res["/SiteMatch/people"]).Should(ContainSubstring(matchStr))
}

// The input test data...
var csv06input = readData("csv06input")

// The expected result...
var csv06output = readData("csv06output")

func TestParseCsv06(t *testing.T) {
	RegisterTestingT(t)

	yaml := `
---
format : CSV
frequency : Once
feed source : https://sheets.googleapis.com/v4/spreadsheets/1h2jihswhVlUpysistWq0GWSshXxUyQcy2Wu1_Gpc1zQ/values/Sheet1!A1:E6
paths :
  -
    source : .
    destination : /SiteMatch/people
`
	fc := config.FeedConfig{}
	fc.Populate([]byte(yaml))
	feed := &SpreadSheetParser{Config: fc}
	res := feed.ParseCSV([]byte(csv06input))
	fmt.Println("res:", res) // ####
	matchStr := string(csv06output)
	Expect(res["/SiteMatch/people"]).Should(ContainSubstring(matchStr))
}

// The input test data...
var csv07input = readData("csv07input")

// The expected result...
var csv07output = readData("csv07output")

func TestParseCsv07(t *testing.T) {
	RegisterTestingT(t)

	yaml := `
---
format : CSV
frequency : Once
feed source : https://sheets.googleapis.com/v4/spreadsheets/1h2jihswhVlUpysistWq0GWSshXxUyQcy2Wu1_Gpc1zQ/values/Sheet1!A1:E6
paths :
  -
    source : .
    destination : /SiteMatch/people
`
	fc := config.FeedConfig{}
	fc.Populate([]byte(yaml))
	feed := &SpreadSheetParser{Config: fc}
	res := feed.ParseCSV([]byte(csv07input))
	fmt.Println("res:", res) // ####
	matchStr := string(csv07output)
	Expect(res["/SiteMatch/people"]).Should(ContainSubstring(matchStr))
}

func TestParseCsvEvent(t *testing.T) {
	RegisterTestingT(t)

	yaml := `
---
format : CSV
frequency : Once
feed source : https://sheets.googleapis.com/v4/spreadsheets/1h2jihswhVlUpysistWq0GWSshXxUyQcy2Wu1_Gpc1zQ/values/Sheet1!A1:E6
paths :
  -
    source : .
    destination : /SiteMatch/people
params :
  headerrow : 0
  generic :
    order : ["poi", "event"]
    poi :
      Id : "ID"
      Title : "ProjectName"
      Geotype : "=point"
      Lat : "Lat"
      Long : "Long"
    event :
      timeFormat : "=2006-01-02"
      Id : "=auto"
      Refs : "=poi"
      Title : "LBC Planning Reference"
      Start : "Construction Starts"
      End : "Construction Ends"
`

	csvInput := readData("csv_poi")
	fc := config.FeedConfig{}
	fc.Populate([]byte(yaml))
	feed := &SpreadSheetParser{Config: fc}
	res := feed.ParseCSV([]byte(csvInput))
	fmt.Println("res:", res) // ####
	gResult := GenericResult{}
	err := json.Unmarshal([]byte(res["/SiteMatch/people"]), &gResult)
	Expect(err).To(BeNil())
	Expect(gResult.Pois[0].ID).To(Equal("1"))
	Expect(*gResult.Pois[0].Geolocation[0].Latitude).To(Equal(51.376262))
	Expect(gResult.Events[1].Refs[0]).To(Equal("2"))
	Expect(gResult.Events[1].Start).To(Equal("2018-02-01T00:00:00Z"))
}

func TestParseCsvNearBy(t *testing.T) {
	RegisterTestingT(t)

	yaml := `
---
format : CSV
frequency : Once
feed source : https://sheets.googleapis.com/v4/spreadsheets/1h2jihswhVlUpysistWq0GWSshXxUyQcy2Wu1_Gpc1zQ/values/Sheet1!A1:E6
paths :
  -
    source : .
    destination : /SiteMatch/people
params :
    headerrow : 0
    headermap:
      title: Title
      latitude: Lat
      longitude: Long
      artistdescription: Desc

    generic:
      poi:
        Id: poid
        Geotype: "=point"
        Title: Title
        Description: Desc
        Lat: Lat
        Long: Long
        Properties :
            Category: category
            NearBy: nearby
`

	csvInput := readData("csv_nearby")
	fc := config.FeedConfig{}
	fc.Populate([]byte(yaml))
	feed := &SpreadSheetParser{Config: fc}
	res := feed.ParseCSV([]byte(csvInput))
	//fmt.Println("res:", res) // ####
	gResult := GenericResult{}
	err := json.Unmarshal([]byte(res["/SiteMatch/people"]), &gResult)
	Expect(err).To(BeNil())
	Expect(gResult.Pois[0].ID).To(Equal("art001"))
	Expect(*gResult.Pois[0].Geolocation[0].Latitude).To(Equal(51.3727097))
	Expect(gResult.Pois[0].Properties["Category"]).To(Equal("Croydon_Collection"))
	Expect(gResult.Pois[0].Properties["NearBy"]).Should(ContainSubstring(`{"poiId":"2410","title":"Luna"`))
}
