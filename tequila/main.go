package tequila

import (
	"log"
	"net/http"
)

type route struct {
	handleFunc map[string]http.HandlerFunc
}

func (r *route) Add(name string, handlerFunc http.HandlerFunc) {
	r.handleFunc[name] = handlerFunc
}

type Engine struct {
	route
}

func New() *Engine {
	return &Engine{
		route{handleFunc: make(map[string]http.HandlerFunc)},
	}
}

func (e *Engine) Run() {
	for key, value := range e.handleFunc {
		http.HandleFunc(key, value)
	}
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal(err)
	}
}
