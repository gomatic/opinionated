package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("opinionated", func() {
	Title("An opinionated application starter")
	Description("A basic starter service")
	Scheme("http")
	Host("localhost:3080")
})

var _ = Resource("user", func() {
	BasePath("/u")
	DefaultMedia(UserJSON)

	Action("login", func() {
		Description("Login")
		Routing(GET("/:username"))
		Params(func() {
			Param("username", String, "Username")
		})
		Response(OK)
		Response(NotFound)
	})
})

var UserJSON = MediaType("application/json", func() {
	Description("Credentials")
	Attributes(func() {
		Attribute("id", String, "Username")
		Attribute("token", String, "Token")
		Required("id", "token")
	})
	View("default", func() {
		Attribute("id")
	})
})
