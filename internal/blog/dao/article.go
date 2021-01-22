package dao

import (
	"github.com/weirubo/blog-go/internal/blog/model"
	"gorm.io/gorm"
)

func (d *Dao) CreateArticle(title, description, content, author, categoryName, tagName string, authorId, categoryId, tagId int) error {
	article := model.Article{
		Model:        gorm.Model{},
		Title:        title,
		Description:  description,
		Content:      content,
		AuthorId:     authorId,
		Author:       author,
		CategoryId:   categoryId,
		CategoryName: categoryName,
		TagId:        tagId,
		TagName:      tagName,
	}
	return article.Create(d.dbEngine)
}

func (d *Dao) GetArticleById(id int) (model.Article, error) {
	article := model.Article{}
	return article.GetById(d.dbEngine, id)
}

func (d *Dao) GetArticlesList() ([]model.Article, error) {
	article := model.Article{}
	return article.List(d.dbEngine)
}

func (d *Dao) UpdateArticle(id uint, title, description, content, author, categoryName, tagName string, authorId, categoryId, tagId int) (int64, error) {
	article := model.Article{
		Model: gorm.Model{
			ID: id,
		},
		Title:        title,
		Description:  description,
		Content:      content,
		AuthorId:     authorId,
		Author:       author,
		CategoryId:   categoryId,
		CategoryName: categoryName,
		TagId:        tagId,
		TagName:      tagName,
	}
	return article.Update(d.dbEngine)
}

func (d *Dao) DeleteArticle(id uint) (int64, error) {
	article := model.Article{
		Model: gorm.Model{
			ID: id,
		},
	}
	return article.Delete(d.dbEngine)
}
