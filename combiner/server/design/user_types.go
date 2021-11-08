package design

import (
	. "goa.design/goa/v3/dsl"
)

var GeoPath = Type("GeoPath", func() {
	Attribute("id", String, func() {
		Example("59b287534defd02dace3671b")
	})
	Attribute("title", String, "Path title/name")
	Attribute("type", String, "must be route | point | area")
	Attribute("geopoints", ArrayOf(ArrayOf(Float32)), "path")
})

var PoiPayload = Type("PoiPayload", func() {
	Attribute("id", String, "ID of poi", func() {
		Example("59b287534defd02dace3671b")
	})
	Attribute("title", String, "Poi title/name")
	Attribute("description", String)
	Attribute("deactivated", Boolean)
	Attribute("location", String, "name of location")
	Attribute("geopath", GeoPath)
	Attribute("geojson", String)
	Attribute("properties", MapOf(String, String), "Hash of application specific properties")
	Attribute("refs", ArrayOf(String), "references")
	Attribute("path", String)
	Required("id", "title", "geojson")
})

var EventPayload = Type("EventPayload", func() {
	//	Attribute("path", String, "Path of Event", func() {
	//		Example("59b287534defd02dace3671b")
	//	})
	Attribute("id", String, "ID of Event", func() {
		Example("59b287534defd02dace3671b")
	})
	Attribute("title", String, "Event title/name")
	Attribute("description", String)
	Attribute("deactivated", Boolean)
	Attribute("start", String, "", func() {
		Format("date-time")
		Example("2017-08-22T12:06:30.696Z")
	})
	Attribute("end", String, "", func() {
		Format("date-time")
		Example("2017-08-22T12:06:30.696Z")
	})
	Attribute("properties", MapOf(String, String), "Hash of application specific properties")
	Attribute("refs", ArrayOf(String), "references")
	Attribute("path", String)
	Required("id", "title", "start")
})

var OrganismPayload = Type("OrganismPayload", func() {
	Attribute("id", String, "ID of Organism", func() {
		Example("59b287534defd02dace3671b")
	})
	Attribute("title", String, "Organism title")
	Attribute("description", String)
	Attribute("deactivated", Boolean)
	Attribute("properties", MapOf(String, String), "Hash of application specific properties")
	Attribute("refs", ArrayOf(String), "references")
	Attribute("path", String)
	Required("id", "title")
})
