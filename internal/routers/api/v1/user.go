package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/weirubo/blog-go/internal/service"
	"net/http"
)

type User struct {
}

func (u User) Get(c *gin.Context) {
	data := gin.H{}
	c.JSON(http.StatusOK, data)
}

func (u User) List(c *gin.Context) {
	data := gin.H{}
	c.JSON(http.StatusOK, data)
}

func (u User) Modify(c *gin.Context) {
	data := gin.H{}
	c.JSON(http.StatusOK, data)
}

func (u User) Delete(c *gin.Context) {
	data := gin.H{}
	c.JSON(http.StatusOK, data)
}

func (u User) Register(c *gin.Context) {
	param := service.CreateUserRequest{}
	c.ShouldBind(&param)
	s := service.New()
	err := s.CreateUser(&param)
	if err != nil {
		fmt.Println("Register err:", err)
		return
	}
	data := gin.H{}
	c.JSON(http.StatusOK, data)
}

func (u User) Login(c *gin.Context) {
	data := gin.H{}
	c.JSON(http.StatusOK, data)
}

func (u User) Logout(c *gin.Context) {
	data := gin.H{}
	c.JSON(http.StatusOK, data)
}
