package main

import (
	"log"
	"os"
	"os/exec"

	"github.com/awoelf/go-web-app/db"
	"github.com/awoelf/go-web-app/router"
	"github.com/awoelf/go-web-app/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
)

func loadEnv() {
    err := godotenv.Load()
	if err != nil {
		log.Fatal("Unable to load .env file.")
	}
}

func newApp() *fiber.App {
	dbConn, err := db.Connect()
	if err != nil {
		log.Panic(err)
	}
	defer dbConn.Client.Close()

	services.Register(dbConn.Client)

	engine := html.New("./views", ".html")
	f := fiber.New(fiber.Config{
		Views: engine,
	})
	f.Static("/", "./public")

	f.Mount("/", router.ViewsRouter())
	f.Mount("/api", router.APIRouter())

	return f
}

func main() {
	cmd := exec.Command("/bin/sh", "refresh.sh")
	
	loadEnv()
	port := os.Getenv("PORT")

	app := newApp()

	go cmd.Run()
	go log.Fatal(app.Listen(":" + port))
}
