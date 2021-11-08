package events

import (
	"context"
	"fmt"
	"os"
	"testing"

	. "github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/derbexuk/wurzel/harvester/datawarehouse"
)

var dw *datawarehouse.Datawarehouse

// Test Main
func TestMain(m *testing.M) {
	dw = &datawarehouse.Datawarehouse{}
	dw.Open()
	dw.EnsurePathIndex(DB, COLLECTION)
	c := dw.Client.Database(DB).Collection(COLLECTION)
	c.DeleteMany(context.Background(), bson.D{})
	code := m.Run()
	os.Exit(code)
}

// Test Create
func TestCreate(t *testing.T) {
	RegisterTestingT(t)
	// Create Event
	eventPayload := makePostPayload()
	path := "festivals/croydon"
	id := "001"
	eventPayload.ID = id
	desc := "TestCreate"
	eventPayload.Description = &desc
	Create(path, &eventPayload, dw)
	// Retrieve the Event just created
	var result map[string]interface{}
	fullpath := fmt.Sprintf("/%s/%s", path, id)
	c := dw.Client.Database(DB).Collection(COLLECTION)
	res := c.FindOne(context.Background(), bson.D{{"path", fullpath}})
	Expect(res).ShouldNot(BeNil())
	// Check retrieved data
	err := res.Decode(&result)
	Expect(err).Should(BeNil())
	Expect(result["data"].(map[string]interface{})["title"].(string)).To(Equal("Factory Day"))
	Expect(result["data"].(map[string]interface{})["description"].(string)).To(Equal("TestCreate"))
	Expect(result["data"].(map[string]interface{})["deactivated"].(bool)).To(Equal(false))
	Expect(result["data"].(map[string]interface{})["start"].(string)).To(Equal("2018-01-15 14:00:00"))
	Expect(result["data"].(map[string]interface{})["end"].(string)).To(Equal("2018-01-15 16:00:00"))
	Expect(result["data"].(map[string]interface{})["properties"].(map[string]interface{})["emblem"].(string)).To(Equal("tractor tire"))
	Expect(result["data"].(map[string]interface{})["properties"].(map[string]interface{})["theme"].(string)).To(Equal("Brumm Brumm My Tractor Goes"))
	Expect(result["data"].(map[string]interface{})["refs"].(primitive.A)[0].(string)).To(Equal("/pois/POI01"))
	Expect(result["data"].(map[string]interface{})["refs"].(primitive.A)[1].(string)).To(Equal("/organisms/ORG01"))
}

// Test Create - with optional fields empty
func TestCreateOptional(t *testing.T) {
	RegisterTestingT(t)
	// Create Event
	eventPayload := makePostPayloadOptional()
	id := "001"
	path := "optionals/test"
	Create(path, &eventPayload, dw)
	// Retrieve the Event just created
	var result map[string]interface{}
	fullpath := fmt.Sprintf("/%s/%s", path, id)
	c := dw.Client.Database(DB).Collection(COLLECTION)
	res := c.FindOne(context.Background(), bson.D{{"path", fullpath}})
	Expect(res).ShouldNot(BeNil())
	// Check retrieved data
	err := res.Decode(&result)
	Expect(err).Should(BeNil())
	Expect(result["data"].(map[string]interface{})["title"].(string)).To(Equal("Optional"))
	Expect(result["data"].(map[string]interface{})["description"].(string)).To(Equal("Optional fields test"))
	Expect(result["data"].(map[string]interface{})["deactivated"].(bool)).To(Equal(false))
	Expect(result["data"].(map[string]interface{})["start"].(string)).To(Equal("2018-01-15 14:00:00"))
	Expect(result["data"].(map[string]interface{})["end"].(string)).To(Equal(""))
	properties := result["data"].(map[string]interface{})["properties"].(map[string]interface{})
	propLength := len(properties)
	Expect(propLength).To(Equal(0))
	refs := result["data"].(map[string]interface{})["refs"].(primitive.A)
	refsLength := len(refs)
	Expect(refsLength).To(Equal(0))
}

