package services

import (
	"Blog/internal/modules/user/request/auth"
	UserResponse "Blog/internal/modules/user/responses"
)

type UserServiceInterface interface {
	Create(request auth.RegisterRequest) (UserResponse.User, error)
	CheckUserExists(email string) bool
	HandleUserLogin(request auth.LoginRequest) (UserResponse.User, error)
}
