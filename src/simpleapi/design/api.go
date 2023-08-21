package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("simpleapi", func() {
	Title("Simple Api service")
	Description("Service for testing api gateways")
	Server("simpleapi", func() {
		Host("", func() {
			URI("http://")
		})
	})
})

var StoredArticle = ResultType("application/vnd.article", func() {
	Description("A StoredArticle describes article")
	TypeName("StoredArticle")

	Attributes(func() {
		Attribute("id", Int, "ID is the unique id of the article")
		Attribute("title", String)
		Attribute("author", String)
		Attribute("desc", String)
		Attribute("content", String)
		Attribute("admin", String)
	})

	View("default", func() {
		Attribute("id")
		Attribute("title")
		Attribute("author")
		Attribute("desc")
		Attribute("content")
		Attribute("admin")
	})
})

var TimeOutRate = ResultType("application/vnd.timeoutrate", func() {
	Description("Rate of the timeouts response of the endpoint")
	TypeName("TimeOutRate")

	Attributes(func() {
		Attribute("timeoutrate", Int, "rate ot the timeouts in percent")
	})

	View("default", func() {
		Attribute("timeoutrate")
	})
})
