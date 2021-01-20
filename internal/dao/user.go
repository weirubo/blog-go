package dao

import (
	"github.com/weirubo/blog-go/internal/model"
	"gorm.io/gorm"
)

func (d *Dao) CreateUser(username, email, password string) error {
	user := model.User{
		Model:    gorm.Model{},
		Username: username,
		Email:    email,
		Password: password,
	}
	return user.CreateUser(d.dbEngine)
}
