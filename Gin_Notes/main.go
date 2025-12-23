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
	r.GET("/notes/new", controllers.NotesNew)
	r.POST("/notes", controllers.NotesCreate)
	r.GET("/notes/:id", controllers.NotesShow)
	r.GET("/notes/edit/:id", controllers.NotesEditPage)
	r.POST("/notes/:id", controllers.NotesUpdate)
	r.DELETE("/notes/:id", controllers.NotesDelete)

	r.GET("/signup", controllers.SignupPage)

	r.GET("/login", controllers.LoginPage)

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home/index.html", gin.H{
			"title": "Notes Application",
		})
	})

	log.Println("Listening and serving HTTP on :8080")
	r.Run(":8080")
}
