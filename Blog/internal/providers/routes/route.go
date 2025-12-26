package routes

import (
	articleRoutes "Blog/internal/modules/article/routes"
	homeRoutes "Blog/internal/modules/home/routes"
	userRoutes "Blog/internal/modules/user/routes"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	homeRoutes.Routes(router)
	articleRoutes.Routes(router)
	userRoutes.Routes(router)
}
