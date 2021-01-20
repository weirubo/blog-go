package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// 创建数据库
func NewDBEngine() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/blog-go?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		TablePrefix:   "blog_",
		SingularTable: true,
	}})
	if err != nil {
		fmt.Println("NewDBEngine error:", err)
		return nil, err
	}
	return db, nil
}
