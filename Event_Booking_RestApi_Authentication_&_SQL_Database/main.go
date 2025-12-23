package main

import (
	"event-booking-api/db"
	"event-booking-api/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Could not load .env file")
	}

	err = db.InitDB()
	if err != nil {
		panic("Could not initialize database: " + err.Error())
	}

	server := gin.Default()

	routes.RegisterRoutes(server)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server.Run(":" + port)
}
