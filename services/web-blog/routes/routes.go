package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vvk17/microservices-go/services/web-blog/validators"
	"github.com/vvk17/microservices-go/services/web-blog/utilities"
	"github.com/vvk17/microservices-go/services/web-blog/models"
	"github.com/vvk17/microservices-go/services/web-blog/database"
	"log"
)

func HelloMate (c *fiber.Ctx) error {
	return c.SendString("Hello Mate, 1")
}

func AddAuthor (c *fiber.Ctx) error {
	response := utilities.GetBaseResponseObject()
	postBody := &validators.AuthorAddPostBody{}
	log.Println("pkg routes func AddAuthor postBody ", *postBody)

	if err := c.BodyParser(postBody); err != nil {
		response["error"] = err.Error()
		log.Println("routes-AddAuthor BodyParser error", response)
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	} else {
		if err := validators.ValidateStruct(postBody); err != nil {
			log.Println("routes-AddAuthor ValidateStruct error ", response)
			return c.Status(fiber.StatusInternalServerError).JSON(response)
		} else {
			author := models.Authors{Title: postBody.Title}
			if _, err := database.Database.Orm.Insert(&author); err != nil {
				response["error"] = err.Error()
				log.Println("routes-AddAuthor Orm.Insert error", response)
				return c.Status(fiber.StatusInternalServerError).JSON(response)
			} else {
				response["message"] = "Author added successfully"
				response["status"] = "OK"
				log.Println("routes-AddAuthor Orm.Insert OK", response)
				return c.Status(fiber.StatusCreated).JSON(response)
			}
		}
	}
}

func GetAllAuthors (c *fiber.Ctx) error {
	return c.SendString("All Authors")
}

func GetSingleAuthor (c *fiber.Ctx) error {
	return c.SendString("Single Author")
}

func DeleteAuthor (c *fiber.Ctx) error {
	return c.SendString("Delete Author")
}

func UpdateAuthor (c *fiber.Ctx) error {
	return c.SendString("Update Author")
}

