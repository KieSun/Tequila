package tequila

import (
	"log"
	"net/http"
)

type routeGroup struct {
	name       string
	handleFunc map[string]http.HandlerFunc
}

func (r *routeGroup) Add(name string, handlerFunc http.HandlerFunc) {
	r.handleFunc[name] = handlerFunc
}

type router struct {
	routeGroups []*routeGroup
}

func (r *router) Group(name string) *routeGroup {
	group := &routeGroup{name: name, handleFunc: make(map[string]http.HandlerFunc)}
	r.routeGroups = append(r.routeGroups, group)
	return group
}

type Engine struct {
	router
}

func New() *Engine {
	return &Engine{
		router{},
	}
}

func (e *Engine) Run() {
	for _, group := range e.routeGroups {
		for key, value := range group.handleFunc {
			http.HandleFunc("/"+group.name+key, value)
		}
	}

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal(err)
	}
}
