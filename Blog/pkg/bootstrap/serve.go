package bootstrap

import (
	"Blog/pkg/config"
	"Blog/pkg/html"
	"Blog/pkg/routing"
	"Blog/pkg/static"
)

func Serve() {
	config.Set()

	routing.Init()

	static.LoadStatic(routing.GetRouter())

	html.LoadHTML(routing.GetRouter())

	routing.RegisterRoutes()

	routing.Serve()
}
