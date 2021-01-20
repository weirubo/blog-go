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

func (u User) GetUser(db *gorm.DB) (User, error) {
	user := User{}
	tx := db.Where(&u).Find(&user)
	if tx.Error != nil {
		return user, tx.Error
	}
	return user, nil
}

func (u User) GetUserByUid(db *gorm.DB, uid int) (User, error) {
	user := User{}
	tx := db.Where("id", uid).Find(&user)
	if tx.Error != nil {
		return user, tx.Error
	}
	return user, nil
}

func (u User) GetUserList(db *gorm.DB, pageOffset, pageSize int) ([]User, error) {
	var users []User
	tx := db.Offset(pageOffset).Limit(pageSize).Find(&users)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return users, nil
}

func (u User) UpdateUser(db *gorm.DB) (int64, error) {
	tx := db.Where("id", u.ID).Updates(&u)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return tx.RowsAffected, nil
}

func (u User) DeleteUser(db *gorm.DB) (int64, error) {
	tx := db.Where("id", u.ID).Delete(&User{})
	if tx.Error != nil {
		return 0, tx.Error
	}
	return tx.RowsAffected, nil
}
