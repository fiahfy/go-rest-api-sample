package usecase

import (
	"github.com/fiahfy/go-rest-api-sample/domain/model"
	"github.com/fiahfy/go-rest-api-sample/domain/repository"
)

type TodoUseCase interface {
	Find(int) (*model.Todo, error)
	FindAll() ([]*model.Todo, error)
	Create(*model.Todo) (*model.Todo, error)
	Update(*model.Todo) error
	Delete(int) error
}

type todoUseCase struct {
	r repository.TodoRepository
}

func NewTodoUseCase(r repository.TodoRepository) TodoUseCase {
	return &todoUseCase{r}
}

func (u *todoUseCase) Find(id int) (*model.Todo, error) {
	return u.r.Find(id)
}

func (u *todoUseCase) FindAll() ([]*model.Todo, error) {
	return u.r.FindAll()
}

func (u *todoUseCase) Create(m *model.Todo) (*model.Todo, error) {
	return u.r.Create(m)
}

func (u *todoUseCase) Update(m *model.Todo) error {
	return u.r.Update(m)
}

func (u *todoUseCase) Delete(id int) error {
	return u.r.Delete(id)
}
