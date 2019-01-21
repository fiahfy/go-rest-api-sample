package registry

import (
	"github.com/fiahfy/go-todo-rest-api/application/usecase"
	"github.com/fiahfy/go-todo-rest-api/infrastructure/memory"
	"github.com/fiahfy/go-todo-rest-api/interfaces/handler"
)

type Registry interface {
	NewAppHandler() handler.AppHandler
}

type registry struct{}

func NewRegistry() Registry {
	return &registry{}
}

func newHelloHandler() handler.HelloHandler {
	return handler.NewHelloHandler()
}

func newTodoHandler() handler.TodoHandler {
	r := memory.NewTodoRepository()
	u := usecase.NewTodoUseCase(r)
	return handler.NewTodoHandler(u)
}
func (r *registry) NewAppHandler() handler.AppHandler {
	return handler.NewAppHandler(
		newHelloHandler(),
		newTodoHandler(),
	)
}