// Test Get
func TestGetByPath(t *testing.T) {
	RegisterTestingT(t)
	// Create Event
	eventPayload := makePostPayload()
	path := "celebrations"
	id := "002"
	eventPayload.ID = id
	description := "TestGet"
	eventPayload.Description = &description
	Create(path, &eventPayload, dw)
	// Retrieve created event
	fullpath := fmt.Sprintf("%s/%s", path, id)
	result := GetByPath(fullpath, dw)
	// Check results
	Expect(result).ShouldNot(BeNil())
	Expect(result.Title).To(Equal("Factory Day"))
	Expect(*result.Description).To(Equal("TestGet"))
	Expect(*result.Deactivated).To(Equal(false))
	Expect(result.Start).To(Equal("2018-01-15 14:00:00"))
	Expect(*result.End).To(Equal("2018-01-15 16:00:00"))
	Expect(result.Properties["emblem"]).To(Equal("tractor tire"))
	Expect(result.Properties["theme"]).To(Equal("Brumm Brumm My Tractor Goes"))
	Expect(len(result.Refs)).To(Equal(2))
	Expect(result.Refs[0]).To(Equal("/pois/POI01"))
	Expect(result.Refs[1]).To(Equal("/organisms/ORG01"))
}

// Test ListByPath
func TestListByPath(t *testing.T) {
	RegisterTestingT(t)
	// Path
	path := "farming/parties"
	// Create 1st Event
	eventPayload1 := makePostPayload()
	id1 := "001"
	eventPayload1.ID = id1
	title1 := "Tractor Day"
	eventPayload1.Title = title1
	desc1 := "TestGetByPath Event 1"
	eventPayload1.Description = &desc1
	Create(path, &eventPayload1, dw)
	// Create 2nd Event
	eventPayload2 := makePostPayload()
	id2 := "002"
	eventPayload2.ID = id2
	title2 := "Plough Day"
	eventPayload2.Title = title2
	desc2 := "TestGetByPath Event 2"
	eventPayload2.Description = &desc2
	Create(path, &eventPayload2, dw)
	// Retrieve created Events
	results, err := ListByPath(path, dw)
	Expect(err).Should(BeNil())
	// Check number of Events
	length := len(results)
	Expect(length).To(Equal(2))
	// Check 1st Event
	result1 := results[0]
	Expect(result1.Title).To(Equal("Tractor Day"))
	Expect(result1.Start).To(Equal("2018-01-15 14:00:00"))
	Expect(*result1.End).To(Equal("2018-01-15 16:00:00"))

	// Check 2nd Event
	result2 := results[1]
	Expect(result2.Title).To(Equal("Plough Day"))
	Expect(result2.Start).To(Equal("2018-01-15 14:00:00"))
	Expect(*result2.End).To(Equal("2018-01-15 16:00:00"))
}

// Test GetByTimeAndPath - only start time
func TestListByTimeAndPathStartOnly(t *testing.T) {
	RegisterTestingT(t)
	// Path
	path := "time/start"
	// Create 1st Event
	eventPayload1 := makePostPayload()
	id1 := "001"
	eventPayload1.ID = id1
	title1 := "TimeTest1"
	eventPayload1.Title = title1
	desc1 := "TestListByTimeAndPath Event 1"
	eventPayload1.Description = &desc1
	start1 := "2018-01-01 12:00:00"
	eventPayload1.Start = start1
	end1 := ""
	eventPayload1.End = &end1
	Create(path, &eventPayload1, dw)
	// Create 2nd Event
	eventPayload2 := makePostPayload()
	id2 := "002"
	eventPayload2.ID = id2
	title2 := "TimeTest2"
	eventPayload2.Title = title2
	desc2 := "TestListByTimeAndPath Event 2"
	eventPayload2.Description = &desc2
	start2 := "2018-02-02 12:00:00"
	eventPayload2.Start = start2
	end2 := ""
	eventPayload2.End = &end2
	Create(path, &eventPayload2, dw)
	// Select on Time 1
	results, err := ListByTimeAndPath(path, start1, "", dw)
	// Check results
	Expect(err).Should(BeNil())
	Expect(results).ShouldNot(BeNil())
	Expect(len(results)).To(Equal(2))
	// Check 1st Event
	result1 := results[0]
	Expect(result1.Title).To(Equal("TimeTest1"))
	// Check 2nd Event
	result2 := results[1]
	Expect(result2.Title).To(Equal("TimeTest2"))
	// Select on Time 2
	results, err = ListByTimeAndPath(path, start2, "", dw)
	// Check results
	Expect(err).Should(BeNil())
	Expect(results).ShouldNot(BeNil())
	Expect(len(results)).To(Equal(1))
	// Check 1st Event
	result1 = results[0]
	Expect(result1.Title).To(Equal("TimeTest2"))
}

