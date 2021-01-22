package service

import "github.com/weirubo/blog-go/internal/blog/model"

type CreateArticleRequest struct {
	Title        string `json:"title"`
	Description  string `json:"description"`
	Content      string `json:"content"`
	AuthorId     int    `json:"authorId"`
	Author       string `json:"author"`
	CategoryId   int    `json:"categoryId"`
	CategoryName string `json:"categoryName"`
	TagId        int    `json:"tagId"`
	TagName      string `json:"tagName"`
}

type EditArticleRequest struct {
	CreateArticleRequest
	ID uint `json:"id"`
}

func (s Service) AddArticle(param *CreateArticleRequest) error {
	return s.dao.CreateArticle(param.Title, param.Description, param.Content, param.Author, param.CategoryName, param.TagName, param.AuthorId, param.CategoryId, param.TagId)
}

func (s Service) GetArticle(id uint) (model.Article, error) {
	return s.dao.GetArticleById(id)
}

func (s Service) ArticlesList() ([]model.Article, error) {
	return s.dao.GetArticlesList()
}

func (s Service) EditArticle(param *EditArticleRequest) (int64, error) {
	return s.dao.UpdateArticle(param.ID, param.Title, param.Description, param.Content, param.Author, param.CategoryName, param.TagName, param.AuthorId, param.CategoryId, param.TagId)
}

func (s Service) DeleteArticle(id uint) (int64, error) {
	return s.dao.DeleteArticle(id)
}
