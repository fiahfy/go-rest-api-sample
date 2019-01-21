package memory

import (
	"sort"

	"github.com/fiahfy/go-todo-rest-api/domain/model"
	"github.com/fiahfy/go-todo-rest-api/domain/repository"
)

type TodoRepository struct {
	todos map[int]*model.Todo
}

func NewTodoRepository() repository.TodoRepository {
	todos := make(map[int]*model.Todo)
	todos[1] = &model.Todo{ID: 1}
	todos[2] = &model.Todo{ID: 2}
	todos[3] = &model.Todo{ID: 3}

	return &TodoRepository{todos}
}

func (r *TodoRepository) Find(id int) (*model.Todo, error) {
	todo, _ := r.todos[id]
	return todo, nil
}

func (r *TodoRepository) FindAll() ([]*model.Todo, error) {
	todos := []*model.Todo{}
	for _, v := range r.todos {
		todos = append(todos, v)
	}
	sort.Slice(todos, func(i, j int) bool { return todos[i].ID < todos[j].ID })
	return todos, nil
}
