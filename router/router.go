package router

import (
	"github.com/awoelf/go-web-app/controllers"
	"github.com/gofiber/fiber/v2"
)

func APIRouter() *fiber.App {
	api := fiber.New()

	api.Get("/comments", func(c *fiber.Ctx) error {
		return controllers.GetAllComments(c)
	})
	api.Get("/comment/:id", func(c *fiber.Ctx) error {
		return controllers.GetComment(c)
	})
	api.Post("/comment", func(c *fiber.Ctx) error {
		return controllers.CreateComment(c)
	})
	api.Put("/comment/:id", func(c *fiber.Ctx) error {
		return controllers.UpdateComment(c)
	})
	api.Delete("/comment/:id", func(c *fiber.Ctx) error {
		return controllers.DeleteComment(c)
	})

	return api
}