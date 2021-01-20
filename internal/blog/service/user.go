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

type ModifyUserRequest struct {
	CreateUserRequest
	ID uint `json:"uid"`
}

func (s Service) Register(param *CreateUserRequest) error {
	return s.dao.CreateUser(param.UserName, param.Email, param.PassWord)
}

func (s Service) Login(param *CreateUserRequest) (model.User, error) {
	return s.dao.GetUser(param.Email, param.PassWord)
}

func (s Service) User(uid int) (model.User, error) {
	return s.dao.GetUserByUid(uid)
}

func (s Service) List(param *UserListByPageRequest) ([]model.User, error) {
	pageOffset := param.PageNum - 1
	return s.dao.GetUserList(pageOffset, param.PageSize)
}

func (s Service) Modify(param *ModifyUserRequest) (int64, error) {
	return s.dao.UpdateUser(param.ID, param.UserName, param.Email, param.PassWord)
}

func (s Service) Delete(uid uint) (int64, error) {
	return s.dao.DeleteUser(uid)
}
