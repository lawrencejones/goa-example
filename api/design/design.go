package design

import (
	. "goa.design/goa/v3/dsl"
)

var Incident = Type("Incident", func() {
	Attribute("id", String, "Unique identifier for the incident", func() {
		Example("01FDAG4SAP5TYPT98WGR2N7W91")
	})
	Attribute("name", String, "Name of the incident", func() {
		Example("Full service outage")
	})
	Required(
		"id",
		"name",
	)
})

var _ = Service("Incidents", func() {
	Description("Manage incidents")

	HTTP(func() {
		Path("/api/incidents")
	})

	Method("Create", func() {
		Description("Create a new incident")

		Payload(func() {
			Reference(Incident)
			Attribute("name")
		})

		Result(Incident)

		HTTP(func() {
			POST("/")
		})
	})
})
