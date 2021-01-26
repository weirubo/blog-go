package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/weirubo/blog-go/internal/blog/service"
	"net/http"
	"strconv"
)

type Tag struct {
}

func (t Tag) Add(c *gin.Context) {
	param := service.CreateTagRequest{}
	err := c.ShouldBind(&param)
	if err != nil {
		fmt.Println("add error:", err)
		return
	}
	s := service.New()
	err = s.AddTag(&param)
	if err != nil {
		fmt.Println("add error:", err)
		return
	}
	data := gin.H{}
	c.JSON(http.StatusOK, data)
}

func (t Tag) Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println("get error:", err)
		return
	}
	s := service.New()
	tag, err := s.GetTag(uint(id))
	if err != nil {
		fmt.Println("get error:", err)
		return
	}
	data := gin.H{
		"data": tag,
	}
	c.JSON(http.StatusOK, data)
}

func (t Tag) List(c *gin.Context) {
	s := service.New()
	tags, err := s.TagsList()
	if err != nil {
		fmt.Println("list error:", err)
		return
	}
	data := gin.H{
		"list": tags,
	}
	c.JSON(http.StatusOK, data)
}

func (t Tag) Edit(c *gin.Context) {
	param := service.EditTagRequest{}
	err := c.ShouldBindHeader(&param)
	if err != nil {
		fmt.Println("edit error:", err)
		return
	}
	s := service.New()
	rows, err := s.EditTag(&param)
	if err != nil {
		fmt.Println("edit error:", err)
		return
	}
	data := gin.H{
		"rows": rows,
	}
	c.JSON(http.StatusOK, data)
}

func (t Tag) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println("delete error:", err)
		return
	}
	s := service.New()
	rows, err := s.DeleteTag(uint(id))
	if err != nil {
		fmt.Println("delete error:", err)
		return
	}
	data := gin.H{
		"rows": rows,
	}
	c.JSON(http.StatusOK, data)
}
