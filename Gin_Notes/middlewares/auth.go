package middlewares

import (
	"fmt"
	"notes_application/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthenticateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		sessionID := session.Get("user_id")
		fmt.Printf("DEBUG: Session ID retrieved: %v (Type: %T)\n", sessionID, sessionID)

		var user *models.User
		userPresent := true
		if sessionID == nil {
			userPresent = false
			fmt.Println("DEBUG: Session ID is nil")
		} else {
			// Handle potential type mismatch if underlying store changes numbers
			switch id := sessionID.(type) {
			case uint64:
				user = models.UserFind(id)
			case int:
				user = models.UserFind(uint64(id))
			case float64: // JSON numbers sometimes come as float64
				user = models.UserFind(uint64(id))
			default:
				fmt.Printf("DEBUG: Unexpected type for sessionID: %T\n", id)
				userPresent = false
			}
			
			if user != nil {
				fmt.Printf("DEBUG: User found: ID=%d, Name=%s\n", user.ID, user.Username)
				userPresent = (user.ID > 0)
			} else {
				fmt.Println("DEBUG: UserFind returned nil")
				userPresent = false
			}
		}
		if userPresent {
			c.Set("user_id", user.ID)
			c.Set("email", user.Username)
		} else {
			fmt.Println("DEBUG: User not present or invalid")
		}
		c.Next()
	}
}
