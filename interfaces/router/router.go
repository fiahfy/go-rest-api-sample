package router

import (
	"net/http"
	"regexp"
)

type Handler func(*Context)

type route struct {
	method  string
	pattern *regexp.Regexp
	handler Handler
}

type Router interface {
	http.Handler
	Handle(method string, pattern string, handler Handler)
	HandleDefault(handler Handler)
	Get(pattern string, handler Handler)
	Post(pattern string, handler Handler)
	Put(pattern string, handler Handler)
	Delete(pattern string, handler Handler)
}

type router struct {
	routes       []route
	defaultRoute Handler
}

func New() Router {
	return &router{
		defaultRoute: func(c *Context) {
			c.Text(http.StatusNotFound, "Not found")
		},
	}
}

func (a *router) Handle(method string, pattern string, handler Handler) {
	reg := regexp.MustCompile(pattern)
	route := route{method, reg, handler}

	a.routes = append(a.routes, route)
}

func (a *router) HandleDefault(handler Handler) {
	a.defaultRoute = handler
}

func (a *router) Get(pattern string, handler Handler) {
	a.Handle(http.MethodGet, pattern, handler)
}

func (a *router) Post(pattern string, handler Handler) {
	a.Handle(http.MethodPost, pattern, handler)
}

func (a *router) Put(pattern string, handler Handler) {
	a.Handle(http.MethodPut, pattern, handler)
}

func (a *router) Delete(pattern string, handler Handler) {
	a.Handle(http.MethodDelete, pattern, handler)
}

func (a *router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := &Context{w: w, r: r}

	for _, rt := range a.routes {
		if matches := rt.pattern.FindStringSubmatch(c.r.URL.Path); len(matches) > 0 && rt.method == c.r.Method {
			if len(matches) > 1 {
				c.Params = matches[1:]
			}

			rt.handler(c)
			return
		}
	}

	a.defaultRoute(c)
}
