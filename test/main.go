package main

import (
	"fmt"
	"github/KieSun/Tequila/tequila"
	"net/http"
)

func main() {
	engine := tequila.New()
	user := engine.Group("user")
	user.Add("/1", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Printf("123123")
	})
	engine.Run()
}
