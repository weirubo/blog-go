package service

import "github.com/weirubo/blog-go/internal/blog/model"

type CreateUserRequest struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	PassWord string `json:"password"`
}

type UserListByPageRequest struct {
	PageNum  int `json:"pageNum" form:"pageNum"`
	PageSize int `json:"pageSize" form:"pageSize"`
}

type EditUserRequest struct {
	CreateUserRequest
	ID uint `json:"id"`
}

func (s Service) AddUser(param *CreateUserRequest) error {
	return s.dao.CreateUser(param.UserName, param.Email, param.PassWord)
}

func (s Service) GetUser(id uint) (model.User, error) {
	return s.dao.GetUserByUid(id)
}

func (s Service) UserList(param *UserListByPageRequest) ([]model.User, error) {
	pageOffset := param.PageNum - 1
	return s.dao.GetUserList(pageOffset, param.PageSize)
}

func (s Service) EditUser(param *EditUserRequest) (int64, error) {
	return s.dao.UpdateUser(param.ID, param.UserName, param.Email, param.PassWord)
}

func (s Service) DeleteUser(id uint) (int64, error) {
	return s.dao.DeleteUser(id)
}

func (s Service) Login(param *CreateUserRequest) (model.User, error) {
	return s.dao.GetUserByEmail(param.Email, param.PassWord)
}
