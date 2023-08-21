package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("stuff", func() {
	Description("Service provides management od articles stuff")

	Error("NotFound", func() {
		Description("Notfound is the error returned by the service methods when the id of the stuff is not found.")
	})

	Error("Timeout", func() {
		Description("Operation timed out")
	})

	Error("InternalError", func() {
		Description("Internal Server Error")
	})

	Method("list", func() {

		Description("List all stored articles")
		Result(CollectionOf(StoredArticle), func() {
			View("default")
		})

		HTTP(func() {
			GET("/articles")
			Response(StatusOK)
		})
	})

	Method("Add", func() {

		Payload(func() {
			Field(1, "Title", String, "Title")
			Field(2, "Author", String, "Author name")
			Field(3, "Desc", String, "Description")
			Field(4, "Content", String, "Content description")
			Attribute("Admin", String, "Who stored the column")
			Required("Author", "Title")
		})
		Result(Empty)
		Description("Add article to the articles")
		HTTP(func() {
			POST("/article")
			Header("Admin:X-Simple-Admin")
			Response(StatusCreated)
		})
	})

	Method("show", func() {

		Description("Show article by id")
		Payload(func() {
			Field(1, "id", Int, "Id of article to show")
			Required("id")
		})
		Result(StoredArticle)

		HTTP(func() {
			GET("/article/{id}")
			Response("NotFound", StatusNotFound)
			Response("Timeout", StatusInternalServerError)
			Response("InternalError", StatusInternalServerError)
			Response(StatusOK)
		})
	})

})
