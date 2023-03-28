package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/vvk17/microservices-go/services/web-blog/database"
)

func main() {
	app := fiber.New()
	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello World, 1")
	})
	database.ConnectDB()
	log.Fatal(app.Listen(":3069"))
}