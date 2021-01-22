package model

import (
	"fmt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null;index:idx_username_email;comment:用户名"`
	Email    string `gorm:"type:varchar(20);not null;unique;index:idx_username_email;comment:邮箱地址"`
	Password string `gorm:"type:varchar(20);not null;comment:密码"`
}

func (u User) Create(db *gorm.DB) error {
	// 迁移
	err := db.AutoMigrate(&u)
	if err != nil {
		fmt.Println("db AutoMigrate user table error:", err)
		return err
	}
	return db.Create(&u).Error
}

func (u User) GetByEmail(db *gorm.DB) (User, error) {
	user := User{}
	tx := db.Where(&u).Find(&user)
	if tx.Error != nil {
		return user, tx.Error
	}
	return user, nil
}

func (u User) GetById(db *gorm.DB, id uint) (User, error) {
	user := User{}
	tx := db.Where("id", id).Find(&user)
	if tx.Error != nil {
		return user, tx.Error
	}
	return user, nil
}

func (u User) List(db *gorm.DB, pageOffset, pageSize int) ([]User, error) {
	var users []User
	tx := db.Offset(pageOffset).Limit(pageSize).Find(&users)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return users, nil
}

func (u User) Update(db *gorm.DB) (int64, error) {
	tx := db.Where("id", u.ID).Updates(&u)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return tx.RowsAffected, nil
}

func (u User) Delete(db *gorm.DB) (int64, error) {
	tx := db.Where("id", u.ID).Delete(&User{})
	if tx.Error != nil {
		return 0, tx.Error
	}
	return tx.RowsAffected, nil
}
