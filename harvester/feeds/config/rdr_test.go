package config

import (
	. "github.com/onsi/gomega"
	"testing"
	"time"
)

const yaml_config = `
---
format : XML
frequency : 60m
credentials :
  ApiKey : KEY123
feed source : http://api.erg.kcl.ac.uk/AirQuality/Hourly/MonitoringIndex/
sub feeds :
  SiteCode :
    - CR5
paths :
  -
    source : ./HourlyAirQualityIndex/LocalAuthority/Site/species/attrs/AirQualityBand
    destination : /environment/london/|SiteCode|/air-quality/band
    mungers : toupper
  -
    source : ./HourlyAirQualityIndex/LocalAuthority/Site/species/attrs/AirQualityIndex
    destination : /environment/london/|SiteCode|/air-quality/index
`

func TestReadFail(t *testing.T) {

	RegisterTestingT(t)

	Expect(func() {
		fc := FeedConfig{}
		fc.Read("reference_feed_config.arse")
	}).Should(Panic())
}

func TestRead(t *testing.T) {
	RegisterTestingT(t)

	fc := FeedConfig{}
	fc.Read("reference_feed_config.yaml")
	Expect(fc.Format).To(Equal("XML"))
}

func TestDuffJSON(t *testing.T) {
	RegisterTestingT(t)
	Expect(func() {
		fc := FeedConfig{}
		fc.Populate([]byte(": value"))
	}).Should(Panic())
}

func TestPopulate(t *testing.T) {
	RegisterTestingT(t)

	fc := FeedConfig{}
	fc.Populate([]byte(yaml_config))
	Expect(fc.Format).To(Equal("XML"))
	Expect(fc.Frequency).To(Equal("60m"))
	Expect(fc.Source).Should(ContainSubstring("MonitoringIndex"))
	Expect(fc.Credentials["ApiKey"]).To(Equal("KEY123"))
	Expect(fc.Subs["SiteCode"][0]).To(Equal("CR5"))
	Expect(fc.Paths[0]["destination"]).Should(ContainSubstring("london"))
	Expect(fc.Paths[0]["source"]).Should(ContainSubstring("AirQualityBand"))
	Expect(fc.Paths[0]["mungers"]).To(Equal("toupper"))
	Expect(fc.Paths[1]).ShouldNot(HaveKey("mungers"))
}

func TestNoSubs(t *testing.T) {
	RegisterTestingT(t)

	yaml := `
---
format : XML
frequency : 60m
feed source : http://api.erg.kcl.ac.uk/AirQuality/Hourly/MonitoringIndex/
paths :
`
	fc := FeedConfig{}
	fc.Populate([]byte(yaml))
	Expect(len(fc.Subs)).To(Equal(0))
}

func TestNoCred(t *testing.T) {
	RegisterTestingT(t)

	yaml := `
---
format : XML
frequency : 60m
feed source : http://api.erg.kcl.ac.uk/AirQuality/Hourly/MonitoringIndex/
paths :
`
	fc := FeedConfig{}
	fc.Populate([]byte(yaml))
	fc.ApplyCredentials()
	Expect(fc.Source).To(Equal("http://api.erg.kcl.ac.uk/AirQuality/Hourly/MonitoringIndex/"))
}

func TestCred(t *testing.T) {
	RegisterTestingT(t)

	yaml := `
---
format : XML
frequency : 60m
feed source : http://api.erg.kcl.ac.uk/|ApiKey|/AirQuality/Hourly/MonitoringIndex/
credentials :
  ApiKey : KEY123
paths :
`
	fc := FeedConfig{}
	fc.Populate([]byte(yaml))
	fc.ApplyCredentials()
	Expect(fc.Source).To(Equal("http://api.erg.kcl.ac.uk/KEY123/AirQuality/Hourly/MonitoringIndex/"))
}

func TestToBytes(t *testing.T) {
	RegisterTestingT(t)

	fc := FeedConfig{}
	_, err := fc.ToBytes()
	Expect(err.Error()).To(Equal("Config file must have a Format"))
	fc.Format = "XML"
	_, err = fc.ToBytes()
	Expect(err.Error()).To(Equal("Config file must have a Frequency"))
	fc.Frequency = "60s"
	_, err = fc.ToBytes()
	Expect(err.Error()).To(Equal("Config file must have a Source"))
	fc.Source = "http://api.erg.kcl.ac.uk/KEY123/AirQuality/Hourly/MonitoringIndex/"
	_, err = fc.ToBytes()
	Expect(err.Error()).To(Equal("Config file must have at least one Path"))

	var paths []map[string]string
	p := map[string]string{"source": "./HourlyAirQualityIndex/LocalAuthority/Site/species/attrs/AirQualityBand",
		"destination": "/environment/london/|SiteCode|/air-quality/band"}
	paths = append(paths, p)

	fc.Paths = paths
	yaml, err := fc.ToBytes()
	Expect(err).To(BeNil())
	Expect(string(yaml)).To(Equal(`format: XML
frequency: 60s
feed source: http://api.erg.kcl.ac.uk/KEY123/AirQuality/Hourly/MonitoringIndex/
paths:
- destination: /environment/london/|SiteCode|/air-quality/band
  source: ./HourlyAirQualityIndex/LocalAuthority/Site/species/attrs/AirQualityBand
`))
}

func TestDatDelta(t *testing.T) {
  RegisterTestingT(t)

  dstr, err := DateDelta("TODAY15D")
  Expect(err.Error()).To(Equal(`strconv.Atoi: parsing "5D": invalid syntax`))
  Expect(dstr).To(Equal(""))

  tm := time.Now()
  tm = tm.AddDate(0, 0, -15)

  dstr, err = DateDelta("TODAY-15D")
  Expect(err).Should(BeNil())
  Expect(dstr).To(Equal(tm.Format("2006/01/02")))

}

func TestExpandDates(t *testing.T) {
  RegisterTestingT(t)

  fc := FeedConfig{}
  fc.Source = "https://api.tfl.gov.uk/Road/All/Disruption?startDate=TODAY-15D&endDate=TODAY+03M/xxx"

  tm := time.Now()
  tm = tm.AddDate(0, 0, -15)
  d1 := tm.Format("2006/01/02")

  tm = time.Now()
  tm = tm.AddDate(0, 3, 0)
  d2 := tm.Format("2006/01/02")
  expectedStr := "https://api.tfl.gov.uk/Road/All/Disruption?startDate=" + d1 +
    "&endDate=" + d2 + "/xxx"

  fc.ExpandDates()
  Expect(fc.Source).To(Equal(expectedStr))

}
