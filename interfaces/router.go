package interfaces

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/fiahfy/go-todo-rest-api/app/usecase"
	"github.com/fiahfy/go-todo-rest-api/infra/inmem"
	"github.com/fiahfy/go-todo-rest-api/interfaces/handler"
	"github.com/fiahfy/go-todo-rest-api/interfaces/httputils"
)

type Handler func(*httputils.Context)

type Route struct {
	Pattern *regexp.Regexp
	Handler Handler
}

type Router struct {
	Routes       []Route
	DefaultRoute Handler
}

func NewRouter() *Router {
	r := &Router{
		DefaultRoute: func(ctx *httputils.Context) {
			ctx.Text(http.StatusNotFound, "Not found")
		},
	}

	r.Handle(`^/hello$`, func(ctx *httputils.Context) {
		ctx.Text(http.StatusOK, "Hello world")
	})

	r.Handle(`^/hello/([\w\._-]+)$`, func(ctx *httputils.Context) {
		ctx.Text(http.StatusOK, fmt.Sprintf("Hello %s", ctx.Params[0]))
	})

	// r.Handle(`^/todos$`, func(ctx *Context) {
	// 	todos := FindAll()
	// 	fmt.Printf("%+v", todos)
	// 	ctx.Text(http.StatusOK, fmt.Sprintf("Todo Index"))
	// })
	// r.Handle(`^/todos/(\d+)$`, func(ctx *Context) {
	// 	id, _ := strconv.Atoi(ctx.Params[0])
	// 	todo := Find(id)
	// 	fmt.Printf("%+v", todo)
	// 	ctx.Text(http.StatusOK, fmt.Sprintf("Todo %s", ctx.Params[0]))
	// })

	repo := inmem.NewTodoRepository()
	u := usecase.NewTodoUseCase(repo)
	h := handler.NewTodoHandler(u)

	r.Handle(`^/todos/(\d+)$`, h.GetTodo)
	r.Handle(`^/todos$`, h.ListTodos)

	return r
}

func (a *Router) Handle(pattern string, handler Handler) {
	re := regexp.MustCompile(pattern)
	route := Route{Pattern: re, Handler: handler}

	a.Routes = append(a.Routes, route)
}

func (a *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := &httputils.Context{Request: r, ResponseWriter: w}

	for _, rt := range a.Routes {
		if matches := rt.Pattern.FindStringSubmatch(ctx.URL.Path); len(matches) > 0 {
			if len(matches) > 1 {
				ctx.Params = matches[1:]
			}

			rt.Handler(ctx)
			return
		}
	}

	a.DefaultRoute(ctx)
}
