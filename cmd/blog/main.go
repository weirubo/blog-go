package main

import (
	"github.com/weirubo/blog-go/internal/routers"
	"net/http"
	"time"
)

func main () {
	router := routers.NewRouter()
	// 自定义 http.Server
	server := new(http.Server)
	server.Addr = ":8080"
	server.Handler = router
	server.ReadTimeout = time.Second * 10
	server.WriteTimeout = time.Second * 10
	server.MaxHeaderBytes = 1 << 20
	server.ListenAndServe()
}
