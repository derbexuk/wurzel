package fetchers

import (
	"encoding/json"
	. "github.com/onsi/gomega"
	"log"
	"testing"

	"github.com/derbexuk/wurzel/harvester/feeds/config"
)

func TestSimpleJson(t *testing.T) {
	RegisterTestingT(t)

	yaml := `
---
format : JSON
frequency : 60m
feed source : http://3foxinternational.com/endpoints/get_news.php?key=433114771&s=0&n=10
paths :
  -
    source : .
    destination : /SiteMatch/news
`
	data := `
  {
  "returnCode":true,
  "errorMsg":"",
  "numRecords":543,
  "recordsReturned":10,
  "0":{"id":"8454",
    "headline":"Khan doubles housing targets ",
    "description":"Mayor of London Sadiq Khan announced plans on 29 November to build more than 250,000 homes in the capital's 13 outer suburbs as part of his draft London Plan.",
    "author":null,
    "body":"Story body",
    "image":"http:\/\/3foxinternational.com\/display_image.php?p=2930",
    "date_added":"2017-12-01 10:58:58"},
  "1":{"id":"8453",
    "headline":"Workshops return for Sitematch 2018",
    "description":null,
    "author":null,
    "body":"Story body",
    "image":"http:\/\/3foxinternational.com\/display_image.php?p=2931",
    "date_added":"2017-12-01 10:57:58"}
  }
`
	fc := config.FeedConfig{}
	fc.Populate([]byte(yaml))
	feed := &JSONFetcher{Config: fc}
	res := feed.process([]byte(data), nil)
	Expect(res["/SiteMatch/news"]).Should(ContainSubstring("Sadiq Khan"))
}

var genData = `
{
    "HourlyAirQualityIndex": {
        "@TimeToLive": "54",
        "LocalAuthority": {
            "@LocalAuthorityName": "Croydon",
            "@LaCentreLatitude": "51.372361",
            "@LaCentreLongitude": "-0.100401",
            "Site": {
                "@Longitude": "-0.12311",
                "@Latitude": "51.411349",
                "@SiteType": "Kerbside",
                "@SiteName": "Croydon - Norbury",
                "@BulletinDate": "2018-07-06 07:00:00",
                "species": {
                    "@SpeciesName": "Nitrogen Dioxide",
                    "@SpeciesCode": "NO2",
                    "@AirQualityIndex": "1",
                    "@AirQualityBand": "Low",
                    "@IndexSource": "Measurement"
                }
            }
        }
    }
}
`

func TestGenEvent(t *testing.T) {
	RegisterTestingT(t)

	genYaml := `
---
format : JSON
frequency : 60m
feed source : http://3foxinternational.com/endpoints/get_news.php?key=433114771&s=0&n=10
paths :
  -
    source : .
    destination : /airQuality/test
params :
  generic :
    event :
      timeFormat : "=2006-01-02 15:04:05"
      Id : "=auto"
      Title : "HourlyAirQualityIndex.LocalAuthority.Site.@SiteName"
      Start : "HourlyAirQualityIndex.LocalAuthority.Site.@BulletinDate"
      End : "=start"
      Description : "HourlyAirQualityIndex.LocalAuthority.Site.species.@IndexSource"
      Properties : 
        LocalAuthority : "HourlyAirQualityIndex.LocalAuthority.@LocalAuthorityName" 
        SiteType : "HourlyAirQualityIndex.LocalAuthority.Site.@SiteType"
        SpeciesName : "HourlyAirQualityIndex.LocalAuthority.Site.species.@SpeciesName"
        SpeciesCode : "HourlyAirQualityIndex.LocalAuthority.Site.species.@SpeciesCode"
        AirQualityIndex : "HourlyAirQualityIndex.LocalAuthority.Site.species.@AirQualityIndex"
        AirQualityBand : "HourlyAirQualityIndex.LocalAuthority.Site.species.@AirQualityBand"
`

	fc := config.FeedConfig{}
	fc.Populate([]byte(genYaml))
	feed := &JSONFetcher{Config: fc}
	res := feed.process([]byte(genData), nil)
	log.Println(res)
	Expect(res["/airQuality/test"]).Should(ContainSubstring(`"Title":"Croydon - Norbury"`))
	Expect(res["/airQuality/test"]).Should(ContainSubstring(`"AirQualityBand":"Low"`))
	Expect(res["/airQuality/test"]).Should(ContainSubstring(`"Description":"Measurement"`))
	Expect(res["/airQuality/test"]).Should(ContainSubstring(`"Start":"2018-07-06T07:00:00Z"`))
	Expect(res["/airQuality/test"]).Should(ContainSubstring(`"End":"2018-07-06T07:00:00Z"`))
}

