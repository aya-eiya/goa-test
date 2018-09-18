package design

import (
	. "github.com/goadesign/goa/design" // Use . imports to enable the DSL
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("goa test", func() { // API defines the microservice endpoint and
	Title("goa test")       // other global properties. There should be one
	Description("goa test") // and exactly one API definition appearing in
	Contact(func() {
		Name("goa test")
		Email("hidetoshi.ayabe@speee.jp")
		URL("http://goa.design")
	})
	Scheme("http") // the design.
	Host("localhost:8080")
})
