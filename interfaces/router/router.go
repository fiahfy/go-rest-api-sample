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

func (r *router) Handle(method string, pattern string, handler Handler) {
	reg := regexp.MustCompile(pattern)
	route := route{method, reg, handler}

	r.routes = append(r.routes, route)
}

func (r *router) HandleDefault(handler Handler) {
	r.defaultRoute = handler
}

func (r *router) Get(pattern string, handler Handler) {
	r.Handle(http.MethodGet, pattern, handler)
}

func (r *router) Post(pattern string, handler Handler) {
	r.Handle(http.MethodPost, pattern, handler)
}

func (r *router) Put(pattern string, handler Handler) {
	r.Handle(http.MethodPut, pattern, handler)
}

func (r *router) Delete(pattern string, handler Handler) {
	r.Handle(http.MethodDelete, pattern, handler)
}

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := &Context{w: w, r: req}

	for _, rt := range r.routes {
		if matches := rt.pattern.FindStringSubmatch(c.r.URL.Path); len(matches) > 0 && rt.method == c.r.Method {
			if len(matches) > 1 {
				c.Params = matches[1:]
			}

			rt.handler(c)
			return
		}
	}

	r.defaultRoute(c)
}
