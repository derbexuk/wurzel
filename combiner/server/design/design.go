package design

import (
	. "goa.design/goa/v3/dsl"
	cors "goa.design/plugins/v3/cors/dsl"
)

var _ = API("PoiEventApi", func() {
	Title("The Place of Interest, Event, and Organism API")
	Description("Generic Poi, Event, and Organism API server")

	Server("PoiEventApi", func() {
		Services("events", "organisms", "pois", "upload")
		Host("production", func() {
			URI("http://0.0.0.0:8081")
		})
	})
	HTTP(func() {
		Path("/things")
	})

	cors.Origin("*", func() {
		cors.Headers("Authorization", "Origin", "X-Requested-With", "Content-Type", "Accept")
		cors.Methods("POST", "GET", "PUT", "DELETE", "OPTIONS")
		cors.Expose("X-Time")
		cors.MaxAge(600)
		cors.Credentials()
	})
})

// Generic upload
var _ = Service("upload", func() {
	HTTP(func() {
		Path("/upload/")
	})
	Error("bad_req")
	Error("internal_error", func() {
                Fault()
        })

	// Poi - Upload (bulk Create)
	Method("fetch", func() {
		Description("Import the Google Doc For Generic Upload")
		Payload(func() {
			Attribute("configFile", String, "Config file name", func() {
				Example("test.yaml")
			})
		})
		HTTP(func() {
			POST("/fetch/{configFile}")
			Response(StatusOK)
			Response("bad_req", StatusBadRequest)
		})
	})

	Method("csv", func() {
		Description("Upload CSV file")
		Payload(func() {
			Attribute("configFile", String, "Config file name", func() {
				Example("test.yaml")
			})
			Attribute("resource", String, "Service we are uploading", func() {
				Example("bizFeed.txt")
			})
		})
		HTTP(func() {
			POST("/csv/{configFile}/{resource}")
			Response(StatusOK)
			Response("bad_req", StatusBadRequest)
		})
	})
})

// Pois - Places of Interest
var _ = Service("pois", func() {
	HTTP(func() {
		Path("/pois/")
	})
	Error("bad_req")
	Error("not_found")
	Error("un_auth")
	Error("internal_error", func() {
                Fault()
        })

	// Poi - Post (Create)
	Method("post", func() {
		Description("Add a POI")
		Payload(func() {
			Attribute("path", String, "the path of the POI")
			Attribute("pois", ArrayOf(PoiPayload))
			Required("path")
		})
		HTTP(func() {
			POST("/{*path}")
			Response(StatusOK)
			Response("un_auth", StatusUnauthorized)
			Response("bad_req", StatusBadRequest)
		})
	})

	// Poi - Show (Read)
	Method("show", func() {
		Description("Show a POI")
		Payload(func() {
			Attribute("path", String, "the path of the POI")
			Required("path")
		})
		Result(PoiMedia)
		HTTP(func() {
			GET("/path/{*path}")
			Response("not_found", StatusNotFound)
			Response(StatusOK)
		})
	})

	// Poi - List (Read multiple)
	Method("ListByPath", func() {
		Description("List POIs for a path")
		Payload(func() {
			Attribute("path", String, "the path of the POI")
			Required("path")
		})
		Result(CollectionOf(PoiMedia))
		HTTP(func() {
			GET("/paths/{*path}")
			Response(StatusOK)
			Response("not_found", StatusNotFound)
		})
	})

	Method("ListByReference", func() {
		Description("List POIs for a path")
		Payload(func() {
			Attribute("path", String, "the path of the POI")
			Required("path")
		})
		Result(CollectionOf(PoiMedia))
		HTTP(func() {
			GET("/refs/{*path}")
			Response(StatusOK)
			Response("not_found", StatusNotFound)
		})
	})

	// Poi - Update
	Method("update", func() {
		Description("Update a POI")
		Payload(func() {
			Attribute("path", String, "the path of the POI")
			Attribute("poi", PoiPayload)
			Required("path")
		})
		HTTP(func() {
			PUT("/update/{*path}")
			Response(StatusOK)
			Response("bad_req", StatusBadRequest)
		})
	})

	Method("deactivate", func() {
		Description("Update a POI")
		Payload(func() {
			Attribute("path", String, "the path of the POI")
			Required("path")
		})
		HTTP(func() {
			PUT("/deactivate/{*path}")
			Response(StatusOK)
			Response("not_found", StatusNotFound)
		})
	})

	// Poi - Delete
	Method("delete", func() {
		Description("Delete a POI")
		Payload(func() {
			Attribute("path", String, "the path of the POI")
			Required("path")
		})

		HTTP(func() {
			DELETE("/delete/{*path}")
			Response(StatusOK)
			Response("not_found", StatusNotFound)
		})
	})
})

