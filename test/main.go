package main

import (
	"fmt"
	"github/KieSun/Tequila/tequila"
)

func main() {
	engine := tequila.New()
	user := engine.Group("user")
	//user.Post("/1", func(writer http.ResponseWriter, request *http.Request) {
	//	fmt.Printf("post")
	//})
	user.Use(func(handlerFunc tequila.HandlerFunc) tequila.HandlerFunc {
		return func(ctx *tequila.Context) {
			fmt.Printf("middleware")
			handlerFunc(ctx)
		}
	})
	user.Get("/:id", func(c *tequila.Context) {
		fmt.Printf("get")
	})
	engine.Run()
}
