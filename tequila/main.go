package tequila

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

const Any = "Any"

type HandlerFunc func(ctx *Context)
type MiddlewareFunc func(handlerFunc HandlerFunc) HandlerFunc

type routeGroup struct {
	name             string
	handleFuncMap    map[string]map[string]HandlerFunc
	handleMethodFunc map[string][]string
	treeNode         *treeNode
	middlewares      []MiddlewareFunc
}

func (r *routeGroup) Use(middlewares ...MiddlewareFunc) {
	r.middlewares = append(r.middlewares, middlewares...)
}

func (r *routeGroup) HandleMiddleware(handlerFunc HandlerFunc, ctx *Context) {
	for _, middleware := range r.middlewares {
		handlerFunc = middleware(handlerFunc)
	}
	handlerFunc(ctx)
}

func (r *routeGroup) handle(name string, method string, handlerFunc HandlerFunc) {
	_, ok := r.handleFuncMap[name]
	if !ok {
		r.handleFuncMap[name] = make(map[string]HandlerFunc)
	}
	_, ok = r.handleFuncMap[name][method]
	if ok {
		panic("重复路由")
	}
	r.handleFuncMap[name][method] = handlerFunc
	r.handleMethodFunc[method] = append(r.handleMethodFunc[method], name)
	r.treeNode.Put(name)
}

func (r *routeGroup) Any(name string, handlerFunc HandlerFunc) {
	r.handle(name, Any, handlerFunc)
}

func (r *routeGroup) Get(name string, handlerFunc HandlerFunc) {
	r.handle(name, http.MethodGet, handlerFunc)
}

func (r *routeGroup) Post(name string, handlerFunc HandlerFunc) {
	r.handle(name, http.MethodPost, handlerFunc)
}

type router struct {
	routeGroups []*routeGroup
	engine      *Engine
}

func (r *router) Group(name string) *routeGroup {
	group := &routeGroup{
		name:             name,
		handleFuncMap:    make(map[string]map[string]HandlerFunc),
		handleMethodFunc: make(map[string][]string),
		treeNode:         &treeNode{name: "/", children: make([]*treeNode, 0)},
	}
	r.routeGroups = append(r.routeGroups, group)
	return group
}

type Engine struct {
	router
	pool sync.Pool
}

func New() *Engine {
	engine := &Engine{
		router: router{},
	}
	engine.pool.New = func() any {
		return &Context{engine: engine}
	}
	return engine
}

func Default() *Engine {
	engine := New()
	return engine
}

func (e *Engine) allocateContext() any {
	return &Context{engine: e}
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := e.pool.Get().(*Context)
	ctx.W = w
	ctx.R = r
	e.httpRequestHandle(ctx, w, r)
	e.pool.Put(ctx)
}

func (e *Engine) httpRequestHandle(ctx *Context, w http.ResponseWriter, r *http.Request) {
	method := r.Method
	for _, group := range e.router.routeGroups {
		routerName := SubString(r.URL.Path, group.name)
		node := group.treeNode.Get(routerName)
		if node != nil {
			fn, ok := group.handleFuncMap[node.routerName][Any]
			if ok {
				group.HandleMiddleware(fn, ctx)
				return
			}
			fn, ok = group.handleFuncMap[node.routerName][method]
			if ok {
				group.HandleMiddleware(fn, ctx)
				return
			}
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Print("not allowed")
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
