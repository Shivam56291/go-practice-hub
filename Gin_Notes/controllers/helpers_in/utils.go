package helpers_in

import (
	"notes_application/models"

	"github.com/gin-gonic/gin"
)

func GetUserFromRequest(c *gin.Context) *models.User {
	UserID := c.GetUint64("user_id")
	var currentUser *models.User
	if UserID > 0 {
		currentUser = models.UserFind(UserID)
	} else {
		currentUser = nil
	}
	return currentUser
}
