package bootstrap

import (
	"Blog/pkg/config"
	"Blog/pkg/database"
	"Blog/pkg/html"
	"Blog/pkg/routing"
	"Blog/pkg/sessions"
	"Blog/pkg/static"
)

func Serve() {
	config.Set()

	database.Connect()

	routing.Init()

	sessions.Start(routing.GetRouter())

	static.LoadStatic(routing.GetRouter())

	html.LoadHTML(routing.GetRouter())

	routing.RegisterRoutes()

	routing.Serve()
}
