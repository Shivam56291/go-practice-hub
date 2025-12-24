package main

import (
	"log"
	"net/http"
	"notes_application/controllers"
	"notes_application/controllers/helpers_in"
	"notes_application/middlewares"
	"notes_application/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Static("/vendor", "./static/vendor")
	r.Static("/assets", "./static")

	r.LoadHTMLGlob("templates/**/**")

	models.ConnectDatabase()
	models.DBMigrate()

	store := memstore.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("notes", store))

	r.Use(middlewares.AuthenticateUser())

	notes := r.Group("/notes")
	notes.GET("", controllers.NotesIndex)
	notes.GET("/new", controllers.NotesNew)
	notes.POST("", controllers.NotesCreate)
	notes.GET("/:id", controllers.NotesShow)
	notes.GET("/edit/:id", controllers.NotesEditPage)
	notes.POST("/:id", controllers.NotesUpdate)
	notes.DELETE("/:id", controllers.NotesDelete)

	r.GET("/signup", controllers.SignupPage)
	r.GET("/login", controllers.LoginPage)

	r.POST("/signup", controllers.Signup)

	r.POST("/login", controllers.Login)
	r.GET("/logout", controllers.Logout)

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home/index.html", helpers_in.SetPayload(c, gin.H{
			"title": "Notes Application",
		}))
	})

	log.Println("Listening and serving HTTP on :8080")
	r.Run(":8080")
}