// Test GetByTimeAndPath - start and end time
func TestListByTimeAndPathStartAndEnd(t *testing.T) {
	RegisterTestingT(t)
	// Path
	path := "time/startandend"
	// Create 1st Event
	eventPayload1 := makePostPayload()
	id1 := "010"
	eventPayload1.ID = id1
	title1 := "TimeTest1"
	eventPayload1.Title = title1
	desc1 := "TestListByTimeAndPath Event 1"
	eventPayload1.Description = &desc1
	start1 := "2018-01-01 12:00:00"
	eventPayload1.Start = start1
	end1 := "2018-07-01 12:00:00"
	eventPayload1.End = &end1
	Create(path, &eventPayload1, dw)
	// Create 2nd Event
	eventPayload2 := makePostPayload()
	id2 := "020"
	eventPayload2.ID = id2
	title2 := "TimeTest2"
	eventPayload2.Title = title2
	desc2 := "TestListByTimeAndPath Event 2"
	eventPayload2.Description = &desc2
	start2 := "2018-02-02 12:00:00"
	eventPayload2.Start = start2
	end2 := "2018-12-31 12:00:00"
	eventPayload2.End = &end2
	Create(path, &eventPayload2, dw)
	// Select on start 1 and end 1 - between 2018-01-01 12:00:00 and 2018-07-01 12:00:00 inclusive
	results, err := ListByTimeAndPath(path, start1, end1, dw)
	// Check results
	Expect(err).Should(BeNil())
	Expect(results).ShouldNot(BeNil())
	Expect(len(results)).To(Equal(2))
	// Check 1st Event
	result1 := results[0]
	Expect(result1.Title).To(Equal("TimeTest1"))
	// Check 2nd Event
	result2 := results[1]
	Expect(result2.Title).To(Equal("TimeTest2"))
	// Select on start 2 and end 2 - between 2018-02-02 12:00:00 and 2018-12-31 12:00:00 inclusive
	results, err = ListByTimeAndPath(path, start2, end2, dw)
	// Check results
	Expect(err).Should(BeNil())
	Expect(results).ShouldNot(BeNil())
	Expect(len(results)).To(Equal(1))
	// Check 1st Event
	result1 = results[0]
	Expect(result1.Title).To(Equal("TimeTest2"))
	// Select on start 2 and end 1 - between 2018-02-02 12:00:00 and 2018-07-01 12:00:00 inclusive
	results, err = ListByTimeAndPath(path, "2018-01-02 12:00:00", "2018-01-02 12:00:00", dw)
	// Check results
	Expect(err).Should(BeNil())
	Expect(len(results)).To(Equal(0))
}

// Test GetByTimeAndPath - check time difference
// Two Events on same day but at different times
func TestListByTimeAndPathByTime(t *testing.T) {
	RegisterTestingT(t)
	// Path
	path := "time/time"
	// Create 1st Event
	eventPayload1 := makePostPayload()
	id1 := "001"
	eventPayload1.ID = id1
	title1 := "TimeTest1"
	eventPayload1.Title = title1
	desc1 := "TestListByTimeAndPath Event 1"
	eventPayload1.Description = &desc1
	start1 := "2018-01-01 12:00:00"
	eventPayload1.Start = start1
	end1 := ""
	eventPayload1.End = &end1
	Create(path, &eventPayload1, dw)
	// Create 2nd Event
	eventPayload2 := makePostPayload()
	id2 := "002"
	eventPayload2.ID = id2
	title2 := "TimeTest2"
	eventPayload2.Title = title2
	desc2 := "TestListByTimeAndPath Event 2"
	eventPayload2.Description = &desc2
	start2 := "2018-01-01 14:00:00"
	eventPayload2.Start = start2
	end2 := ""
	eventPayload2.End = &end2
	Create(path, &eventPayload2, dw)
	// Select on Time 1
	testTime1 := "2018-01-01 09:00:00"
	results, err := ListByTimeAndPath(path, testTime1, "", dw)
	// Check results
	Expect(err).Should(BeNil())
	Expect(results).ShouldNot(BeNil())
	Expect(len(results)).To(Equal(2))
	// Check 1st Event
	result1 := results[0]
	Expect(result1.Title).To(Equal("TimeTest1"))
	// Check 2nd Event
	result2 := results[1]
	Expect(result2.Title).To(Equal("TimeTest2"))
	// Select on Time 2
	testTime2 := "2018-01-01 13:00:00"
	results, err = ListByTimeAndPath(path, testTime2, "", dw)
	// Check results
	Expect(err).Should(BeNil())
	Expect(results).ShouldNot(BeNil())
	Expect(len(results)).To(Equal(1))
	// Check 1st Event
	result1 = results[0]
	Expect(result1.Title).To(Equal("TimeTest2"))
}

