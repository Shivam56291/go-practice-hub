package controllers

import (
	"fmt"
	"net/http"
	"notes_application/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func NotesIndex(c *gin.Context) {
	notes := models.NotesAll()
	c.HTML(http.StatusOK, "notes/index.html", gin.H{
		"notes": notes,
	})
}

func NotesNew(c *gin.Context) {
	c.HTML(http.StatusOK, "notes/new.html", gin.H{})
}

func NotesCreate(c *gin.Context) {
	name := c.PostForm("name")
	content := c.PostForm("content")
	models.NotesCreate(name, content)
	c.Redirect(http.StatusMovedPermanently, "/notes")
}

func NotesShow(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	note := models.NotesFind(id)
	c.HTML(http.StatusOK, "notes/show.html", gin.H{
		"note": note,
	})
}

func NotesEditPage(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	note := models.NotesFind(id)
	c.HTML(http.StatusOK, "notes/edit.html", gin.H{
		"note": note,
	})
}

func NotesUpdate(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	note := models.NotesFind(id)
	note.Update(c.PostForm("name"), c.PostForm("content"))
	c.Redirect(http.StatusSeeOther, "/notes/"+idStr)
}

func NotesDelete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	models.NotesMarkDeleted(id)
	c.Redirect(http.StatusSeeOther, "/notes")
}
