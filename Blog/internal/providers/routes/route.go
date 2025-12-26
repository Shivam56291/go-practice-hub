package routes

import (
	homeRoutes "Blog/internal/modules/home/routes"

	articleRoutes "Blog/internal/modules/article/routes"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	homeRoutes.Routes(router)
	articleRoutes.Routes(router)
}
