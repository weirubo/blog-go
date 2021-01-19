package service

type CreateUserRequest struct {
	UserName string `json:"username"`
	Email string `json:"email"`
	PassWord string `json:"password"`
}

func (s Service) CreateUser(param CreateUserRequest) error {
	return nil
}
