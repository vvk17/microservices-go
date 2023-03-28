package main

import (
	"log"
	"github.com/beego/beego/v2/client/orm"
	"github.com/gofiber/fiber/v2"
	"github.com/vvk17/microservices-go/services/web-blog/database"
	"github.com/vvk17/microservices-go/services/web-blog/models"
)

func main() {
	app := fiber.New()
	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello World, 1")
	})
	
	orm.RegisterModel(new(models.Authors))
	log.Println("orm.RegisterModel - success")

	database.ConnectDB()
	log.Print("ConnectDB - success")

	log.Fatal(app.Listen(":3069"))
}