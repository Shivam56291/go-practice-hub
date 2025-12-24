package controllers

import (
	"fmt"
	"net/http"
	"notes_application/controllers/helpers_in"
	"notes_application/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func NotesIndex(c *gin.Context) {
	currentUser := helpers_in.GetUserFromRequest(c)
	if currentUser == nil || currentUser.ID == 0 {
		helpers_in.SetFlash(c, "Please login to view your notes")
		c.HTML(http.StatusUnauthorized, "home/index.html", gin.H{})
		return
	}
	notes := models.NotesAll(currentUser)
	c.HTML(http.StatusOK, "notes/index.html", gin.H{
		"notes": notes,
	})
}

func NotesNew(c *gin.Context) {
	c.HTML(http.StatusOK, "notes/new.html", gin.H{})
}

func NotesCreate(c *gin.Context) {
	currentUser := helpers_in.GetUserFromRequest(c)
	if currentUser == nil || currentUser.ID == 0 {
		helpers_in.SetFlash(c, "Please login to view your notes")
		c.HTML(http.StatusUnauthorized, "home/index.html", gin.H{})
		return
	}
	name := c.PostForm("name")
	content := c.PostForm("content")
	models.NotesCreate(currentUser, name, content)
	c.Redirect(http.StatusMovedPermanently, "/notes")
}

func NotesShow(c *gin.Context) {
	currentUser := helpers_in.GetUserFromRequest(c)
	if currentUser == nil || currentUser.ID == 0 {
		helpers_in.SetFlash(c, "Please login to view your notes")
		c.HTML(http.StatusUnauthorized, "home/index.html", gin.H{})
		return
	}
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	note := models.NotesFind(currentUser, id)
	c.HTML(http.StatusOK, "notes/show.html", gin.H{
		"note": note,
	})
}

func NotesEditPage(c *gin.Context) {
	currentUser := helpers_in.GetUserFromRequest(c)
	if currentUser == nil || currentUser.ID == 0 {
		helpers_in.SetFlash(c, "Please login to view your notes")
		c.HTML(http.StatusUnauthorized, "home/index.html", gin.H{})
		return
	}
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	note := models.NotesFind(currentUser, id)
	c.HTML(http.StatusOK, "notes/edit.html", gin.H{
		"note": note,
	})
}

func NotesUpdate(c *gin.Context) {
	currentUser := helpers_in.GetUserFromRequest(c)
	if currentUser == nil || currentUser.ID == 0 {
		helpers_in.SetFlash(c, "Please login to view your notes")
		c.HTML(http.StatusUnauthorized, "home/index.html", gin.H{})
		return
	}
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	note := models.NotesFind(currentUser, id)
	note.Update(c.PostForm("name"), c.PostForm("content"))
	c.Redirect(http.StatusSeeOther, "/notes/"+idStr)
}

func NotesDelete(c *gin.Context) {
	currentUser := helpers_in.GetUserFromRequest(c)
	if currentUser == nil || currentUser.ID == 0 {
		helpers_in.SetFlash(c, "Please login to view your notes")
		c.HTML(http.StatusUnauthorized, "home/index.html", gin.H{})
		return
	}
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	models.NotesMarkDeleted(currentUser, id)
	c.Redirect(http.StatusSeeOther, "/notes")
}
