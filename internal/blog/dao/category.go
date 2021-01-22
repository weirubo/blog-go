package dao

import (
	"github.com/weirubo/blog-go/internal/blog/model"
	"gorm.io/gorm"
)

func (d *Dao) CreateCategory(name string) error {
	category := model.Category{
		Model: gorm.Model{},
		Name:  name,
	}
	return category.Create(d.dbEngine)
}

func (d *Dao) GetCategoryById(id int) (model.Category, error) {
	category := model.Category{}
	return category.GetById(d.dbEngine, id)
}

func (d *Dao) GetCategoriesList() ([]model.Category, error) {
	category := model.Category{}
	return category.List(d.dbEngine)
}

func (d *Dao) UpdateCategory(id uint, name string) (int64, error) {
	category := model.Category{
		Model: gorm.Model{
			ID: id,
		},
		Name: name,
	}
	return category.Update(d.dbEngine)
}

func (d *Dao) DeleteCategory(id uint) (int64, error) {
	category := model.Category{
		Model: gorm.Model{
			ID: id,
		},
	}
	return category.Delete(d.dbEngine)
}