// Test Update
func TestUpdate(t *testing.T) {
	RegisterTestingT(t)
	// Create Event
	eventPayload := makePostPayload()
	id := "100"
	eventPayload.ID = id
	path := "updates"
	Create(path, &eventPayload, dw)
	// Build full path
	fullpath := fmt.Sprintf("%s/%s", path, id)
	// Update event
	updatePayload := makeEventPayload(eventPayload)
	newTitle := "UPDATED TITLE"
	newDescription := "UPDATED DESCRIPTION"
	updatePayload.Title = newTitle
	updatePayload.Description = &newDescription
	updateCount := Update(fullpath, &updatePayload, dw)
	Expect(updateCount).To(Equal(1))
	// Retrieve the Event just updated
	c := dw.Client.Database(DB).Collection(COLLECTION)
	res := c.FindOne(context.Background(), bson.D{{"path", "/" + fullpath}})
	Expect(res).ShouldNot(BeNil())
	// Check retrieved data
	var result map[string]interface{}
	err := res.Decode(&result)
	Expect(err).Should(BeNil())
	Expect(result["data"].(map[string]interface{})["title"].(string)).To(Equal(newTitle))
	Expect(result["data"].(map[string]interface{})["description"].(string)).To(Equal(newDescription))
	Expect(result["data"].(map[string]interface{})["deactivated"].(bool)).To(Equal(false))
	Expect(result["data"].(map[string]interface{})["start"].(string)).To(Equal("2018-01-15 14:00:00"))
	Expect(result["data"].(map[string]interface{})["end"].(string)).To(Equal("2018-01-15 16:00:00"))
	Expect(result["data"].(map[string]interface{})["properties"].(map[string]interface{})["emblem"].(string)).To(Equal("tractor tire"))
	Expect(result["data"].(map[string]interface{})["properties"].(map[string]interface{})["theme"].(string)).To(Equal("Brumm Brumm My Tractor Goes"))
	Expect(result["data"].(map[string]interface{})["refs"].(primitive.A)[0].(string)).To(Equal("/pois/POI01"))
	Expect(result["data"].(map[string]interface{})["refs"].(primitive.A)[1].(string)).To(Equal("/organisms/ORG01"))
}

// Test Delete
func TestDelete(t *testing.T) {
	RegisterTestingT(t)
	// Path
	path := "removals"
	// Create original Event
	eventPayload := makePostPayload()
	id := "200"
	eventPayload.ID = id
	Create(path, &eventPayload, dw)
	// Build path
	fullpath := fmt.Sprintf("%s/%s", path, id)
	// Delete event
	updateCount := Delete(fullpath, dw)
	Expect(updateCount).To(Equal(1))
	// Attempt to retrieve the Event just deleted
	c := dw.Client.Database(DB).Collection(COLLECTION)
	res := c.FindOne(context.Background(), bson.D{{"path", fullpath}})
	Expect(res).ShouldNot(BeNil())
}

