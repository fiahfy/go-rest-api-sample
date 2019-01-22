package router

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

type Context struct {
	w      http.ResponseWriter
	r      *http.Request
	Params []string
}

func (c *Context) Text(code int, body string) {
	c.w.Header().Set("Content-Type", "text/plain")
	c.w.WriteHeader(code)
	fmt.Fprintf(c.w, fmt.Sprintf("%s\n", body))
}

func (c *Context) Json(code int, data interface{}) {
	o := &response{code, data}
	res, err := json.Marshal(o)
	if err != nil {
		http.Error(c.w, err.Error(), http.StatusInternalServerError)
		return
	}
	c.w.Header().Set("Content-Type", "application/json")
	c.w.WriteHeader(code)
	c.w.Write(res)
}
