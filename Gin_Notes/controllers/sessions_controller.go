package controllers

import (
	"net/http"
	"notes_application/controllers/helpers_in"
	"notes_application/helpers"
	"notes_application/models"

	"github.com/gin-gonic/gin"
)

func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "home/login.html", gin.H{})
}

func SignupPage(c *gin.Context) {
	c.HTML(http.StatusOK, "home/signup.html", gin.H{})
}

func Signup(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	confirmPassword := c.PostForm("confirm_password")

	availabe := models.UserCheckAvailability(email)
	if !availabe {
		helpers_in.SetFlash(c, "Email already exists")
		c.HTML(http.StatusOK, "home/signup.html", helpers_in.SetPayload(c, gin.H{}))
		return
	}

	if password != confirmPassword {
		helpers_in.SetFlash(c, "Passwords do not match")
		c.HTML(http.StatusNotAcceptable, "home/signup.html", helpers_in.SetPayload(c, gin.H{}))
		return
	}

	user := models.UserCreate(email, password)
	if user.ID == 0 {
		helpers_in.SetFlash(c, "Unable to create user")
		c.HTML(http.StatusInternalServerError, "home/signup.html", helpers_in.SetPayload(c, gin.H{}))
	} else {
		helpers.SessionSet(c, user.ID)
		helpers_in.SetFlash(c, "Signup successful")
		c.Redirect(http.StatusSeeOther, "/")
	}
}

func Login(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	user := models.UserCheck(email, password)
	if user != nil { // valid login
		helpers.SessionSet(c, user.ID)
		helpers_in.SetFlash(c, "Logged in successfully")
		c.Redirect(http.StatusSeeOther, "/")
	} else { // invalid login
		helpers_in.SetFlash(c, "Invalid credentials")
		c.HTML(http.StatusOK, "home/login.html", helpers_in.SetPayload(c, gin.H{}))
	}
}

func Logout(c *gin.Context) {
	helpers.SessionClear(c)
	helpers_in.SetFlash(c, "Logged out successfully")
	c.Redirect(http.StatusSeeOther, "/login")
}
