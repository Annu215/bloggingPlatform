package router

import (
	"github.com/gofiber/fiber/v2"
	"main.go/modules/blog"
)

func Configure(app *fiber.App) {
	api := app.Group("/api")
	blog.Route(api)
}
