package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("todo", func() {
	Title("Todo Service")
	Description("Service for managing tasks")
	Server("todo", func() {
		Host("localhost", func() {
			URI("http://localhost:8080")
			URI("grpc://localhost:8888")
		})
	})
})

var _ = Service("todo", func() {
	Method("get", func() {
		Payload(func() {
			Field(1, "id", String, "Todo ID")
			Required("id")
		})

		Result(Todo)

		HTTP(func() {
			GET("/{id}")
		})

		GRPC(func() {
		})
	})

	Files("/openapi.json", "./gen/http/openapi.json")
})

// Todo is ResultType
var Todo = ResultType("Todo", func() {
	Attribute("id", String, "Todo ID", func() {
		Example("todo1")
		Meta("rpc:tag", "1")
	})
	Attribute("task", String, "", func() {
		Example("build an API")
		Meta("rpc:tag", "2")
	})
	Required("id", "task")
})
