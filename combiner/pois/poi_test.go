package pois

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/derbexuk/wurzel/harvester/datawarehouse"
	. "github.com/onsi/gomega"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	//"go.mongodb.org/mongo-driver/mongo/options"
)

var dw *datawarehouse.Datawarehouse

/*
  Test Main
*/
func TestMain(m *testing.M) {
	fmt.Println("TestMain")
	dw = &datawarehouse.Datawarehouse{}
	dw.Open()
	dw.EnsurePathIndex(DB, COLLECTION)

	c := dw.Client.Database(DB).Collection(COLLECTION)
	c.DeleteMany(context.Background(), bson.D{})

	code := m.Run()
	os.Exit(code)
}

func TestPoiCreate(t *testing.T) {
	RegisterTestingT(t)

	poi := createPayload()
	path := fmt.Sprintf("%s/%s", "/test", poi.ID)
	println("path value", path)
	Create(poi, "/test", dw)

	c := dw.Client.Database(DB).Collection(COLLECTION)
	res := c.FindOne(context.Background(), bson.D{{"path", path}})
	Expect(res).ShouldNot(BeNil())
	// Check retrieved data
	var result map[string]interface{}
	err := res.Decode(&result)
	Expect(err).Should(BeNil())
	Expect(result["data"].(map[string]interface{})["properties"].(map[string]interface{})["popcorn"].(string)).To(Equal("expensive"))
	Expect(result["data"].(map[string]interface{})["title"].(string)).To(Equal("Croydon"))
	Expect(result["data"].(map[string]interface{})["description"].(string)).To(Equal("A true multiplex"))
	Expect(result["data"].(map[string]interface{})["deactivated"].(bool)).To(Equal(false))
	Expect(result["data"].(map[string]interface{})["location"].(string)).To(Equal("back of the bus station"))
	Expect(result["data"].(map[string]interface{})["geotype"].(string)).To(Equal("point"))
	Expect(result["data"].(map[string]interface{})["refs"].(primitive.A)[0].(string)).To(Equal("/events/cinema/vue"))
}

// Test Duplicates
func TestDuplicates(t *testing.T) {
	RegisterTestingT(t)

	poi := createPayload()
	path := fmt.Sprintf("%s/%s", "/test", poi.ID)
	println(path)
	err := Create(poi, "/test", dw)
	// Check results
	Expect(err).ShouldNot(BeNil())
	Expect(err.Error()).Should(ContainSubstring("duplicate key error"))
}

/*func TestGetPath(t *testing.T) {
	RegisterTestingT(t)

	result := GetByPath(dw)
	Expect(len(result)).Should(BeNumerically(">", 1))
}*/

func TestGet(t *testing.T) {
	RegisterTestingT(t)
	path := "test/abc123"
	result := GetByPath(path, dw)
	Expect(result.Title).To(Equal("Croydon"))
	Expect(*result.Description).To(Equal("A true multiplex"))
	Expect(*result.Deactivated).To(Equal(false))
}

func TestInvalidGet(t *testing.T) {
	RegisterTestingT(t)

	path := "invalid/id"
	result := GetByPath(path, dw)
	Expect(result).Should(BeNil())
}

func TestGetByRefs(t *testing.T) {
	RegisterTestingT(t)
	ref := "/events/cinema/vue"

	result := ListByRefs(ref, dw)
	Expect(len(result)).Should(BeNumerically(">=", 1))
}

func TestInvalidGetByRefs(t *testing.T) {
	RegisterTestingT(t)

	ref := "/invalid"
	result := ListByRefs(ref, dw)
	Expect(len(result)).Should(BeNumerically("==", 0))
}

func TestUpdate(t *testing.T) {
	RegisterTestingT(t)
	// Create original POI
	poiPayload := createUpdatePayload()
	path := "/test/xyz007"

	Create(poiPayload, "/test", dw)

	// Create updated POI
	updatedPayload := updatePayload("xyz007")
	// Update POI
	updatedCount := Update(path, &updatedPayload, dw)
	// Check updated count
	Expect(updatedCount).To(Equal(1))
}

