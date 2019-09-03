package todoapi

import (
	"context"
	"log"
	todo "todo/gen/todo"
)

// todo service example implementation.
// The example methods log the requests and return zero values.
type todosrvc struct {
	logger *log.Logger
}

// NewTodo returns the todo service implementation.
func NewTodo(logger *log.Logger) todo.Service {
	return &todosrvc{logger}
}

// Get implements get.
func (s *todosrvc) Get(ctx context.Context, p *todo.GetPayload) (res *todo.Todo, err error) {
	res = &todo.Todo{}
	s.logger.Print("todo.get")
	return &todo.Todo{ID: "todo1", Task: "build an API"}, nil
}
