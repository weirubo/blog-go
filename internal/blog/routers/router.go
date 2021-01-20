package routers

import (
	"github.com/gin-gonic/gin"
	apiV1 "github.com/weirubo/blog-go/internal/blog/routers/api/v1"
)

func NewRouter() *gin.Engine {
	// 创建路由
	// 返回一个不包含任何中间件的 Engine 实例
	r := gin.New()
	user := apiV1.User{}
	// 路由组
	v1 := r.Group("v1")
	{
		v1.GET("/user/uid/:uid", user.User)
		v1.GET("/user/list", user.List)
		v1.PUT("/user/modify", user.Modify)
		v1.DELETE("/user/delete/:uid", user.Delete)
		v1.POST("/user/register", user.Register)
		v1.POST("/user/login", user.Login)
		v1.POST("/user/logout/:uid", user.Logout)
	}
	return r
}
