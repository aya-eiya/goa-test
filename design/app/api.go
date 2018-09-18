package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var MyTypes = Type("MyTypes", func() {
	Attribute("arrs", ArrayOf(MyType), "MyType Array")
	Attribute("intVal", Integer, "non-required Integer")
})

var MyType = Type("MyType", func() {
	Attribute("arrs", ArrayOf(String), "String Array")
	Attribute("intVal", Integer, "non-required Integer")
})

var _ = Resource("my_resorce", func() {

	BasePath("/actions")
	Action("my_action", func() {
		Description("my action")
		Routing(POST("/my_action"))
		Payload(MyType)
		MultipartForm()
		Response(OK, MyTypes)
	})
})
