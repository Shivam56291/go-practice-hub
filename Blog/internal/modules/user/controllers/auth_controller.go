package controllers

import (
	"Blog/internal/modules/user/request/auth"
	"Blog/pkg/converters"
	"Blog/pkg/errors"
	"Blog/pkg/html"
	"Blog/pkg/old"
	"Blog/pkg/sessions"
	"log"
	"net/http"
	"strconv"

	UserService "Blog/internal/modules/user/services"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	UserService UserService.UserServiceInterface
}

func New() *Controller {
	return &Controller{
		UserService: UserService.New(),
	}
}

func (controller *Controller) Register(c *gin.Context) {
	html.Render(c, http.StatusOK, "modules/user/html/register", gin.H{
		"title": "Register",
	})
}

func (controller *Controller) HandleRegister(c *gin.Context) {
	var request auth.RegisterRequest
	if err := c.ShouldBind(&request); err != nil {
		errors.Init()
		errors.SetFromErrors(err)

		sessions.Set(c, "errors", converters.MapToString(errors.Get()))

		old.Init()
		old.Set(c)
		sessions.Set(c, "old", converters.UrlValuesToString(old.Get()))

		c.Redirect(http.StatusFound, "/register")
		return
	}

	if controller.UserService.CheckUserExists(request.Email) {
		errors.Init()
		errors.Add("email", "User with this email already exists")

		sessions.Set(c, "errors", converters.MapToString(errors.Get()))

		old.Init()
		old.Set(c)
		sessions.Set(c, "old", converters.UrlValuesToString(old.Get()))

		c.Redirect(http.StatusFound, "/register")
		return
	}

	user, err := controller.UserService.Create(request)
	if err != nil {
		c.Redirect(http.StatusFound, "/register")
		return
	}

	sessions.Set(c, "auth", strconv.Itoa(int(user.ID)))

	log.Printf("user: %v", user)
	c.Redirect(http.StatusFound, "/")
}

func (controller *Controller) Login(c *gin.Context) {
	html.Render(c, http.StatusOK, "modules/user/html/login", gin.H{
		"title": "Login",
	})
}

func (controller *Controller) HandleLogin(c *gin.Context) {
	var request auth.LoginRequest

	if err := c.ShouldBind(&request); err != nil {
		errors.Init()
		errors.SetFromErrors(err)

		sessions.Set(c, "errors", converters.MapToString(errors.Get()))

		old.Init()
		old.Set(c)
		sessions.Set(c, "old", converters.UrlValuesToString(old.Get()))

		c.Redirect(http.StatusFound, "/login")
		return
	}

	user, err := controller.UserService.HandleUserLogin(request)
	if err != nil {
		errors.Init()
		errors.Add("email", err.Error())

		sessions.Set(c, "errors", converters.MapToString(errors.Get()))

		old.Init()
		old.Set(c)
		sessions.Set(c, "old", converters.UrlValuesToString(old.Get()))

		c.Redirect(http.StatusFound, "/login")
		return
	}

	sessions.Set(c, "auth", strconv.Itoa(int(user.ID)))

	log.Printf("user: %v", user)
	c.Redirect(http.StatusFound, "/")
}

func (controller *Controller) HandleLogout(c *gin.Context) {
	sessions.Remove(c, "auth")
	c.Redirect(http.StatusFound, "/")
}
