package main

import (
	"github/KieSun/Tequila/tequila"
	"net/http"
)

func main() {
	server := &http.Server{
		// 自定义的请求核心处理函数
		Handler: tequila.NewCore(),
		// 请求监听地址
		Addr: ":8080",
	}
	_ = server.ListenAndServe()
}
