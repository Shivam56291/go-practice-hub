package helpers_in

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// SetFlash sets a one-time flash alert in the session
func SetFlash(c *gin.Context, message string) {
	session := sessions.Default(c)
	session.Set("flash_alert", message)
	session.Save()
}

// GetFlash retrieves and clears the flash alert from the session
func GetFlash(c *gin.Context) string {
	session := sessions.Default(c)
	flash := session.Get("flash_alert")
	if flash != nil {
		session.Delete("flash_alert")
		session.Save()
		return flash.(string)
	}
	return ""
}
