package main

import (
	"fmt"
	"github.com/weirubo/blog-go/global"
	"github.com/weirubo/blog-go/internal/model"
	"github.com/weirubo/blog-go/internal/routers"
	"net/http"
	"time"
)

func init() {
	err := initDBEngine()
	if err != nil {
		fmt.Println("initDBEngine error:", err)
	}
}
func main() {
	router := routers.NewRouter()
	// 自定义 http.Server
	server := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    time.Second * 10,
		WriteTimeout:   time.Second * 10,
		MaxHeaderBytes: 1 << 20,
	}
	_ = server.ListenAndServe()
}

// 初始化数据库
func initDBEngine() error {
	var err error
	// 全局变量赋值
	global.DBEngine, err = model.NewDBEngine()
	if err != nil {
		fmt.Println("initDBEngine err:", err)
		return err
	}
	return nil
}
