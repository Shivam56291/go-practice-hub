package main

import (
	"log"
	"net/http"
	"notes_application/controllers"
	"notes_application/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Static("/vendor", "./static/vendor")

	r.LoadHTMLGlob("templates/**/**")

	models.ConnectDatabase()
	models.DBMigrate()

	r.GET("/notes", controllers.NotesIndex)

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "views/index.html", gin.H{
			"title": "Notes Application",
		})
	})

	log.Println("Listening and serving HTTP on :8080")
	r.Run(":8080")
}
