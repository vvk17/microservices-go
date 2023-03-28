package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/vvk17/microservices-go/services/web-blog/database"
	"github.com/vvk17/microservices-go/services/web-blog/models"
)

func main() {
	app := fiber.New()
	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello World, 1")
	})
	database.ConnectDB()
	log.Print("ConnectDB - success")
	orm.RegisterModel(new(models.Authors))
	log.Print("orm.RegisterModel - success")
	log.Fatal(app.Listen(":3069"))
}