package blog

import "github.com/gofiber/fiber/v2"

func Route(router fiber.Router) {
	group := router.Group("/blog")
	group.Post("/create", create)
	group.Post("/delete", delete)
	group.Post("/update", update)
	group.Post("/fetchOne", fetchOne)
	group.Post("/fetchAll", fetchAll)
}
