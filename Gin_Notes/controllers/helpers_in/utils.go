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

func IsUserLoggedIn(c *gin.Context) bool {
	return c.GetUint64("user_id") > 0
}

func SetPayload(c *gin.Context, h gin.H) gin.H {
	// Add email if user is logged in
	user := GetUserFromRequest(c)
	if user != nil && user.ID > 0 {
		h["email"] = user.Username
		h["logged_in"] = true
	} else {
		h["logged_in"] = false
	}

	// Add flash alert if exists
	flash := GetFlash(c)
	if len(flash) > 0 {
		h["alert"] = flash
	}

	return h
}