func TestGenOrganism(t *testing.T) {
	RegisterTestingT(t)

	genYaml := `
---
format : JSON
frequency : 60m
feed source : http://3foxinternational.com/endpoints/get_news.php?key=433114771&s=0&n=10
paths :
  -
    source : .
    destination : /airQuality/test
params :
  generic :
    organism :
      Id : "=auto"
      Title : "HourlyAirQualityIndex.LocalAuthority.@LocalAuthorityName"
      Description : "HourlyAirQualityIndex.LocalAuthority.Site.@SiteName"
      Properties : 
        Lat : "HourlyAirQualityIndex.LocalAuthority.@LaCentreLatitude"
        Long : "HourlyAirQualityIndex.LocalAuthority.@LaCentreLongitude"
`

	fc := config.FeedConfig{}
	fc.Populate([]byte(genYaml))
	feed := &JSONFetcher{Config: fc}
	res := feed.process([]byte(genData), nil)
	log.Println(res)
	Expect(res["/airQuality/test"]).Should(ContainSubstring(`"Title":"Croydon"`))
	Expect(res["/airQuality/test"]).Should(ContainSubstring(`"Description":"Croydon - Norbury"`))
	Expect(res["/airQuality/test"]).Should(ContainSubstring(`"Properties":{"Lat":"51.372361","Long":"-0.100401"}`))
}

func TestGenPoi(t *testing.T) {
	RegisterTestingT(t)

	genYaml := `
---
format : JSON
frequency : 60m
feed source : http://3foxinternational.com/endpoints/get_news.php?key=433114771&s=0&n=10
paths :
  -
    source : .
    destination : /airQuality/test
params :
  generic :
    poi :
      Id : "=auto"
      Title : "HourlyAirQualityIndex.LocalAuthority.Site.@SiteName"
      Geotype : "=point"
      Lat : "HourlyAirQualityIndex.LocalAuthority.Site.@Latitude"
      Long : "HourlyAirQualityIndex.LocalAuthority.Site.@Longitude"
      Properties : 
        LocalAuthority : "HourlyAirQualityIndex.LocalAuthority.@LocalAuthorityName" 
        SiteType : "HourlyAirQualityIndex.LocalAuthority.Site.@SiteType"
`

	fc := config.FeedConfig{}
	fc.Populate([]byte(genYaml))
	feed := &JSONFetcher{Config: fc}
	res := feed.process([]byte(genData), nil)
	log.Println(res)
	Expect(res["/airQuality/test"]).Should(ContainSubstring(`"Title":"Croydon - Norbury"`))
	Expect(res["/airQuality/test"]).Should(ContainSubstring(`"Geotype":"point"`))
	Expect(res["/airQuality/test"]).Should(ContainSubstring(`"Geolocation":[{"Latitude":51.411349,"Longitude":-0.12311}]`))
	Expect(res["/airQuality/test"]).Should(ContainSubstring(`"Properties":{"LocalAuthority":"Croydon","SiteType":"Kerbside"}`))
}

func TestGenRefs(t *testing.T) {
	RegisterTestingT(t)

	genYaml := `
---
format : JSON
frequency : 60m
feed source : http://3foxinternational.com/endpoints/get_news.php?key=433114771&s=0&n=10
paths :
  -
    source : .
    destination : /airQuality/test
params :
  generic :
    order : ["poi", "event"]
    poi :
      Id : "=geo"
      Title : "HourlyAirQualityIndex.LocalAuthority.Site.@SiteName"
      Geotype : "=point"
      Lat : "HourlyAirQualityIndex.LocalAuthority.Site.@Latitude"
      Long : "HourlyAirQualityIndex.LocalAuthority.Site.@Longitude"
    event :
      timeFormat : "=2006-01-02 15:04:05"
      Id : "=auto"
      Refs : "=poi"
      Title : "HourlyAirQualityIndex.LocalAuthority.Site.@SiteName"
      Start : "HourlyAirQualityIndex.LocalAuthority.Site.@BulletinDate"
      End : "=start"
`

	fc := config.FeedConfig{}
	fc.Populate([]byte(genYaml))
	feed := &JSONFetcher{Config: fc}
	res := feed.process([]byte(genData), nil)
	gResult := GenericResult{}
	log.Println(res["/airQuality/test"])
	err := json.Unmarshal([]byte(res["/airQuality/test"]), &gResult)
	Expect(err).To(BeNil())

	Expect(gResult.Pois[0].Title).Should(Equal("Croydon - Norbury"))
	Expect(gResult.Pois[0].Geotype).Should(Equal("point"))
	Expect(*gResult.Pois[0].Geolocation[0].Latitude).Should(Equal(51.411349))
	Expect(gResult.Events[0].Start).Should(Equal("2018-07-06T07:00:00Z"))
	Expect(gResult.Events[0].Refs[0]).Should(Equal(gResult.Pois[0].ID))
}

