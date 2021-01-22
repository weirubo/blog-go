package model

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	Name string `gorm:"type:varchar(20);not null;comment:标签名称"`
}

func (t Tag) Create(db *gorm.DB) error {
	err := db.AutoMigrate(&t)
	if err != nil {
		return err
	}
	return db.Create(&t).Error
}

func (t Tag) GetById(db *gorm.DB, id uint) (Tag, error) {
	tag := Tag{}
	tx := db.Where("id", id).Find(&tag)
	if tx.Error != nil {
		return tag, tx.Error
	}
	return tag, nil
}

func (t Tag) List(db *gorm.DB) ([]Tag, error) {
	var tags []Tag
	tx := db.Find(&tags)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return tags, nil
}

func (t Tag) Update(db *gorm.DB) (int64, error) {
	tx := db.Where("id", t.ID).Updates(&t)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return tx.RowsAffected, nil
}

func (t Tag) Delete(db *gorm.DB) (int64, error) {
	tx := db.Where("id", t.ID).Delete(&Tag{})
	if tx.Error != nil {
		return 0, tx.Error
	}
	return tx.RowsAffected, nil
}
