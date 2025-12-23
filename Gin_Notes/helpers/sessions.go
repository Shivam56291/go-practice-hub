package helpers

import "github.com/gin-gonic/gin"

func SessionSet(c *gin.Context, userID uint64){}

func SessionGet(c *gin.Context) uint64{
	return 0
}

func SessionClear(c *gin.Context){}