package dao

import (
	"github.com/weirubo/blog-go/internal/blog/model"
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

func (d *Dao) GetUser(email, password string) (model.User, error) {
	user := model.User{
		Email:    email,
		Password: password,
	}
	return user.GetUser(d.dbEngine)
}

func (d *Dao) GetUserByUid(uid int) (model.User, error) {
	user := model.User{}
	return user.GetUserByUid(d.dbEngine, uid)
}

func (d *Dao) GetUserList(pageOffset, pageSize int) ([]model.User, error) {
	user := model.User{}
	return user.GetUserList(d.dbEngine, pageOffset, pageSize)
}

func (d *Dao) UpdateUser(uid uint, username, email, password string) (int64, error) {
	user := model.User{
		Model: gorm.Model{
			ID: uid,
		},
		Username: username,
		Email:    email,
		Password: password,
	}
	return user.UpdateUser(d.dbEngine)
}

func (d *Dao) DeleteUser(uid uint) (int64, error) {
	user := model.User{
		Model: gorm.Model{
			ID: uid,
		},
	}
	return user.DeleteUser(d.dbEngine)
}
