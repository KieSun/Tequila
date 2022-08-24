package tequila

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Context struct {
	W http.ResponseWriter
	R *http.Request
}

func (ctx *Context) Json(status int, data any) error {
	jsonData, err := json.Marshal(data)
	fmt.Print(jsonData)
	if err != nil {
		return err
	}
	ctx.W.Header().Set("Content-type", "application/json; charset=utf-8")
	ctx.W.WriteHeader(status)
	_, err = ctx.W.Write(jsonData)
	return err
}
