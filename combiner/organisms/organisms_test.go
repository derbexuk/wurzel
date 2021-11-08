package organisms

import (
	"context"
	"github.com/derbexuk/wurzel/harvester/datawarehouse"
	"fmt"
	//"log"
	"os"
	"testing"

	. "github.com/onsi/gomega"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var dw *datawarehouse.Datawarehouse

/*
  Test Main
*/
func TestMain(m *testing.M) {
	dw = &datawarehouse.Datawarehouse{}
	dw.Open()
	dw.EnsurePathIndex(DB, COLLECTION)

	c := dw.Client.Database(DB).Collection(COLLECTION)
	c.DeleteMany(context.Background(), bson.D{})

	code := m.Run()
	os.Exit(code)
}

func TestOrgCreate(t *testing.T) {
	RegisterTestingT(t)

	org := createPayload()
	err := Create(org, "/test", dw)
	Expect(err).Should(BeNil())

	path := "/test/abc123"
	c := dw.Client.Database(DB).Collection(COLLECTION)
	res := c.FindOne(context.Background(), bson.D{{"path", path}})
	Expect(res).ShouldNot(BeNil())
	// Check retrieved data
	var result map[string]interface{}
	err = res.Decode(&result)
	Expect(err).Should(BeNil())

	Expect(result["path"].(string)).To(Equal(path))
	Expect(result["data"].(map[string]interface{})["properties"].(map[string]interface{})["name"].(string)).To(Equal("CS"))
	Expect(result["data"].(map[string]interface{})["title"].(string)).To(Equal("Org1"))
	Expect(result["data"].(map[string]interface{})["description"].(string)).To(Equal("Org Desc"))
	Expect(result["data"].(map[string]interface{})["deactivated"].(bool)).To(Equal(false))
	Expect(result["data"].(map[string]interface{})["refs"].(primitive.A)[0].(string)).To(Equal("/events/techmeet"))
}

// Test Duplicates
func TestDuplicates(t *testing.T) {
	RegisterTestingT(t)

	org := createPayload()
	path := fmt.Sprintf("%s/%s", "/test", org.ID)
	println(path)
	err := Create(org, "/test", dw)

	// Check results
	Expect(err).ShouldNot(BeNil())
	Expect(err.Error()).Should(ContainSubstring("duplicate key error"))
}

func TestGet(t *testing.T) {
	RegisterTestingT(t)
	path := "/test/abc123"
	result := GetByPath(path, dw)
	Expect(result.Title).To(Equal("Org1"))
	Expect(*result.Description).To(Equal("Org Desc"))
	Expect(*result.Deactivated).To(Equal(false))
}

func TestInvalidGet(t *testing.T) {
	RegisterTestingT(t)
	path := "/invalid"
	result := GetByPath(path, dw)
	Expect(result).Should(BeNil())
}

func TestUpdate(t *testing.T) {
	RegisterTestingT(t)
	// Create original ORG
	path := "/test/xyz007"
	orgPayload := createUpdatePayload()
	Create(orgPayload, "/test", dw)

	// Create updated ORG
	updatedPayload := updatePayload("xyz007")

	// Update ORG
	updatedCount := Update(path, &updatedPayload, dw)

	// Check updated count
	Expect(updatedCount).To(Equal(1))
}

func TestInvalidUpdate(t *testing.T) {
	RegisterTestingT(t)

	path := "/test/invalid"
	orgPayload := createUpdatePayload()
	Create(orgPayload, "/test", dw)

	// Create updated ORG
	updatedPayload := updatePayload("xyz007")
	updatedCount := Update(path, &updatedPayload, dw)
	// Check updated count
	Expect(updatedCount).To(Equal(0))
}

func TestGetByRefs(t *testing.T) {
	RegisterTestingT(t)
	ref := "/events/techmeet"

	result := ListByRefs(ref, dw)
	Expect(len(result)).Should(BeNumerically(">=", 1))
}

func TestInvalidGetByRefs(t *testing.T) {
	RegisterTestingT(t)

	ref := "/invalid"
	result := ListByRefs(ref, dw)
	Expect(len(result)).Should(BeNumerically("==", 0))
}

func TestDeactivate(t *testing.T) {
	RegisterTestingT(t)

	path := "/test/abc123"
	updatedCount := Deactivate(path, dw)
	Expect(updatedCount).To(Equal(1))
}

func TestInvalidDeactivate(t *testing.T) {
	RegisterTestingT(t)

	path := "/test/invalid"
	updatedCount := Deactivate(path, dw)
	Expect(updatedCount).To(Equal(0))
}

//Organism should not be retrieved after deactivation
func TestGetPostDeactivation(t *testing.T) {
	RegisterTestingT(t)
	path := "/test/abc123"
	result := GetByPath(path, dw)
	Expect(result).Should(BeNil())
}

func TestDelete(t *testing.T) {
	RegisterTestingT(t)
	// Create Payload
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

// Create an POI object for testing updates
func updatePayload(orgID string) Organism {
	// Create data
	title := "New Org"
	description := "New description"
	deactivated := true
	properties := map[string]string{"building": "sunley house"}
	refs := []string{"/events/NewEvent"}
	// Create Org
	org := Organism{
		ID:          orgID,
		Title:       title,
		Description: &description,
		Deactivated: &deactivated,
		Properties:  properties,
		Refs:        refs,
	}
	return org
}

func createPayload() *Organism {

	title := "Org1"
	id := "abc123"
	description := "Org Desc"
	properties := map[string]string{"name": "CS", "place": "Croydon"}
	var refs []string
	refs = append(refs, "/events/techmeet")
	deactivated := false

	org := Organism{Description: &description, ID: id, Properties: properties, Title: title, Refs: refs, Deactivated: &deactivated}
	return &org
}

func createUpdatePayload() *Organism {

	title := "Org2"
	id := "xyz007"
	description := "Org Desc2"
	properties := map[string]string{"name": "CS", "place": "Croydon"}
	var refs []string
	refs = append(refs, "/events/techmeet")
	deactivated := false

	org := Organism{Description: &description, ID: id, Properties: properties, Title: title, Refs: refs, Deactivated: &deactivated}
	return &org
}
