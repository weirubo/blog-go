package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/weirubo/blog-go/internal/blog/service"
	"net/http"
	"strconv"
)

type Category struct {
}

func (cate Category) Add(c *gin.Context) {
	param := service.CreateCategoryRequest{}
	err := c.ShouldBind(&param)
	if err != nil {
		fmt.Println("add error:", err)
		return
	}
	s := service.New()
	err = s.AddCategory(&param)
	if err != nil {
		fmt.Println("add error:", err)
		return
	}
	data := gin.H{}
	c.JSON(http.StatusOK, data)
}

func (cate Category) Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println("get error:", err)
		return
	}
	s := service.New()
	category, err := s.GetCategory(id)
	if err != nil {
		fmt.Println("get error:", err)
		return
	}
	data := gin.H{
		"data": category,
	}
	c.JSON(http.StatusOK, data)
}

func (cate Category) List(c *gin.Context) {
	s := service.New()
	categories, err := s.CategoriesList()
	if err != nil {
		fmt.Println("list error:", err)
		return
	}
	data := gin.H{
		"list": categories,
	}
	c.JSON(http.StatusOK, data)
}

func (cate Category) Edit(c *gin.Context) {
	param := service.EditCategoryRequest{}
	err := c.ShouldBind(&param)
	if err != nil {
		fmt.Println("edit error:", err)
		return
	}
	s := service.New()
	rows, err := s.EditCategory(&param)
	if err != nil {
		fmt.Println("edit error:", err)
		return
	}
	data := gin.H{
		"rows": rows,
	}
	c.JSON(http.StatusOK, data)
}

func (cate Category) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println("delete error:", err)
		return
	}
	s := service.New()
	rows, err := s.DeleteCategory(uint(id))
	if err != nil {
		fmt.Println("delete error:", err)
		return
	}
	data := gin.H{
		"rows": rows,
	}
	c.JSON(http.StatusOK, data)
}