// Events
var _ = Service("events", func() {
	HTTP(func() {
		Path("/events/")
	})
	Error("bad_req")
	Error("not_found")
	Error("un_auth")
	Error("internal_error", func() {
                Fault()
        })

	// Event - Post (Create)
	Method("post", func() {
		Description("Add an event")
		Payload(func() {
			Attribute("path", String, "the path of the Event", func() {
				Example("/croydon/purley/59b287534defd02dace3671b")
			})
			Attribute("events", ArrayOf(EventPayload))
			Required("path")
		})
		HTTP(func() {
			POST("/{*path}")
			Response(StatusOK)
			Response("bad_req", StatusBadRequest)
		})
	})

	// Event - Show (Read)
	Method("show", func() {
		Description("Show an event in full")
		Payload(func() {
			Attribute("path", String, "the path of the Event", func() {
				Example("/croydon/purley/59b287534defd02dace3671b")
			})
			Required("path")
		})
		Result(EventMedia)
		HTTP(func() {
			GET("/path/{*path}")
			Response(StatusOK)
			Response("not_found", StatusNotFound)
		})
	})

	// Event - Show Earliest
	Method("TimeSearch", func() {
		Description("Show events according to search values")
		Payload(func() {
			Attribute("start", String, "Start", func() {
				Format("date-time")
				Example("2017-08-22T12:06:30.696Z")
			})
			Attribute("end", String, "End", func() {
				Format("date-time")
				Example("2017-08-22T12:06:30.696Z")
			})
			Attribute("order", String, "Order", func() {
				Example("A")
			})
			Attribute("limit", Int, "Limit")
			Attribute("path", String, "Path", func() {
				Example("/croydon/purley/59b287534defd02dace3671b")
			})
			Required("start")
			Required("path")
		})
		Result(CollectionOf(EventMedia))
		HTTP(func() {
			GET("/start/{start}/end/{end}/order/{order}/limit/{limit}/path/{*path}")
			Response(StatusOK)
			Response("not_found", StatusNotFound)
			Response("bad_req", StatusBadRequest)
		})
	})

	// Event - List ByTimeAndPath(Read multiple)
	Method("ListByTimeAndPath", func() {
		Description("List Events for a path")
		Payload(func() {
			Attribute("start", String, "Start", func() {
				Format("date-time")
				Example("2017-08-22T12:06:30.696Z")
			})
			Attribute("end", String, "End", func() {
				Format("date-time")
				Example("2017-08-22T12:06:30.696Z")
			})
			Attribute("path", String, "Path", func() {
				Example("/croydon/purley/59b287534defd02dace3671b")
			})
			Required("start")
			Required("path")
		})
		Result(CollectionOf(EventMedia))
		HTTP(func() {
			GET("/start/{start}/end/{end}/path/{*path}")
			Response("not_found", StatusNotFound)
			Response(StatusOK)
		})
	})

	// Event - List (Read multiple)
	Method("ListByPath", func() {
		Description("List Events for a path")
		Payload(func() {
			Attribute("path", String, "Path", func() {
				Example("/croydon/purley/59b287534defd02dace3671b")
			})
			Required("path")
		})
		Result(CollectionOf(EventMedia))
		HTTP(func() {
			GET("/list/{*path}")
			Response("not_found", StatusNotFound)
			Response(StatusOK)
		})
	})

	// Event - Update
	Method("update", func() {
		Description("Update an event")
		Payload(func() {
			Attribute("path", String, "Path", func() {
				Example("/croydon/purley/59b287534defd02dace3671b")
			})
			Attribute("event", EventPayload)
			Required("path")
		})
		HTTP(func() {
			PUT("/update/{*path}")
			Response(StatusOK)
			Response("bad_req", StatusBadRequest)
		})
	})

	// Event - Deactivate (logically delete)
	Method("deactivate", func() {
		Description("Delete an Event")
		Payload(func() {
			Attribute("path", String, "Path", func() {
				Example("/croydon/purley/59b287534defd02dace3671b")
			})
			Required("path")
		})
		HTTP(func() {
			PUT("/deactivate/{*path}")
			Response(StatusOK)
			Response("not_found", StatusNotFound)
		})
	})

	// Event - Delete
	Method("delete", func() {
		Description("Delete an Event")
		Payload(func() {
			Attribute("path", String, "Path", func() {
				Example("/croydon/purley/59b287534defd02dace3671b")
			})
			Required("path")
		})
		HTTP(func() {
			DELETE("/delete/{*path}")
			Response(StatusOK)
			Response("not_found", StatusNotFound)
		})
	})
})

