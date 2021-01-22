package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/weirubo/blog-go/internal/blog/service"
	"net/http"
	"strconv"
)

type Article struct {
}

func (a Article) Add(c *gin.Context) {
	param := service.CreateArticleRequest{}
	err := c.ShouldBind(&param)
	if err != nil {
		fmt.Println("add error:", err)
		return
	}
	s := service.New()
	err = s.AddArticle(&param)
	if err != nil {
		fmt.Println("add error:", err)
		return
	}
	data := gin.H{}
	c.JSON(http.StatusOK, data)
}

func (a Article) Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println("get error:", err)
		return
	}
	s := service.New()
	article, err := s.GetArticle(id)
	if err != nil {
		fmt.Println("get error:", err)
		return
	}
	data := gin.H{
		"data": article,
	}
	c.JSON(http.StatusOK, data)
}

func (a Article) List(c *gin.Context) {
	s := service.New()
	articles, err := s.ArticlesList()
	if err != nil {
		fmt.Println("list error:", err)
		return
	}
	data := gin.H{
		"list": articles,
	}
	c.JSON(http.StatusOK, data)
}

func (a Article) Edit(c *gin.Context) {
	param := service.EditArticleRequest{}
	err := c.ShouldBind(&param)
	if err != nil {
		fmt.Println("edit error:", err)
		return
	}
	s := service.New()
	rows, err := s.EditArticle(&param)
	if err != nil {
		fmt.Println("edit error:", err)
		return
	}
	data := gin.H{
		"rows": rows,
	}
	c.JSON(http.StatusOK, data)
}

func (a Article) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println("delete error:", err)
		return
	}
	s := service.New()
	rows, err := s.DeleteArticle(uint(id))
	if err != nil {
		fmt.Println("delete error:", err)
		return
	}
	data := gin.H{
		"rows": rows,
	}
	c.JSON(http.StatusOK, data)
}
