package model

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name string `gorm:"type:varchar(20);not null;comment:分类名称"`
}

func (c Category) Create(db *gorm.DB) error {
	err := db.AutoMigrate(&c)
	if err != nil {
		return err
	}
	return db.Create(&c).Error
}

func (c Category) GetById(db *gorm.DB, id uint) (Category, error) {
	category := Category{}
	tx := db.Where("id", id).Find(&category)
	if tx.Error != nil {
		return category, tx.Error
	}
	return category, nil
}

func (c Category) List(db *gorm.DB) ([]Category, error) {
	var categories []Category
	tx := db.Find(&categories)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return categories, nil
}

func (c Category) Update(db *gorm.DB) (int64, error) {
	tx := db.Where("id", c.ID).Updates(&c)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return tx.RowsAffected, nil
}

func (c Category) Delete(db *gorm.DB) (int64, error) {
	tx := db.Where("id", c.ID).Delete(&Category{})
	if tx.Error != nil {
		return 0, tx.Error
	}
	return tx.RowsAffected, nil
}
