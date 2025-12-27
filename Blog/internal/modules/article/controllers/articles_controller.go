package controllers

import (
	"Blog/internal/modules/article/requests/articles"
	ArticleService "Blog/internal/modules/article/services"
	"Blog/internal/modules/user/helpers"
	"Blog/pkg/converters"
	"Blog/pkg/errors"
	"Blog/pkg/html"
	"Blog/pkg/old"
	"Blog/pkg/sessions"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	articleService ArticleService.ArticleServiceInterface
}

func New() *Controller {
	return &Controller{
		articleService: ArticleService.New(),
	}
}

func (controller *Controller) Show(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		html.Render(c, http.StatusInternalServerError, "templates/errors/html/500", gin.H{
			"title":   "Internal Server Error",
			"message": "Error converting the id",
		})
		return
	}

	article, err := controller.articleService.Find(id)
	if err != nil {
		html.Render(c, http.StatusNotFound, "templates/errors/html/404", gin.H{
			"title":   "Page not Found",
			"message": "Article not found",
		})
		return
	}

	html.Render(c, http.StatusOK, "modules/article/html/show", gin.H{
		"title":   "Show Article",
		"article": article,
	})
}

func (controller *Controller) Create(c *gin.Context) {
	html.Render(c, http.StatusOK, "modules/article/html/create", gin.H{
		"title": "Create Article",
	})
}

func (controller *Controller) Store(c *gin.Context) {
	var request articles.StoreRequest

	if err := c.ShouldBind(&request); err != nil {
		errors.Init()
		errors.SetFromErrors(err)
		sessions.Set(c, "errors", converters.MapToString(errors.Get()))

		old.Init()
		old.Set(c)
		sessions.Set(c, "old", converters.UrlValuesToString(old.Get()))

		c.Redirect(http.StatusSeeOther, "/articles/create")
		return
	}

	user := helpers.Auth(c)

	article, err := controller.articleService.StoreAsUser(request, user)
	if err != nil {
		c.Redirect(http.StatusNotFound, "/articles/create")
		return
	}

	c.Redirect(http.StatusSeeOther, fmt.Sprint("/articles/", article.ID))
}
