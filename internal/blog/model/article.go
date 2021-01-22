package model

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title        string `gorm:"type:varchar(60);not null;comment:标题"`
	Description  string `gorm:"type:varchar(600);not null;comment:摘要"`
	Content      string `gorm:"type:text;not null;comment:内容"`
	AuthorId     int    `gorm:"type:int(10);unsigned;not null;comment:作者 ID"`
	Author       string `gorm:"type:varchar(20);not null;comment:作者"`
	CategoryId   int    `gorm:"type:tinyint(3);not null;comment:分类 ID"`
	CategoryName string `gorm:"type:varchar(20);not null;comment:分类名称"`
	TagId        int    `gorm:"type:int(10);unsigned;not null;comment:标签 ID"`
	TagName      string `gorm:"type:varchar(20);not null;comment:标签名称"`
}

func (a Article) Create(db *gorm.DB) error {
	err := db.AutoMigrate(&a)
	if err != nil {
		return err
	}
	return db.Create(&a).Error
}

func (a Article) GetById(db *gorm.DB, id int) (Article, error) {
	article := Article{}
	tx := db.Where("id", id).Find(&article)
	if tx.Error != nil {
		return article, tx.Error
	}
	return article, nil
}

func (a Article) List(db *gorm.DB) ([]Article, error) {
	var articles []Article
	tx := db.Find(&articles)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return articles, nil
}

func (a Article) Update(db *gorm.DB) (int64, error) {
	tx := db.Where("id", a.ID).Updates(&a)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return tx.RowsAffected, nil
}

func (a Article) Delete(db *gorm.DB) (int64, error) {
	tx := db.Where("id", a.ID).Delete(&Article{})
	if tx.Error != nil {
		return 0, tx.Error
	}
	return tx.RowsAffected, nil
}
