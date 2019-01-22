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
	todos[1] = &model.Todo{ID: 1, Title: "Task 1"}
	todos[2] = &model.Todo{ID: 2, Title: "Task 2"}
	todos[3] = &model.Todo{ID: 3, Title: "Task 3"}

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

func (r *TodoRepository) Create(m *model.Todo) (*model.Todo, error) {
	id := 0
	for k := range r.todos {
		if k > id {
			id = k
		}
	}
	id++
	m.ID = id
	r.todos[id] = m
	return m, nil
}

func (r *TodoRepository) Update(m *model.Todo) error {
	_, ok := r.todos[m.ID]
	if !ok {
		return nil
	}
	r.todos[m.ID] = m
	return nil
}

func (r *TodoRepository) Delete(id int) error {
	delete(r.todos, id)
	return nil
}
