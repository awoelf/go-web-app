package main

import (
	"log"
	"os/exec"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	cmd := exec.Command("/bin/sh", "refresh.sh")
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Static("/", "./public")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"title": "Add Your Comment!",
		}, "layouts/main")
	})

	go cmd.Run()
	go log.Fatal(app.Listen(":3000"))
}
