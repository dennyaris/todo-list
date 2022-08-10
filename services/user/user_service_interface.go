package services

import "github.com/dennyaris/todo-list/entities"

type UserService interface {
	Register(data entities.UserRequest) error
	Login(data entities.LoginRequest) (string, error)
	GetProfile(user_id uint) (entities.User, error)
}