// Organism - an individual or organisation
var _ = Service("organisms", func() {
	HTTP(func() {
		Path("/organisms/")
	})
	Error("bad_req")
	Error("not_found")
	Error("un_auth")
	Error("internal_error", func() {
                Fault()
        })

	// Organism - Post (Create)
	Method("post", func() {
		Description("Add an organism")
		Payload(func() {
			Attribute("path", String, "the path of the org")
			Attribute("organisms", ArrayOf(OrganismPayload))
			Required("path")
		})
		HTTP(func() {
			POST("/{*path}")
			Response(StatusOK)
			Response("bad_req", StatusBadRequest)
		})
	})

	// Organism - Show (Read)
	Method("show", func() {
		Description("Show an organism in full")
		Payload(func() {
			Attribute("path", String, "the path of the org")
			Required("path")
		})
		Result(OrganismMedia)
		HTTP(func() {
			GET("/path/{*path}")
			Response(StatusOK)
			Response("not_found", StatusNotFound)
		})
	})

	// Organism - Update
	Method("update", func() {
		Description("Update an organism")
		Payload(func() {
			Attribute("path", String, "the path of the org")
			Attribute("organism", OrganismPayload)
			Required("path")
		})
		HTTP(func() {
			PUT("/update/{*path}")
			Response(StatusOK)
			Response("bad_req", StatusBadRequest)
		})
	})

	// Organism - Delete
	Method("delete", func() {
		Description("Delete an Organism")
		Payload(func() {
			Attribute("path", String, "the path of the org")
			Required("path")
		})
		HTTP(func() {
			DELETE("/delete/{*path}")
			Response(StatusOK)
			Response("not_found", StatusNotFound)
		})
	})

	// Organism - Deactivate (logically delete)
	Method("deactivate", func() {
		Description("Delete an Organism")
		Payload(func() {
			Attribute("path", String, "Path", func() {
				Example("/croydon/charity")
			})
			Required("path")
		})
		HTTP(func() {
			PUT("/deactivate/{*path}")
			Response(StatusOK)
			Response("not_found", StatusNotFound)
		})
	})

	// Org - List (Read multiple)
	Method("ListByPath", func() {
		Description("List Orgs for a Path")
		Payload(func() {
			Attribute("path", String, "the path of the Org")
			Required("path")
		})
		Result(CollectionOf(OrganismMedia))
		HTTP(func() {
			GET("/list/{*path}")
			Response("not_found", StatusNotFound)
			Response(StatusOK)
		})
	})

	// Org - List By Reference (Read multiple)
	Method("ListByReference", func() {
		Description("List Orgs for a Ref")
		Payload(func() {
			Attribute("path", String, "the path of the Org")
			Required("path")
		})
		Result(CollectionOf(OrganismMedia))
		HTTP(func() {
			GET("/refs/{*path}")
			Response(StatusOK)
			Response("not_found", StatusNotFound)
		})
	})
})
