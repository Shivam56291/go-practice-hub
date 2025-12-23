package controllers

import (
	"net/http"
	"notes_application/models"

	"github.com/gin-gonic/gin"
)

func NotesIndex(c *gin.Context) {
	notes := models.NotesAll()
	c.HTML(http.StatusOK, "notes/index.html", gin.H{
		"notes": notes,
	})
}
