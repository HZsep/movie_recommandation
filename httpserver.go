                                                                                 package main

import (
 "fmt"
 "net/http"
)

// HandlerFunc defines the request handler used by the web framework
type HandlerFunc func(http.ResponseWriter, *http.Request)

// Engine is the core of the web framework
type Engine struct {
 router map[string]HandlerFunc
}

// NewEngine creates an instance of the web framework
func NewEngine() *Engine {
 return &Engine{router: make(map[string]HandlerFunc)}
}

// AddRoute adds a new route to the web framework
func (engine *Engine) AddRoute(method string, pattern string, handler HandlerFunc) {
 key := method + "-" + pattern
 engine.router[key] = handler
}

// ServeHTTP implements the http.Handler interface.
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
 key := req.Method + "-" + req.URL.Path
 if handler, ok := engine.router[key]; ok {
  handler(w, req)
 } else {
  w.WriteHeader(http.StatusNotFound)
  fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL.Path)
 }
}

func main() {
 engine := NewEngine()

 engine.AddRoute("GET", "/", func(w http.ResponseWriter, req *http.Request) {
  fmt.Fprintf(w, "Welcome to the home page!")
 })

 engine.AddRoute("GET", "/hello", func(w http.ResponseWriter, req *http.Request) {
  fmt.Fprintf(w, "Hello, world!")
 })

 http.ListenAndServe(":8080", engine)
}                                                                         
