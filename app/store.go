package app

import (
	"sort"
)

type Todo struct {
	ID    int
	Title string
	Done  bool
}

var todos map[int]Todo

func InitStore() {
	todos := make(map[int]Todo)
	todos[1] = Todo{ID: 1}
	todos[2] = Todo{ID: 2}
	todos[3] = Todo{ID: 3}
}

func FindAll() []Todo {
	s := []Todo{}
	for _, v := range todos {
		s = append(s, v)
	}
	sort.Slice(s, func(i, j int) bool { return s[i].ID < s[j].ID })
	return s
}

func Find(id int) Todo {
	todo, _ := todos[id]
	return todo
}
