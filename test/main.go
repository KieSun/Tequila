package main

import (
	"fmt"
	"github/KieSun/Tequila/tequila"
	"net/http"
)

type User struct {
	Name string
}

func main() {
	engine := tequila.New()
	user := engine.Group("user")
	//user.Post("/1", func(writer http.ResponseWriter, request *http.Request) {
	//	fmt.Printf("post")
	//})
	user.Use(func(handlerFunc tequila.HandlerFunc) tequila.HandlerFunc {
		return func(ctx *tequila.Context) {
			handlerFunc(ctx)
		}
	})
	user.Use(func(handlerFunc tequila.HandlerFunc) tequila.HandlerFunc {
		return func(ctx *tequila.Context) {
			handlerFunc(ctx)
		}
	})
	user.Get("/:id", func(c *tequila.Context) {
		err := c.Json(http.StatusOK, &User{Name: "yck"})
		if err != nil {
			fmt.Print(err)
		}
		fmt.Printf("get")
	})

	engine.Run()
}