// Test Deactivate
func TestDeactivate(t *testing.T) {
	RegisterTestingT(t)
	// Path
	path := "deactivate"
	// Create Event
	eventPayload := makePostPayload()
	id := "300"
	eventPayload.ID = id
	description := "TestDeactivate"
	eventPayload.Description = &description
	Create(path, &eventPayload, dw)
	// Build path
	fullpath := fmt.Sprintf("%s/%s", path, id)
	// Deactivate event
	updateCount := Deactivate(fullpath, dw)
	Expect(updateCount).To(Equal(1))
	// Retrieve the Event just updated
	c := dw.Client.Database(DB).Collection(COLLECTION)
	res := c.FindOne(context.Background(), bson.D{{"path", "/" + fullpath}})
	Expect(res).ShouldNot(BeNil())
	// Check retrieved data
	var result map[string]interface{}
	err := res.Decode(&result)
	Expect(err).Should(BeNil())
	Expect(result["data"].(map[string]interface{})["title"].(string)).To(Equal("Factory Day"))
	Expect(result["data"].(map[string]interface{})["description"].(string)).To(Equal("TestDeactivate"))
	Expect(result["data"].(map[string]interface{})["deactivated"].(bool)).To(Equal(true))
	Expect(result["data"].(map[string]interface{})["start"].(string)).To(Equal("2018-01-15 14:00:00"))
	Expect(result["data"].(map[string]interface{})["end"].(string)).To(Equal("2018-01-15 16:00:00"))
	Expect(result["data"].(map[string]interface{})["properties"].(map[string]interface{})["emblem"].(string)).To(Equal("tractor tire"))
	Expect(result["data"].(map[string]interface{})["properties"].(map[string]interface{})["theme"].(string)).To(Equal("Brumm Brumm My Tractor Goes"))
	Expect(result["data"].(map[string]interface{})["refs"].(primitive.A)[0].(string)).To(Equal("/pois/POI01"))
	Expect(result["data"].(map[string]interface{})["refs"].(primitive.A)[1].(string)).To(Equal("/organisms/ORG01"))
}

func TestGetPostDeactivation(t *testing.T) {
	RegisterTestingT(t)
	path := "/deactivate/300"
	result := GetByPath(path, dw)
	Expect(result).Should(BeNil())
}

// Test Duplicates
func TestDuplicates(t *testing.T) {
	RegisterTestingT(t)
	// Path
	path := "duplicates/test"
	// Create 1st Event
	eventPayload1 := makePostPayload()
	id1 := "400"
	eventPayload1.ID = id1
	title1 := "Duplicate 1"
	eventPayload1.Title = title1
	desc1 := "Test Duplicates Event 1"
	eventPayload1.Description = &desc1
	err := Create(path, &eventPayload1, dw)
	Expect(err).Should(BeNil())
	// Create 2nd Event - should Panic
	eventPayload2 := makePostPayload()
	id2 := "400"
	eventPayload2.ID = id2
	title2 := "Duplicate 2"
	eventPayload2.Title = title2
	desc2 := "Test Duplicates Event 2"
	eventPayload2.Description = &desc2
	err = Create(path, &eventPayload2, dw)
	// Check results
	Expect(err).ShouldNot(BeNil())
	Expect(err.Error()).Should(ContainSubstring("duplicate key error"))
}

// Test TimeSearch - sorted in ascending order
func TestTimeSearchAscending(t *testing.T) {
	RegisterTestingT(t)
	// Path
	path := "timesearch/ascending"
	// Create 1st Event
	eventPayload1 := makePostPayload()
	id1 := "001"
	eventPayload1.ID = id1
	title1 := "TimeSearch1"
	eventPayload1.Title = title1
	desc1 := "TimeSearch Event 1"
	eventPayload1.Description = &desc1
	start1 := "2018-01-01T14:00:00.696Z"
	eventPayload1.Start = start1
	end1 := "2018-01-15T16:00:00.696Z"
	eventPayload1.End = &end1
	Create(path, &eventPayload1, dw)
	// Create 2nd Event
	eventPayload2 := makePostPayload()
	id2 := "002"
	eventPayload2.ID = id2
	title2 := "TimeSearch2"
	eventPayload2.Title = title2
	desc2 := "TimeSearch Event 2"
	eventPayload2.Description = &desc2
	start2 := "2018-01-03T14:00:00.696Z"
	eventPayload2.Start = start2
	end2 := "2018-01-15T16:00:00.696Z"
	eventPayload2.End = &end2
	Create(path, &eventPayload2, dw)
	// Create 3rd Event
	eventPayload3 := makePostPayload()
	id3 := "003"
	eventPayload3.ID = id3
	title3 := "TimeSearch3"
	eventPayload3.Title = title3
	desc3 := "TimeSearch Event 3"
	eventPayload3.Description = &desc3
	start3 := "2018-01-02T14:00:00.696Z"
	eventPayload3.Start = start3
	end3 := "2018-01-15T16:00:00.696Z"
	eventPayload3.End = &end3
	Create(path, &eventPayload3, dw)
	// Retrieve created Events
	// Params: (path string, start string, end string, order string, limit int, dw *datawarehouse.Datawarehouse)
	results, err := TimeSearch(path, "2018-01-01T14:00:00.696Z", "2018-01-15T16:00:00.696Z", "a", 100, dw)
	// Check for errors
	Expect(err).Should(BeNil())
	// Check number of Events
	length := len(results)
	Expect(length).To(Equal(3))
	// SORTED IN ASCENDING ORDER THE RESULTS SHOULD BE: TimeSearch1, TimeSearch3, TimeSearch2
	// Check 1st Event
	result1 := results[0]
	Expect(result1.Title).To(Equal("TimeSearch1"))
	Expect(result1.Start).To(Equal("2018-01-01T14:00:00.696Z"))
	Expect(*result1.End).To(Equal("2018-01-15T16:00:00.696Z"))

	// Check 2nd Event
	result2 := results[1]
	Expect(result2.Title).To(Equal("TimeSearch3"))
	Expect(result2.Start).To(Equal("2018-01-02T14:00:00.696Z"))
	Expect(*result2.End).To(Equal("2018-01-15T16:00:00.696Z"))

	// Check 3rd Event
	result3 := results[2]
	Expect(result3.Title).To(Equal("TimeSearch2"))
	Expect(result3.Start).To(Equal("2018-01-03T14:00:00.696Z"))
	Expect(*result3.End).To(Equal("2018-01-15T16:00:00.696Z"))
}