func TestInvalidUpdate(t *testing.T) {
	RegisterTestingT(t)
	poiPayload := createUpdatePayload()
	path := "test/invalid"

	Create(poiPayload, "/test", dw)

	// Create updated POI
	updatedPayload := updatePayload("xyz007")
	// Update POI
	updatedCount := Update(path, &updatedPayload, dw)
	// Check updated count
	Expect(updatedCount).To(Equal(0))
}

func TestGetByPaths(t *testing.T) {
	RegisterTestingT(t)
	path := "/test"

	result := ListByPath(path, dw)
	Expect(len(result)).Should(BeNumerically(">=", 1))
}

func TestInvalidGetByPaths(t *testing.T) {
	RegisterTestingT(t)

	path := "invalid"
	result := ListByPath(path, dw)
	Expect(len(result)).Should(BeNumerically("==", 0))
}

func TestDeactivate(t *testing.T) {
	RegisterTestingT(t)

	path := "test/abc123"
	updatedCount := Deactivate(path, dw)
	Expect(updatedCount).To(Equal(1))
}

func TestGetPostDeactivation(t *testing.T) {
	RegisterTestingT(t)
	path := "test/abc123"
	result := GetByPath(path, dw)
	Expect(result).Should(BeNil())
}

func TestInvalidDeactivate(t *testing.T) {
	RegisterTestingT(t)

	path := "test/invalid"
	updatedCount := Deactivate(path, dw)
	Expect(updatedCount).To(Equal(0))
}

func TestDelete(t *testing.T) {
	RegisterTestingT(t)
	path := "/test/abc123"

	updatedCount := Delete(path, dw)
	Expect(updatedCount).To(Equal(1))
}

func TestInvalidDelete(t *testing.T) {
	RegisterTestingT(t)

	path := "/test/invalid"
	updatedCount := Delete(path, dw)
	Expect(updatedCount).To(Equal(0))
}

//Create a new PostPOIPayload
func createPayload() *Poi {

	lat := 51.376495
	long := -0.100594

	title := "Croydon"
	id := "abc123"
	geotype := "point"
	var geoloc []*Geolocation
	geoloc = append(geoloc, &Geolocation{&lat, &long})
	location := "back of the bus station"
	description := "A true multiplex"
	properties := map[string]string{"popcorn": "expensive", "adverts": "too long"}
	var refs []string
	refs = append(refs, "/events/cinema/vue")
	deactivated := false

	poi := Poi{Description: &description, ID: id, Properties: properties, Title: title, Location: &location,
		Geotype: geotype, Geolocation: geoloc, Refs: refs, Deactivated: &deactivated}
	return &poi
}

//Create a payload for update
func createUpdatePayload() *Poi {

	lat := 51.376495
	long := -0.100594

	title := "Croydon"
	id := "xyz007"
	geotype := "point"
	var geoloc []*Geolocation
	geoloc = append(geoloc, &Geolocation{&lat, &long})
	location := "back of the bus station"
	description := "A true multiplex"
	properties := map[string]string{"popcorn": "expensive", "adverts": "too long"}
	var refs []string
	refs = append(refs, "/events/cinema/vue")
	deactivated := false

	poi := Poi{Description: &description, ID: id, Properties: properties, Title: title, Location: &location,
		Geotype: geotype, Geolocation: geoloc, Refs: refs, Deactivated: &deactivated}
	return &poi
}

// Create an POI object for testing updates
func updatePayload(poiId string) Poi {
	// Create data
	title := "New Croydon"
	description := "New description"
	deactivated := true
	properties := map[string]string{"shopping": "whitgift"}
	refs := []string{"/events/NewEvent"}
	// Create POI
	poi := Poi{
		ID:          poiId,
		Title:       title,
		Description: &description,
		Deactivated: &deactivated,
		Properties:  properties,
		Refs:        refs,
	}
	return poi
}
