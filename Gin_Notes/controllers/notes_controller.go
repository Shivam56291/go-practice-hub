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
		c.HTML(http.StatusUnauthorized, "home/index.html", helpers_in.SetPayload(c, gin.H{}))
		return
	}
	notes := models.NotesAll(currentUser)
	c.HTML(http.StatusOK, "notes/index.html", helpers_in.SetPayload(c, gin.H{
		"notes": notes,
		"email": currentUser.Username,
	}))
}

func NotesNew(c *gin.Context) {
	c.HTML(http.StatusOK, "notes/new.html", gin.H{})
}

type FormData struct {
	Name    string `form:"name"`
	Content string `form:"content"`
}

func NotesCreate(c *gin.Context) {
	currentUser := helpers_in.GetUserFromRequest(c)
	if currentUser == nil || currentUser.ID == 0 {
		helpers_in.SetFlash(c, "Please login to view your notes")
		c.HTML(http.StatusUnauthorized, "home/index.html", helpers_in.SetPayload(c, gin.H{}))
		return
	}
	var data FormData
	c.Bind(&data)
	models.NotesCreate(currentUser, data.Name, data.Content)
	c.Redirect(http.StatusMovedPermanently, "/notes")
}

func NotesShow(c *gin.Context) {
	currentUser := helpers_in.GetUserFromRequest(c)
	if currentUser == nil || currentUser.ID == 0 {
		helpers_in.SetFlash(c, "Please login to view your notes")
		c.HTML(http.StatusUnauthorized, "home/index.html", helpers_in.SetPayload(c, gin.H{}))
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
		c.HTML(http.StatusUnauthorized, "home/index.html", helpers_in.SetPayload(c, gin.H{}))
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
		c.HTML(http.StatusUnauthorized, "home/index.html", helpers_in.SetPayload(c, gin.H{}))
		return
	}

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	var data FormData
	c.Bind(&data)
	note := models.NotesFind(currentUser, id)
	note.Update(data.Name, data.Content)
	c.Redirect(http.StatusSeeOther, "/notes/"+idStr)
}

func NotesDelete(c *gin.Context) {
	currentUser := helpers_in.GetUserFromRequest(c)
	if currentUser == nil || currentUser.ID == 0 {
		helpers_in.SetFlash(c, "Please login to view your notes")
		c.HTML(http.StatusUnauthorized, "home/index.html", helpers_in.SetPayload(c, gin.H{}))
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
