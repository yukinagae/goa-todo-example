package todoapi

import (
	"context"
	"fmt"
	"log"
	todo "todo/gen/todo"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// todo service example implementation.
// The example methods log the requests and return zero values.
type todosrvc struct {
	db     *gorm.DB
	logger *log.Logger
}

// NewTodo returns the todo service implementation.
func NewTodo(db *gorm.DB, logger *log.Logger) todo.Service {
	return &todosrvc{db, logger}
}

// Get implements get.
func (s *todosrvc) Get(ctx context.Context, p *todo.GetPayload) (res *todo.Todo, err error) {
	res = &todo.Todo{}
	s.logger.Print("todo.get")

	for _, todo := range Todos {
		if todo.ID == p.ID {
			return todo, nil
		}
	}
	return nil, fmt.Errorf("Todo not found: %s", p.ID)
}

// List implements list.
func (s *todosrvc) List(ctx context.Context) (res todo.TodoCollection, err error) {
	s.logger.Print("todo.list")
	s.logger.Print(s.db)
	s.db.Find(&res)
	return res, nil
}

// Add implements add.
func (s *todosrvc) Add(ctx context.Context, p *todo.AddPayload) (res string, err error) {
	s.logger.Print("todo.add")
	Todos = append(Todos, &todo.Todo{ID: p.ID, Task: p.Task})
	return p.ID, nil
}

// Remove implements remove.
func (s *todosrvc) Remove(ctx context.Context, p *todo.RemovePayload) (err error) {
	s.logger.Print("todo.remove")
	for i, todo := range Todos {
		if todo.ID == p.ID {
			Todos[i] = Todos[len(Todos)-1] // Copy last element to index i.
			Todos[len(Todos)-1] = nil      // Erase last element (write zero value).
			Todos = Todos[:len(Todos)-1]   // Truncate slice.
			return nil
		}
	}
	return fmt.Errorf("Todo not found: %s", p.ID)
}

// Todos is dummy data
var Todos = []*todo.Todo{
	{ID: "todo1", Task: "build an API"},
	{ID: "todo2", Task: "?????"},
	{ID: "todo3", Task: "profit!"},
}
