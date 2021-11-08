package design

import (
	. "goa.design/goa/v3/dsl"
)

var Geolocation = Type("Geolocation", func() {
		Attribute("Latitude", Float64)
		Attribute("Longitude", Float64)
})

var PoiMedia = ResultType("application/vnd.cspace.poi+json", func() {
	Description("A Point of Interest")
	Reference(PoiPayload)
	Attributes(func() { // Attributes define the media type shape.
		Attribute("id")
		Attribute("title")
		Attribute("description")
		Attribute("deactivated")
		Attribute("location")
		Attribute("geojson")
		Attribute("properties")
		Attribute("refs", ArrayOf(String), "refs")
		Attribute("path")
		Required("id", "title", "geojson")
	})
	View("default", func() { // View defines a rendering of the media type.
		Attribute("id")
		Attribute("title")
		Attribute("description")
		Attribute("deactivated")
		Attribute("location")
		Attribute("geojson")
		Attribute("properties")
		Attribute("refs")
		Attribute("path")
	})
	View("summary", func() {
		Attribute("id")
		Attribute("title")
		Attribute("geojson")
		Attribute("path")
	})
})

var PoiSummaryMedia = ResultType("application/vnd.cspace.pois+json", func() {
	Attribute("pois", CollectionOf(PoiMedia), func() {
		Description("A list of pois")
		View("summary")
	})
	View("default", func() {
		Attribute("pois")
	})
})

var EventMedia = ResultType("application/vnd.cspace.event+json", func() {
	Description("An Event")
	Reference(EventPayload)
	Attributes(func() { // Attributes define the media type shape.
		Attribute("id")
		Attribute("title")
		Attribute("description")
		Attribute("deactivated")
		Attribute("start")
		Attribute("end")
		Attribute("properties")
		Attribute("refs", ArrayOf(String), "refs")
		Attribute("path")
	})
	View("default", func() { // View defines a rendering of the media type.
		Attribute("id")
		Attribute("title")
		Attribute("description")
		Attribute("deactivated")
		Attribute("start")
		Attribute("end")
		Attribute("properties")
		Attribute("refs")
		Attribute("path")
	})
	View("summary", func() { // View defines a rendering of the media type.
		Attribute("id")
		Attribute("title")
		Attribute("start")
		Attribute("end")
		Attribute("path")
	})
})

var EventSummaryMedia = ResultType("application/vnd.cspace.events+json", func() {
	Attribute("events", CollectionOf(EventMedia), func() {
		Description("An list of events")
		View("summary")
	})
	View("default", func() {
		Attribute("events")
	})
})

var OrganismMedia = ResultType("application/vnd.cspace.organism+json", func() {
	Description("An Organism")
	Reference(OrganismPayload)
	Attributes(func() { // Attributes define the media type shape.
		Attribute("id")
		Attribute("title")
		Attribute("description")
		Attribute("deactivated")
		Attribute("properties")
		Attribute("refs", ArrayOf(String), "refs")
		Attribute("path")
	})
	View("default", func() { // View defines a rendering of the media type.
		Attribute("id")
		Attribute("title")
		Attribute("description")
		Attribute("deactivated")
		Attribute("properties")
		Attribute("refs")
		Attribute("path")
	})
	View("summary", func() { // View defines a rendering of the media type.
		Attribute("id")
		Attribute("title")
		Attribute("path")
	})
})

var OrganismSummaryMedia = ResultType("application/vnd.cspace.organisms+json", func() {
	Attribute("organisms", CollectionOf(OrganismMedia), func() {
		Description("A list of organisms")
		View("summary")
	})
	View("default", func() {
		Attribute("organisms")
	})
})
