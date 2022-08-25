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
	user.Get("/id", func(c *tequila.Context) {
		err := c.Json(http.StatusOK, &User{Name: "yck"})
		value := c.GetQuery("id")
		fmt.Print(value)
		if err != nil {
			fmt.Print(err)
		}
		fmt.Printf("get")
	})

	engine.Run()
}
