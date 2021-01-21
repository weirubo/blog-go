package main

import (
	"fmt"
	"github.com/weirubo/blog-go/global"
	"github.com/weirubo/blog-go/internal/blog/model"
	"github.com/weirubo/blog-go/internal/blog/routers"
	"github.com/weirubo/blog-go/pkg/config"
	"net/http"
	"time"
)

func init() {
	err := initConfig()
	if err != nil {
		fmt.Println("initConfig error:", err)
	}
	err = initDBEngine()
	if err != nil {
		fmt.Println("initDBEngine error:", err)
	}
}
func main() {
	router := routers.NewRouter()
	// 自定义 http.Server
	server := &http.Server{
		Addr:           ":" + global.ServerConfig.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerConfig.ReadTimeout,
		WriteTimeout:   global.ServerConfig.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	_ = server.ListenAndServe()
}

// 初始化配置
func initConfig() error {
	config, err := config.NewViper()
	if err != nil {
		return err
	}
	err = config.ReadConfig("Server", &global.ServerConfig)
	if err != nil {
		return err
	}
	err = config.ReadConfig("Database", &global.DatabaseConfig)
	if err != nil {
		return err
	}
	global.ServerConfig.ReadTimeout *= time.Second
	global.ServerConfig.WriteTimeout *= time.Second
	return nil
}

// 初始化数据库
func initDBEngine() error {
	var err error
	// 全局变量赋值
	global.DBEngine, err = model.NewDBEngine(global.DatabaseConfig)
	if err != nil {
		fmt.Println("initDBEngine err:", err)
		return err
	}
	return nil
}
