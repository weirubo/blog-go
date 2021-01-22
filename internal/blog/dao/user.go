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
	return user.Create(d.dbEngine)
}

func (d *Dao) GetUserByEmail(email, password string) (model.User, error) {
	user := model.User{
		Email:    email,
		Password: password,
	}
	return user.GetByEmail(d.dbEngine)
}

func (d *Dao) GetUserByUid(id uint) (model.User, error) {
	user := model.User{}
	return user.GetById(d.dbEngine, id)
}

func (d *Dao) GetUserList(pageOffset, pageSize int) ([]model.User, error) {
	user := model.User{}
	return user.List(d.dbEngine, pageOffset, pageSize)
}

func (d *Dao) UpdateUser(id uint, username, email, password string) (int64, error) {
	user := model.User{
		Model: gorm.Model{
			ID: id,
		},
		Username: username,
		Email:    email,
		Password: password,
	}
	return user.Update(d.dbEngine)
}

func (d *Dao) DeleteUser(id uint) (int64, error) {
	user := model.User{
		Model: gorm.Model{
			ID: id,
		},
	}
	return user.Delete(d.dbEngine)
}
