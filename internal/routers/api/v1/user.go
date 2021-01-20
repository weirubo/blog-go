package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/weirubo/blog-go/internal/service"
	"net/http"
	"strconv"
)

type User struct {
}

func (u User) User(c *gin.Context) {
	uid, err := strconv.Atoi(c.Param("uid"))
	if err != nil {
		fmt.Println("uid error:", err)
		return
	}
	s := service.New()
	user, err := s.User(uid)
	if err != nil {
		fmt.Println("get error:", err)
		return
	}
	data := gin.H{
		"data": user,
	}
	c.JSON(http.StatusOK, data)
}

func (u User) List(c *gin.Context) {
	param := service.UserListByPage{}
	err := c.ShouldBind(&param)
	if err != nil {
		fmt.Println("list error:", err)
		return
	}
	s := service.New()
	users, err := s.List(&param)
	if err != nil {
		fmt.Println("List error:", err)
		return
	}
	data := gin.H{
		"list": users,
	}
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
	err := c.ShouldBind(&param)
	if err != nil {
		fmt.Println("register error:", err)
		return
	}
	s := service.New()
	err = s.Register(&param)
	if err != nil {
		fmt.Println("Register err:", err)
		return
	}
	data := gin.H{}
	c.JSON(http.StatusOK, data)
}

func (u User) Login(c *gin.Context) {
	param := service.CreateUserRequest{}
	err := c.ShouldBind(&param)
	if err != nil {
		fmt.Println("login error:", err)
		return
	}
	s := service.New()
	user, err := s.Login(&param)
	if err != nil {
		fmt.Println("Login error:", err)
		return
	}
	data := gin.H{
		"uid": user.ID,
	}
	c.JSON(http.StatusOK, data)
}

func (u User) Logout(c *gin.Context) {
	data := gin.H{}
	c.JSON(http.StatusOK, data)
}
