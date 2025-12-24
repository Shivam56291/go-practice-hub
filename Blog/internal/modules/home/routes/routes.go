package routes

import (
	"Blog/pkg/html"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		html.Render(c, http.StatusOK, "modules/home/html/home", gin.H{
			"title": "Home page",
		})
	})

	router.GET("/about", func(c *gin.Context) {
		html.Render(c, http.StatusOK, "modules/home/html/about", gin.H{
			"title": "About page",
		})
	})
}
