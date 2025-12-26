package controllers

import (
	"Blog/internal/modules/user/request/auth"
	"Blog/pkg/converters"
	"Blog/pkg/errors"
	"Blog/pkg/html"
	"Blog/pkg/sessions"
	"log"
	"net/http"

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

		c.Redirect(http.StatusFound, "/register")
		return
	}

	user, err := controller.UserService.Create(request)
	if err != nil {
		c.Redirect(http.StatusFound, "/register")
		return
	}

	log.Printf("user: %v", user)
	c.Redirect(http.StatusFound, "/")
}
