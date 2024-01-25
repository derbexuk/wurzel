package fetchers

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"

	"github.com/derbexuk/wurzel/harvester/feeds/config"
)

type SpreadSheetParser struct {
	Config config.FeedConfig
}

//Read in a CSV sheet
func (ssp *SpreadSheetParser) ParseCSV(contents []byte) map[string]string {
	log.Println("Parsing CSV")
	reader := csv.NewReader(bytes.NewReader(contents))
	csv, err := reader.ReadAll()
	if err != nil {
		log.Println(string(contents))
		log.Panic(err)
	}
	//Rather tedious typecasting
	var sheet []interface{}
	for _, cRow := range csv {
		var row []interface{}
		for _, cell := range cRow {
			row = append(row, cell)
		}
		sheet = append(sheet, interface{}(row))
	}
	return (ssp).parse(sheet)
}

//Read in a GoogleSheet
func (ssp *SpreadSheetParser) ParseGsheet(contents []byte) map[string]string {
	/*
		unmarshal contents
	*/
	var data map[string]interface{}
	if err := json.Unmarshal(contents, &data); err != nil {
		log.Println(string(contents))
		log.Panic(err)
	}
	//Did Google retun an error
	if _, ok := data["error"]; ok {
		log.Panic(string(contents))
	}
	return (ssp).parse(data["values"].([]interface{}))
}

//	Parse a spread sheet
//	The output comes back as a map of paths to data, either a string or stringified JSON
//
//	There is probably only one path
//	If there is a header row then we convert the strings to maps/objects
//	If startrow is present in Params the we read the data from that row
//
//	Some spreadsheet layout issues can also be fixed by just returning areas of the sheet
//	see example config file
func (ssp *SpreadSheetParser) parse(valAry []interface{}) map[string]string {
	var vals []byte
	results := make(map[string]string)
	/*
		loop through Config.Paths, although it's a sheet so just one path
	*/
	for _, path := range ssp.Config.Paths {
		log.Printf("Path %v\n", path)
		/*
			get destination from Config.Paths
		*/
		destination := path["destination"]
		/*
			get header row from Config.Params
		*/
		//header row values will be used to create the attributes in JSON objects and we return an array of
		//those objects
		/*
			if there is a header row use the headermap from Config.Params to substitute different column headings
		*/
		hr, ok := ssp.Config.Params["headerrow"]
		if ok { //There is a header row (probably column titles)
			/*
				get the header row number
			*/
			hdrRow := hr.(int)
			// fmt.Println("hdrRow:", hdrRow) // #
			/*
				get the header map to overwrite the headers in the sheet
			*/
			_, ok := ssp.Config.Params["headermap"]
			hm := make(map[string]string)
			if ok {
				for k, v := range ssp.Config.Params["headermap"].(map[interface{}]interface{}) {
					hm[k.(string)] = v.(string)
				}
				log.Println(hm)
			}
			// fmt.Println("hm:", hm) // #
			/*
				build headers
			*/
			var headers []string
			for _, v := range valAry[hdrRow].([]interface{}) {
				hval, ok := hm[v.(string)]
				if ok {
					headers = append(headers, hval)
				} else {
					headers = append(headers, v.(string))
				}
			}
			// fmt.Println("headers:", headers) // #
			/*
				by default use the row after the header as the start of the data
				but check for an overriding data start row in Config.Params
			*/
			//Default data start is the next row
			startrow := hdrRow + 1
			ds, ok := ssp.Config.Params["datastartrow"]
			if ok { //Override the default data start row
				startrow = ds.(int)
			}
			fmt.Println("startrow:", startrow) // #
			/*
				combine header and data
			*/
			var dataObjs []map[string]interface{}
			fmt.Println("len(valAry):", len(valAry)) // #
			for i := startrow; i < len(valAry); i++ {
				var dm map[string]interface{}
				dm = make(map[string]interface{})
				for j, v := range valAry[i].([]interface{}) {
					dm[headers[j]] = v //.(string)
				}
				//x,_ := json.MarshalIndent(dm, "", "  ")
				//log.Println(string(x))
				dataObjs = append(dataObjs, dm)
			}
			fmt.Println("dataObjs:", dataObjs) // #
			//Is there a Generic definition in the config file
			gD, ok := ssp.Config.Params["generic"]
			if ok {
				//var dataAry []map[string]interface{}
				//dataAry = append(dataAry, dataObjs)

				results[destination] = makeGenerics(gD.(map[interface{}]interface{}), nil, dataObjs)
				return results
			}
			/*
				marshal into json []byte
			*/
			var err error
			if vals, err = json.Marshal(dataObjs); err != nil {
				log.Panic(err)
			}
		} else { //No header row, so can't convert to maps so just return 2D array
			var err error
			//log.Println(data["values"])
			/*
				marshal into json []byte
			*/
			if vals, err = json.Marshal(valAry); err != nil {
				log.Panic(err)
			}
		}
		/*
			add parsed data to the results map using destination as the key
		*/
		log.Println("PROCESSED")
		results[destination] = string(vals)
	}
	/*
		return the complete map of results
	*/
	return results
}
