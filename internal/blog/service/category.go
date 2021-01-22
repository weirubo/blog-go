package service

import "github.com/weirubo/blog-go/internal/blog/model"

type CreateCategoryRequest struct {
	Name string `json:"name"`
}

type EditCategoryRequest struct {
	CreateCategoryRequest
	ID uint `json:"id"`
}

func (s Service) AddCategory(param *CreateCategoryRequest) error {
	return s.dao.CreateCategory(param.Name)
}

func (s Service) GetCategory(id uint) (model.Category, error) {
	return s.dao.GetCategoryById(id)
}

func (s Service) CategoriesList() ([]model.Category, error) {
	return s.dao.GetCategoriesList()
}

func (s Service) EditCategory(param *EditCategoryRequest) (int64, error) {
	return s.dao.UpdateCategory(param.ID, param.Name)
}

func (s Service) DeleteCategory(id uint) (int64, error) {
	return s.dao.DeleteCategory(id)
}
