package main

import (
	"log"
	"net/http"
	"notes_application/controllers"
	"notes_application/middlewares"
	"notes_application/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Static("/vendor", "./static/vendor")

	r.LoadHTMLGlob("templates/**/**")

	models.ConnectDatabase()
	models.DBMigrate()

	store := memstore.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("notes", store))

	r.Use(middlewares.AuthenticateUser())

	r.GET("/notes", controllers.NotesIndex)
	r.GET("/notes/new", controllers.NotesNew)
	r.POST("/notes", controllers.NotesCreate)
	r.GET("/notes/:id", controllers.NotesShow)
	r.GET("/notes/edit/:id", controllers.NotesEditPage)
	r.POST("/notes/:id", controllers.NotesUpdate)
	r.DELETE("/notes/:id", controllers.NotesDelete)

	r.GET("/signup", controllers.SignupPage)
	r.GET("/login", controllers.LoginPage)

	r.POST("/signup", controllers.Signup)

	r.POST("/login", controllers.Login)
	r.GET("/logout", controllers.Logout)

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home/index.html", gin.H{
			"title":     "Notes Application",
			"logged_in": (c.GetUint64("user_id") > 0),
		})
	})

	log.Println("Listening and serving HTTP on :8080")
	r.Run(":8080")
}
