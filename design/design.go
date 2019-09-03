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

	HTTP(func() {
		Path("/")
	})

	Method("get", func() {
		Payload(func() {
			Field(1, "id", String, "Todo ID")
			Required("id")
		})
		Result(Todo)
		Error("not_found", NotFound, "Bottle not found")
		HTTP(func() {
			GET("/{id}")
			Response(StatusOK)
		})
		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("list", func() {
		Result(CollectionOf(Todo))
		HTTP(func() {
			GET("/")
			Response(StatusOK)
		})
		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("add", func() {
		Payload(func() {
			Field(1, "id", String, "Todo ID")
			Field(2, "task", String, "Todo task")
			Required("id", "task")
		})
		Result(String)
		HTTP(func() {
			POST("/")
			Response(StatusCreated)
		})
		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("remove", func() {
		Payload(func() {
			Field(1, "id", String, "Todo ID")
			Required("id")
		})
		Error("not_found", NotFound, "Bottle not found")
		HTTP(func() {
			DELETE("/{id}")
			Response(StatusNoContent)
		})
		GRPC(func() {
			Response(CodeOK)
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

// NotFound is Type
var NotFound = Type("NotFound", func() {
	Description("NotFound is the type returned when attempting to show or delete a todo that does not exist.")
	Attribute("message", String, "Message of error", func() {
		Meta("struct:error:name")
		Example("todo todo1 not found")
		Meta("rpc:tag", "1")
	})
	Field(2, "id", String, "ID of missing Todo")
	Required("message", "id")
})
