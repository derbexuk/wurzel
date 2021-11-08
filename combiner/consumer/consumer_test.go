package main

import (
	"context"
	. "github.com/onsi/gomega"
	"os"
	"testing"

	ds "github.com/derbexuk/wurzel/harvester/datastore"
	"github.com/derbexuk/wurzel/harvester/datawarehouse"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var dw = datawarehouse.Datawarehouse{}

func TestMain(m *testing.M) {
	DB = "collections-test"

	dw.Open()
	defer dw.Close()

	ds.Init()
	d := dw.Client.Database(DB)
	d.Drop(context.Background())
	ds.Del("/airQuality/test")
	code := m.Run()
	os.Exit(code)
}

func TestRead(t *testing.T) {
	RegisterTestingT(t)

	ccs := ConsumerConfigs{}
	err := ccs.Read("test.yaml")
	Expect(err).To(BeNil())
	Expect(len(ccs.Sources)).To(Equal(2))
	Expect(ccs.Sources["/airQuality/test"]["poiDestPath"]).To(Equal("/pois/environment/croydon/monitoring"))
}

func TestProcess(t *testing.T) {
	var testGR = `{"Pois":[{"ID":"P10d7B5mRz","Title":"Croydon - Norbury","Location":"","Geotype":"point","Description":"","Refs":["EJAd7Bcmg"],"Properties":null,"Deactivated":false,"Geolocation":[{"latitude":51.411349,"longitude":-0.12311}]},{"ID":"P10OnB5mR7","Title":"Croydon - Norbury","Location":"","Geotype":"point","Description":"","Refs":["P10O7B5mRM"],"Properties":null,"Deactivated":false,"Geolocation":[{"latitude":51.411349,"longitude":-0.12311}]}],"Events":[{"ID":"EJAO7fcmgm","Title":"Nitrogen Dioxide","Description":"","Deactivated":false,"Start":"2018-07-06T07:00:00Z","End":"2018-07-06T07:00:00Z","Properties":{"AirQualityBand":"Low"},"Refs":["P10d7B5mRz"]},{"ID":"E1Ad7BcmRZ","Title":"PM10","Description":"","Deactivated":false,"Start":"2018-07-06T07:00:00Z","End":"2018-07-06T07:00:00Z","Properties":{"AirQualityBand":"High"},"Refs":["P10d7B5mRz"]},{"ID":"E1Ad7BcmRV","Title":"Nitrogen Dioxide","Description":"","Deactivated":false,"Start":"2018-07-06T07:00:00Z","End":"2018-07-06T07:00:00Z","Properties":{"AirQualityBand":"Low"},"Refs":["P10OnB5mR7"]},{"ID":"P10d7fcigI","Title":"PM10","Description":"","Deactivated":false,"Start":"2018-07-06T07:00:00Z","End":"2018-07-06T07:00:00Z","Properties":{"AirQualityBand":"High"},"Refs":["P10OnB5mR7"]}],"Organisms":[{"ID":"EJAd7Bcmg","Title":"Croydon","Description":"Croydon - Norbury","Refs":[],"Properties":{"Lat":"51.372361","Long":"-0.100401"},"Deactivated":false},{"ID":"P10O7B5mRM","Title":"Croydon","Description":"Croydon - Norbury","Refs":[],"Properties":{"Lat":"51.372361","Long":"-0.100401"},"Deactivated":false}]}`

	RegisterTestingT(t)

	ds.Set("/airQuality/test", testGR)

	ccs := ConsumerConfigs{}
	err := ccs.Read("test.yaml")
	Expect(err).To(BeNil())

	err = ccs.Process("/airQuality/test")
	Expect(err).To(BeNil())

	ps := dw.GetLike(DB, "pois", "/pois")
	Expect(len(ps)).To(Equal(2))
	p := ps[0].Data
	Expect(p["refs"].(primitive.A)[0].(string)).To(Equal("/organisms/local-authorities/EJAd7Bcmg"))

	evts := dw.GetLike(DB, "events", "/events")
	Expect(len(evts)).To(Equal(4))
	e := evts[3].Data
	Expect(e["refs"].(primitive.A)[0].(string)).To(Equal("/pois/environment/croydon/monitoring/P10OnB5mR7"))

	orgs := dw.GetLike(DB, "organisms", "/organisms")
	Expect(len(orgs)).To(Equal(2))
	Expect(len(orgs[0].Data["refs"].(primitive.A))).To(Equal(0))
}

func TestPathSub(t *testing.T) {
	var testGR = `{"Pois":[{"ID":"984","Title":"Barclays Bank (Norfolk House)","Location":"","Geotype":"point","Description":"","Refs":[],"Properties":{"Category":"Financial Services"},"Deactivated":false,"Geolocation":[{"latitude":51.374407,"longitude":-0.096452}]},{"ID":"985","Title":"Barclays Bank (North End)","Location":"","Geotype":"point","Description":"","Refs":[],"Properties":{"Category":"Financial Services"},"Deactivated":false,"Geolocation":[{"latitude":51.374381,"longitude":-0.096438}]},{"ID":"2940","Title":"South West London Law Centres","Location":"","Geotype":"point","Description":"","Refs":[],"Properties":{"Category":"Legal Services"},"Deactivated":false,"Geolocation":[{"latitude":51.371282,"longitude":-0.099516}]},{"ID":"2941","Title":"The Whitgift Foundation","Location":"","Geotype":"point","Description":"","Refs":[],"Properties":{"Category":"Charity"},"Deactivated":false,"Geolocation":[{"latitude":51.37643,"longitude":-0.100456}]},{"ID":"2943","Title":"Five Guys","Location":"","Geotype":"point","Description":"","Refs":[],"Properties":{"Category":"Eating \u0026 Drinking"},"Deactivated":false,"Geolocation":[{"latitude":51.37332,"longitude":-0.100082}]}],"Events":null,"Organisms":null}`

	RegisterTestingT(t)

	c := dw.Client.Database(DB).Collection("pois")
	c.DeleteMany(context.Background(), bson.D{})

	ds.Set("/airQuality/test2", testGR)

	ccs := ConsumerConfigs{}
	err := ccs.Read("test.yaml")
	Expect(err).To(BeNil())

	err = ccs.Process("/airQuality/test2")
	Expect(err).To(BeNil())

	ps := dw.GetLike(DB, "pois", "/pois")
	Expect(len(ps)).To(Equal(5))
	Expect(ps[0].Path).To(Equal("/pois/business/Financial Services/984"))
	Expect(ps[4].Path).To(Equal("/pois/business/Eating \u0026 Drinking/2943"))
}
