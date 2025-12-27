package services

import (
	UserModel "Blog/internal/modules/user/models"
	UserRepository "Blog/internal/modules/user/repositories"
	"Blog/internal/modules/user/request/auth"
	UserResponse "Blog/internal/modules/user/responses"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepository UserRepository.UserRepositoryInterface
}

func New() *UserService {
	return &UserService{
		UserRepository: UserRepository.New(),
	}
}

func (userService *UserService) Create(request auth.RegisterRequest) (UserResponse.User, error) {
	var response UserResponse.User
	var user UserModel.User

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return response, errors.New("error in hashing the password")
	}

	user.Name = request.Name
	user.Email = request.Email
	user.Password = string(hashedPassword)

	newUser := userService.UserRepository.Create(user)
	if newUser.ID == 0 {
		return response, errors.New("error in creating the user")
	}

	return UserResponse.ToUser(newUser), nil
}

func (userService *UserService) CheckUserExists(email string) bool {
	return userService.UserRepository.FindByEmail(email).ID != 0
}

func (userService *UserService) HandleUserLogin(request auth.LoginRequest) (UserResponse.User, error) {
	var response UserResponse.User
	existsUser := userService.UserRepository.FindByEmail(request.Email)
	if existsUser.ID == 0 {
		return response, errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(existsUser.Password), []byte(request.Password)); err != nil {
		return response, errors.New("invalid credentials")
	}

	return UserResponse.ToUser(existsUser), nil
}
