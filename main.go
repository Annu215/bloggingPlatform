package main

import (
	"github.com/gofiber/fiber/v2"
	"main.go/config"
	"main.go/router"
)

func main() {
	config.Init()
	config.ConnectDB()
	app := fiber.New()
	router.Configure(app)
	app.Listen(":7000")
}
