package routers

import (
	"github.com/gin-gonic/gin"
	apiV1 "github.com/weirubo/blog-go/internal/blog/routers/api/v1"
)

func NewRouter() *gin.Engine {
	// 创建路由
	// 返回一个不包含任何中间件的 Engine 实例
	router := gin.New()
	user := apiV1.User{}
	article := apiV1.Article{}
	category := apiV1.Category{}
	tag := apiV1.Tag{}
	// 路由组
	v1 := router.Group("v1")
	{
		v1.GET("/user/get/:uid", user.Get)
		v1.GET("/user/list", user.List)
		v1.POST("/user/add", user.Add)
		v1.PUT("/user/edit", user.Edit)
		v1.DELETE("/user/delete/:uid", user.Delete)
		v1.POST("/user/login", user.Login)
		v1.POST("/user/logout/:uid", user.Logout)

		v1.GET("/category/get/:id", category.Get)
		v1.GET("/category/list", category.List)
		v1.POST("/category/add", category.Add)
		v1.PUT("/category/edit", category.Edit)
		v1.POST("/category/delete/:id", category.Delete)

		v1.GET("/tag/get/:id", tag.Get)
		v1.GET("/tag/list", tag.List)
		v1.POST("/tag/add", tag.Add)
		v1.PUT("/tag/edit", tag.Edit)
		v1.DELETE("/tag/delete/:id", tag.Delete)

		v1.GET("/article/get/:id", article.Get)
		v1.GET("/article/list", article.List)
		v1.POST("/article/add", article.Add)
		v1.PUT("/article/edit", article.Edit)
		v1.POST("/article/delete/:id", article.Delete)
	}
	return router
}
