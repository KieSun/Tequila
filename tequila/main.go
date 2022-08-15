package tequila

import (
	"fmt"
	"log"
	"net/http"
)

const Any = "Any"

type routeGroup struct {
	name             string
	handleFuncMap    map[string]map[string]http.HandlerFunc
	handleMethodFunc map[string][]string
}

func (r *routeGroup) handle(name string, method string, handlerFunc http.HandlerFunc) {
	_, ok := r.handleFuncMap[name]
	if !ok {
		r.handleFuncMap[name] = make(map[string]http.HandlerFunc)
	}
	_, ok = r.handleFuncMap[name][method]
	if ok {
		panic("重复路由")
	}
	r.handleFuncMap[name][method] = handlerFunc
	r.handleMethodFunc[method] = append(r.handleMethodFunc[method], name)
}

func (r *routeGroup) Any(name string, handlerFunc http.HandlerFunc) {
	r.handle(name, Any, handlerFunc)
}

func (r *routeGroup) Get(name string, handlerFunc http.HandlerFunc) {
	r.handle(name, http.MethodGet, handlerFunc)
}

func (r *routeGroup) Post(name string, handlerFunc http.HandlerFunc) {
	r.handle(name, http.MethodPost, handlerFunc)
}

type router struct {
	routeGroups []*routeGroup
}

func (r *router) Group(name string) *routeGroup {
	group := &routeGroup{
		name:             name,
		handleFuncMap:    make(map[string]map[string]http.HandlerFunc),
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
		for name, funcMap := range group.handleFuncMap {
			url := "/" + group.name + name
			if r.RequestURI == url {
				fn, ok := funcMap[Any]
				if ok {
					fn(w, r)
					return
				}
				fn, ok = funcMap[method]
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
