package service

import "github.com/weirubo/blog-go/internal/blog/model"

type CreateTagRequest struct {
	Name string `json:"name"`
}

type EditTagRequest struct {
	CreateTagRequest
	ID uint `json:"id"`
}

func (s Service) AddTag(param *CreateTagRequest) error {
	return s.dao.CreateTag(param.Name)
}

func (s Service) GetTag(id uint) (model.Tag, error) {
	return s.dao.GetTagById(id)
}

func (s Service) TagsList() ([]model.Tag, error) {
	return s.dao.GetTagsList()
}

func (s Service) EditTag(param *EditTagRequest) (int64, error) {
	return s.dao.UpdateTag(param.ID, param.Name)
}

func (s Service) DeleteTag(id uint) (int64, error) {
	return s.dao.DeleteTag(id)
}