func TestGetArray(t *testing.T) {
	RegisterTestingT(t)

	contents := `
    [ {"HourlyAirQualityIndex": {
        "@TimeToLive": "54",
        "LocalAuthority": {
            "@LocalAuthorityName": "Croydon",
            "@LaCentreLatitude": "51.372361",
            "@LaCentreLongitude": "-0.100401",
            "Site": {
                "@SiteName": "Croydon - Norbury",
                "@BulletinDate": "2018-07-06 07:00:00",
                "@Longitude": "-0.12311",
                "@Latitude": "51.411349",
                "species": [ {
                    "@SpeciesName": "Nitrogen Dioxide",
                    "@AirQualityBand": "Low"
                  }, 
                 {
                    "@SpeciesName": "PM10",
                    "@AirQualityBand": "High"
                  } ] 
            }
        }
    }},
    {"HourlyAirQualityIndex": {
        "@TimeToLive": "54",
        "LocalAuthority": {
            "@LocalAuthorityName": "Croydon",
            "@LaCentreLatitude": "51.372361",
            "@LaCentreLongitude": "-0.100401",
            "Site": {
                "@SiteName": "Croydon - Norbury",
                "@BulletinDate": "2018-07-06 07:00:00",
                "@Longitude": "-0.12311",
                "@Latitude": "51.411349",
                "species": [ {
                    "@SpeciesName": "Nitrogen Dioxide",
                    "@AirQualityBand": "Low"
                  }, 
                 {
                    "@SpeciesName": "PM10",
                    "@AirQualityBand": "High"
                  } ] 
            }
        }
    }}]
`

	genYaml := `
---
format : JSON
frequency : 60m
feed source : http://3foxinternational.com/endpoints/get_news.php?key=433114771&s=0&n=10
paths :
  -
    source : .
    destination : /airQuality/test
params :
  generic :
    order : ["organism", "poi", "event"]
    poi :
      Id : "=auto"
      Title : "HourlyAirQualityIndex.LocalAuthority.Site.@SiteName"
      Geotype : "=point"
      Lat : "HourlyAirQualityIndex.LocalAuthority.Site.@Latitude"
      Long : "HourlyAirQualityIndex.LocalAuthority.Site.@Longitude"
      Refs : "=organism"
    organism :
      Id : "=auto"
      Title : "HourlyAirQualityIndex.LocalAuthority.@LocalAuthorityName"
      Description : "HourlyAirQualityIndex.LocalAuthority.Site.@SiteName"
      Properties : 
        Lat : "HourlyAirQualityIndex.LocalAuthority.@LaCentreLatitude"
        Long : "HourlyAirQualityIndex.LocalAuthority.@LaCentreLongitude"
    event :
      iterator : "HourlyAirQualityIndex.LocalAuthority.Site.species"
      timeFormat : "=2006-01-02 15:04:05"
      Id : "=auto"
      Refs : "=poi"
      Title : "iterator.@SpeciesName"
      Start : "HourlyAirQualityIndex.LocalAuthority.Site.@BulletinDate"
      End : "=start"
      Properties :
        AirQualityBand : "iterator.@AirQualityBand"
`

	fc := config.FeedConfig{}
	fc.Populate([]byte(genYaml))
	feed := &JSONFetcher{Config: fc}
	res := feed.process([]byte(contents), nil)
	gResult := GenericResult{}
	log.Println(res["/airQuality/test"])
	err := json.Unmarshal([]byte(res["/airQuality/test"]), &gResult)
	Expect(err).To(BeNil())
	Expect(len(gResult.Events)).To(Equal(4))
	Expect(len(gResult.Pois)).To(Equal(2))
	Expect(len(gResult.Organisms)).To(Equal(2))
	Expect(gResult.Events[0].Refs[0]).To(Equal(gResult.Pois[0].ID))
	Expect(gResult.Pois[0].Refs[0]).To(Equal(gResult.Organisms[0].ID))
}

var contents1 = `
{
    "HourlyAirQualityIndex": {
        "@TimeToLive": "54",
        "LocalAuthority": {
            "@LocalAuthorityName": "Croydon",
            "Site": {
                "@SiteName": "Croydon - Norbury",
                "species": [{
                    "@AirQualityBand": "High"
                } ,
               {
                    "@AirQualityBand": "Low"
                } ]
            }
        }
    }
}
`

func TestGetVal(t *testing.T) {
	RegisterTestingT(t)

	var data map[string]interface{}
	if err := json.Unmarshal([]byte(contents1), &data); err != nil {
		log.Println(contents1)
		log.Panic(err)
	}

	val, err := getVal("HourlyAirQualityIndex.LocalAuthority.@LocalAuthorityName", data)
	Expect(err).To(BeNil())
	Expect(val.(string)).To(Equal("Croydon"))
	val, err = getVal("HourlyAirQualityIndex.LocalAuthority.@Arse", data)
	Expect(err.Error()).To(Equal("Value not found for @Arse"))
	val, err = getVal("HourlyAirQualityIndex.LocalAuthority.Site.species[1].@AirQualityBand", data)
	Expect(err).To(BeNil())
	Expect(val.(string)).To(Equal("Low"))
}

func TestGetLen(t *testing.T) {
	RegisterTestingT(t)

	var data map[string]interface{}
	if err := json.Unmarshal([]byte(contents1), &data); err != nil {
		log.Println(contents1)
		log.Panic(err)
	}

	val, err := getLen("HourlyAirQualityIndex.LocalAuthority.Site.species", data)
	Expect(err).To(BeNil())
	Expect(val).To(Equal(2))
}
