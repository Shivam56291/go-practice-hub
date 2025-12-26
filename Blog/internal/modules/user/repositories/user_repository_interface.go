package repositories

import userModel "Blog/internal/modules/user/models"

type UserRepositoryInterface interface {
	Create(user userModel.User) userModel.User
}
