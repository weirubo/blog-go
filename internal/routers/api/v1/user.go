package v1

import (
	"github.com/gin-gonic/gin"
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
