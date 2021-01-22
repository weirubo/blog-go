package dao

import (
	"github.com/weirubo/blog-go/internal/blog/model"
	"gorm.io/gorm"
)

func (d *Dao) CreateTag(name string) error {
	tag := model.Tag{
		Model: gorm.Model{},
		Name:  name,
	}
	return tag.Create(d.dbEngine)
}

func (d *Dao) GetTagById(id uint) (model.Tag, error) {
	tag := model.Tag{}
	return tag.GetById(d.dbEngine, id)
}

func (d *Dao) GetTagsList() ([]model.Tag, error) {
	tag := model.Tag{}
	return tag.List(d.dbEngine)
}

func (d *Dao) UpdateTag(id uint, name string) (int64, error) {
	tag := model.Tag{
		Model: gorm.Model{
			ID: id,
		},
		Name: name,
	}
	return tag.Update(d.dbEngine)
}

func (d *Dao) DeleteTag(id uint) (int64, error) {
	tag := model.Tag{
		Model: gorm.Model{
			ID: id,
		},
	}
	return tag.Delete(d.dbEngine)
}
