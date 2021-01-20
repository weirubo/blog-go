package model

import (
	"fmt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"type:varchar(20);not null;index:idx_username_email;comment:用户名"`
	Email    string `json:"email" gorm:"type:varchar(20);not null;unique;index:idx_username_email;comment:邮箱地址"`
	Password string `json:"password" gorm:"type:varchar(20);not null;comment:密码"`
}

func (u User) CreateUser(db *gorm.DB) error {
	// 迁移
	err := db.AutoMigrate(&u)
	if err != nil {
		fmt.Println("db AutoMigrate user table error:", err)
		return err
	}
	return db.Create(&u).Error
}