// Test TimeSearch - sorted in descending order
func TestTimeSearchDescending(t *testing.T) {
	RegisterTestingT(t)
	// Path
	path := "timesearch/descending"
	// Create 1st Event
	eventPayload1 := makePostPayload()
	id1 := "001"
	eventPayload1.ID = id1
	title1 := "TimeSearch1"
	eventPayload1.Title = title1
	desc1 := "TimeSearch Event 1"
	eventPayload1.Description = &desc1
	start1 := "2018-01-01T14:00:00.696Z"
	eventPayload1.Start = start1
	end1 := "2018-01-15T16:00:00.696Z"
	eventPayload1.End = &end1
	Create(path, &eventPayload1, dw)
	// Create 2nd Event
	eventPayload2 := makePostPayload()
	id2 := "002"
	eventPayload2.ID = id2
	title2 := "TimeSearch2"
	eventPayload2.Title = title2
	desc2 := "TimeSearch Event 2"
	eventPayload2.Description = &desc2
	start2 := "2018-01-03T14:00:00.696Z"
	eventPayload2.Start = start2
	end2 := "2018-01-15T16:00:00.696Z"
	eventPayload2.End = &end2
	Create(path, &eventPayload2, dw)
	// Create 3rd Event
	eventPayload3 := makePostPayload()
	id3 := "003"
	eventPayload3.ID = id3
	title3 := "TimeSearch3"
	eventPayload3.Title = title3
	desc3 := "TimeSearch Event 3"
	eventPayload3.Description = &desc3
	start3 := "2018-01-02T14:00:00.696Z"
	eventPayload3.Start = start3
	end3 := "2018-01-15T16:00:00.696Z"
	eventPayload3.End = &end3
	Create(path, &eventPayload3, dw)
	// Retrieve created Events
	// Params: (path string, start string, end string, order string, limit int, dw *datawarehouse.Datawarehouse)
	results, err := TimeSearch(path, "2018-01-01T14:00:00.696Z", "2018-01-15T16:00:00.696Z", "d", 100, dw)
	// Check for errors
	Expect(err).Should(BeNil())
	// Check number of Events
	length := len(results)
	Expect(length).To(Equal(3))
	// SORTED IN DESCENDING ORDER THE RESULTS SHOULD BE: TimeSearch2, TimeSearch3, TimeSearch1
	// Check 3rd Event
	result1 := results[2]
	Expect(result1.Title).To(Equal("TimeSearch1"))
	Expect(result1.Start).To(Equal("2018-01-01T14:00:00.696Z"))
	Expect(*result1.End).To(Equal("2018-01-15T16:00:00.696Z"))

	// Check 2nd Event
	result2 := results[1]
	Expect(result2.Title).To(Equal("TimeSearch3"))
	Expect(result2.Start).To(Equal("2018-01-02T14:00:00.696Z"))
	Expect(*result2.End).To(Equal("2018-01-15T16:00:00.696Z"))
	// Check 1st Event
	result3 := results[0]
	Expect(result3.Title).To(Equal("TimeSearch2"))
	Expect(result3.Start).To(Equal("2018-01-03T14:00:00.696Z"))
	Expect(*result3.End).To(Equal("2018-01-15T16:00:00.696Z"))
}

