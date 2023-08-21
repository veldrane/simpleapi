package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("failures", func() {
	Description("Service provides management for generating failures on stuff api")

	Error("BadRequest", func() {
        Description("Error when timeoutrate is out of the scope")
	})

	Method("GetTimeout", func() {

		Description("List all stored articles")
		Result(TimeOutRate)
		HTTP(func() {
			GET("/timeout/article")
			Response(StatusOK)
		})
	})


	Method("SetTimeout", func() {

		Payload(func() {
			Field(1, "TimeOutRate", Int, "Rate of the failures on /article/{id}")
			Required("TimeOutRate")
		})
		Result(Empty)
		Description("Manage timeouts on the /article/{id} endpoint")
		HTTP(func() {
			POST("/timeout/article")
			Response(StatusOK)
			Response(StatusBadRequest)
		})
	})
})