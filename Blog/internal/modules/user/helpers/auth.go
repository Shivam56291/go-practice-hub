package helpers

import (
	UserRepository "Blog/internal/modules/user/repositories"
	UserResponse "Blog/internal/modules/user/responses"
	"Blog/pkg/sessions"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) UserResponse.User {
	authID := sessions.Get(c, "auth")
	userID, _ := strconv.Atoi(authID)

	var userRepo = UserRepository.New()
	user := userRepo.FindByID(userID)

	if user.ID == 0 {
		return UserResponse.User{}
	}

	return UserResponse.ToUser(user)
}
