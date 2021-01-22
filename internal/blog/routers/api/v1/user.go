package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/weirubo/blog-go/internal/blog/service"
	"net/http"
	"strconv"
)

type User struct {
}

func (u User) Add(c *gin.Context) {
	param := service.CreateUserRequest{}
	err := c.ShouldBind(&param)
	if err != nil {
		fmt.Println("add error:", err)
		return
	}
	s := service.New()
	err = s.AddUser(&param)
	if err != nil {
		fmt.Println("add err:", err)
		return
	}
	data := gin.H{}
	c.JSON(http.StatusOK, data)
}

func (u User) Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println("get error:", err)
		return
	}
	s := service.New()
	user, err := s.GetUser(id)
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
	param := service.UserListByPageRequest{}
	err := c.ShouldBind(&param)
	if err != nil {
		fmt.Println("list error:", err)
		return
	}
	s := service.New()
	users, err := s.UserList(&param)
	if err != nil {
		fmt.Println("list error:", err)
		return
	}
	data := gin.H{
		"list": users,
	}
	c.JSON(http.StatusOK, data)
}

func (u User) Edit(c *gin.Context) {
	param := service.EditUserRequest{}
	err := c.ShouldBind(&param)
	if err != nil {
		fmt.Println("edit error:", err)
		return
	}
	s := service.New()
	rows, err := s.EditUser(&param)
	if err != nil {
		fmt.Println("edit error:", err)
		return
	}
	data := gin.H{
		"rows": rows,
	}
	c.JSON(http.StatusOK, data)
}

func (u User) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println("delete error:", err)
		return
	}
	s := service.New()
	rows, err := s.DeleteUser(uint(id))
	if err != nil {
		fmt.Println("delete error:", err)
		return
	}
	data := gin.H{
		"rows": rows,
	}
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
		fmt.Println("login error:", err)
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
