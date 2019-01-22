package repository

import (
	"github.com/fiahfy/go-todo-rest-api/domain/model"
)

type TodoRepository interface {
	Find(int) (*model.Todo, error)
	FindAll() ([]*model.Todo, error)
	Create(*model.Todo) (*model.Todo, error)
	Update(*model.Todo) error
	Delete(int) error
}
