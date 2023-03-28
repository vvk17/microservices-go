package main

import (
	
	"github.com/joho/godotenv"
	"log"
	"github.com/beego/beego/v2/client/orm"
	"github.com/gofiber/fiber/v2"
	"github.com/vvk17/microservices-go/services/web-blog/database"
	"github.com/vvk17/microservices-go/services/web-blog/models"
	"github.com/vvk17/microservices-go/services/web-blog/routes"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("pkg main: Environment load error")
		log.Println(err)
	}
	app := fiber.New()
	
	orm.RegisterModel(new(models.Authors))
	log.Println("pkg main: orm.RegisterModel - success")

	database.ConnectDB()
	log.Print("pkg main: ConnectDB - success")

	SetupRoutes(app)
	log.Print("pkg main: SetupRoutes - success")

	log.Fatal(app.Listen(":3069"))
}

func SetupRoutes (app *fiber.App) {
	app.Get("/hello", routes.HelloMate)
	app.Get("/authors", routes.GetAllAuthors)
	app.Get("/author/:id", routes.GetSingleAuthor)
	app.Delete("/author", routes.DeleteAuthor)
	app.Put("/author", routes.UpdateAuthor)
}