// Test TimeSearch - ascending order limit = 1
func TestTimeSearchLimit(t *testing.T) {
	RegisterTestingT(t)
	// Path
	path := "timesearch/limit"
	// Create 1st Event
	eventPayload1 := makePostPayload()
	id1 := "001"
	eventPayload1.ID = id1
	title1 := "TimeSearch1"
	eventPayload1.Title = title1
	desc1 := "TimeSearch Event 1"
	eventPayload1.Description = &desc1
	start1 := "2018-01-01T14:00:00.696Z"
	eventPayload1.Start = start1
	end1 := "2018-01-15T16:00:00.696Z"
	eventPayload1.End = &end1
	Create(path, &eventPayload1, dw)
	// Create 2nd Event
	eventPayload2 := makePostPayload()
	id2 := "002"
	eventPayload2.ID = id2
	title2 := "TimeSearch2"
	eventPayload2.Title = title2
	desc2 := "TimeSearch Event 2"
	eventPayload2.Description = &desc2
	start2 := "2018-01-03T14:00:00.696Z"
	eventPayload2.Start = start2
	end2 := "2018-01-15T16:00:00.696Z"
	eventPayload2.End = &end2
	Create(path, &eventPayload2, dw)
	// Create 3rd Event
	eventPayload3 := makePostPayload()
	id3 := "003"
	eventPayload3.ID = id3
	title3 := "TimeSearch3"
	eventPayload3.Title = title3
	desc3 := "TimeSearch Event 3"
	eventPayload3.Description = &desc3
	start3 := "2018-01-02T14:00:00.696Z"
	eventPayload3.Start = start3
	end3 := "2018-01-15T16:00:00.696Z"
	eventPayload3.End = &end3
	Create(path, &eventPayload3, dw)
	// Retrieve created Events
	// Params: (path string, start string, end string, order string, limit int, dw *datawarehouse.Datawarehouse)
	results, err := TimeSearch(path, "2018-01-01T14:00:00.696Z", "2018-01-15T16:00:00.696Z", "a", 1, dw)
	// Check for errors
	Expect(err).Should(BeNil())
	// Check number of Events
	length := len(results)
	Expect(length).To(Equal(1))
	// SORTED IN ASCENDING ORDER THE RESULTS SHOULD BE: TimeSearch1, TimeSearch3, TimeSearch2 - only TimeSearch1 should be returned
	// Check 1st Event
	result1 := results[0]
	Expect(result1.Title).To(Equal("TimeSearch1"))
	Expect(result1.Start).To(Equal("2018-01-01T14:00:00.696Z"))
	Expect(*result1.End).To(Equal("2018-01-15T16:00:00.696Z"))
}

// Test TimeSearch - outside time range
func TestTimeSearchBadRange(t *testing.T) {
	RegisterTestingT(t)
	// Path
	path := "timesearch/badrange"
	// Create 1st Event
	eventPayload1 := makePostPayload()
	id1 := "001"
	eventPayload1.ID = id1
	title1 := "TimeSearch1"
	eventPayload1.Title = title1
	desc1 := "TimeSearch Event 1"
	eventPayload1.Description = &desc1
	start1 := "2018-01-01T14:00:00.696Z"
	eventPayload1.Start = start1
	end1 := "2018-01-15T16:00:00.696Z"
	eventPayload1.End = &end1
	Create(path, &eventPayload1, dw)
	// Create 2nd Event
	eventPayload2 := makePostPayload()
	id2 := "002"
	eventPayload2.ID = id2
	title2 := "TimeSearch2"
	eventPayload2.Title = title2
	desc2 := "TimeSearch Event 2"
	eventPayload2.Description = &desc2
	start2 := "2018-01-03T14:00:00.696Z"
	eventPayload2.Start = start2
	end2 := "2018-01-15T16:00:00.696Z"
	eventPayload2.End = &end2
	Create(path, &eventPayload2, dw)
	// Create 3rd Event
	eventPayload3 := makePostPayload()
	id3 := "003"
	eventPayload3.ID = id3
	title3 := "TimeSearch3"
	eventPayload3.Title = title3
	desc3 := "TimeSearch Event 3"
	eventPayload3.Description = &desc3
	start3 := "2018-01-02T14:00:00.696Z"
	eventPayload3.Start = start3
	end3 := "2018-01-15T16:00:00.696Z"
	eventPayload3.End = &end3
	Create(path, &eventPayload3, dw)
	// Retrieve created Events
	// Params: (path string, start string, end string, order string, limit int, dw *datawarehouse.Datawarehouse)
	results, err := TimeSearch(path, "2017-01-01T14:00:00.696Z", "2017-01-15T16:00:00.696Z", "a", 1, dw)
	// Check for errors
	Expect(err).Should(BeNil())
	// Check number of Events
	length := len(results)
	Expect(length).To(Equal(0))
}

