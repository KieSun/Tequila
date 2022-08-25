package tequila

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type Context struct {
	W          http.ResponseWriter
	R          *http.Request
	engine     *Engine
	queryCache url.Values
}

func (ctx *Context) Json(status int, data any) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	ctx.W.Header().Set("Content-type", "application/json; charset=utf-8")
	ctx.W.WriteHeader(status)
	_, err = ctx.W.Write(jsonData)
	return err
}

func (ctx *Context) GetQuery(key string) any {
	ctx.initQueryCache()
	return ctx.queryCache.Get(key)
}

func (ctx *Context) initQueryCache() {
	if ctx.queryCache == nil {
		ctx.queryCache = ctx.R.URL.Query()
	}
}
