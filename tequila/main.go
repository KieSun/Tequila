package tequila

import (
	"fmt"
	"log"
	"net/http"
)

type routeGroup struct {
	name             string
	handleFunc       map[string]http.HandlerFunc
	handleMethodFunc map[string][]string
}

func (r *routeGroup) Any(name string, handlerFunc http.HandlerFunc) {
	r.handleFunc[name] = handlerFunc
	r.handleMethodFunc["Any"] = append(r.handleMethodFunc["Any"], name)
}

func (r *routeGroup) Get(name string, handlerFunc http.HandlerFunc) {
	r.handleFunc[name] = handlerFunc
	r.handleMethodFunc[http.MethodGet] = append(r.handleMethodFunc[http.MethodGet], name)
}

func (r *routeGroup) Post(name string, handlerFunc http.HandlerFunc) {
	r.handleFunc[name] = handlerFunc
	r.handleMethodFunc[http.MethodPost] = append(r.handleMethodFunc[http.MethodPost], name)
}

type router struct {
	routeGroups []*routeGroup
}

func (r *router) Group(name string) *routeGroup {
	group := &routeGroup{
		name:             name,
		handleFunc:       make(map[string]http.HandlerFunc),
		handleMethodFunc: make(map[string][]string),
	}
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

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	for _, group := range e.router.routeGroups {
		for route, fn := range group.handleFunc {
			url := "/" + group.name + route
			if r.RequestURI == url {
				_, ok := group.handleMethodFunc["Any"]
				if ok {
					fn(w, r)
					return
				}
				_, ok = group.handleMethodFunc[method]
				if ok {
					fn(w, r)
					return
				}
				w.WriteHeader(http.StatusMethodNotAllowed)
				fmt.Print("not allowed")
				return
			}
		}
	}
	w.WriteHeader(http.StatusNotFound)
	fmt.Print("not found")
}

func (e *Engine) Run() {
	http.Handle("/", e)

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal(err)
	}
}