// Test TimeSearch - nothing found
func TestTimeSearchNotFound(t *testing.T) {
	RegisterTestingT(t)
	// Path
	path := "timesearch/notfound"
	// Attempt to retrieve non-existent Events
	// Params: (path string, start string, end string, order string, limit int, dw *datawarehouse.Datawarehouse)
	results, err := TimeSearch(path, "2018-01-01T14:00:00.696Z", "2018-01-15T16:00:00.696Z", "a", 1, dw)
	// Check for errors
	Expect(err).Should(BeNil())
	// Check number of Events
	length := len(results)
	Expect(length).To(Equal(0))
}

// Wrapper finction for catching Panic (Cannot have input or output parameters when catching Panic)
func makeDuplicate() {
	// Path
	path := "duplicates/test"
	// Create 2nd Event
	eventPayload2 := makePostPayload()
	id2 := "400" // same id as before
	eventPayload2.ID = id2
	title2 := "Duplicate 2"
	eventPayload2.Title = title2
	desc2 := "Test Duplicates Event 2"
	eventPayload2.Description = &desc2
	Create(path, &eventPayload2, dw)
}

// Create an Event object for test consumption
func makePostPayload() Event {
	// Create data
	id := "001"
	title := "Factory Day"
	description := "A Celebration Of Agricultural Tractors"
	deactivated := false
	start := "2018-01-15 14:00:00"
	end := "2018-01-15 16:00:00"
	properties := map[string]string{"emblem": "tractor tire", "theme": "Brumm Brumm My Tractor Goes"}
	refs := []string{"/pois/POI01", "/organisms/ORG01"}
	// Create Event
	event := Event{
		//		Path:        &path,
		ID:          id,
		Title:       title,
		Description: &description,
		Deactivated: &deactivated,
		Start:       start,
		End:         &end,
		Properties:  properties,
		Refs:        refs,
	}
	return event
}

// Create an Event object for test consumption
func makePostPayloadOptional() Event {
	// Create data
	id := "001"
	title := "Optional"
	description := "Optional fields test"
	deactivated := false
	start := "2018-01-15 14:00:00"
	end := ""
	properties := make(map[string]string)
	refs := make([]string, 0)
	// Create Event
	event := Event{
		ID:          id,
		Title:       title,
		Description: &description,
		Deactivated: &deactivated,
		Start:       start,
		End:         &end,
		Properties:  properties,
		Refs:        refs,
	}
	return event
}

// Create an Event object from a PostEvent object
func makeEventPayload(oldEvent Event) Event {
	event := Event{
		ID:          oldEvent.ID,
		Title:       oldEvent.Title,
		Description: oldEvent.Description,
		Deactivated: oldEvent.Deactivated,
		Start:       oldEvent.Start,
		End:         oldEvent.End,
		Properties:  oldEvent.Properties,
		Refs:        oldEvent.Refs,
	}
	return event
}

func printPayload(p Event) {
	fmt.Println("### EventPayload ###")
	fmt.Println("ID:", p.ID)
	fmt.Println("Title:", p.Title)
	fmt.Println("Description:", p.Description)
	fmt.Println("Deactivated:", p.Deactivated)
	fmt.Println("Start:", p.Start)
	fmt.Println("End:", p.End)
	fmt.Println("Properties:", p.Properties)
	fmt.Println("Refs:", p.Refs)
}
