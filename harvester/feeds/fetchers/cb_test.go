package fetchers

import (
	"encoding/json"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"log"
	"testing"

	"github.com/derbexuk/wurzel/harvester/feeds/config"
)

func TestCoinbase(t *testing.T) {
	RegisterTestingT(t)

	testdata := "./testdata/coinb-ex.json"
	contents, err := ioutil.ReadFile(testdata) // just pass the file name
	if err != nil {
		log.Print(err)
	}

	genYaml := `
---
format : JSON
frequency : 1m
feed source : https://api.coinbase.com/v2/exchange-rates
paths :
  -
    source : .
    destination : /finance/prices/exchangerates/coinbase
params :
  generic :
    event :
      Id : "=auto"
      Title : "=Exchange Rates"
      timeFormat : "=2006-01-02 15:04:05"
      Start: "=now"
      End : "=start"
      Properties : 
        BTC : "data.rates.BTC" 
        ETH : "data.rates.ETH" 
        ETC : "data.rates.ETC" 
        XAG : "data.rates.XAG" 
        XAu : "data.rates.XAU" 
        GBP : "data.rates.GBP" 
`

	fc := config.FeedConfig{}
	fc.Populate([]byte(genYaml))
	feed := &JSONFetcher{Config: fc}
	res := feed.process([]byte(contents), nil)
	gResult := GenericResult{}
	log.Println(res["/finance/prices/exchangerates/coinbase"])
	err = json.Unmarshal([]byte(res["/finance/prices/exchangerates/coinbase"]), &gResult)
	Expect(err).To(BeNil())
}
