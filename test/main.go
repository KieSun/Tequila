package main

import (
	"fmt"
	"github/KieSun/Tequila/tequila"
	"net/http"
)

func main() {
	engine := tequila.New()
	user := engine.Group("user")
	//user.Post("/1", func(writer http.ResponseWriter, request *http.Request) {
	//	fmt.Printf("post")
	//})
	user.Get("/:id", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Printf("get")
	})
	engine.Run()
}
