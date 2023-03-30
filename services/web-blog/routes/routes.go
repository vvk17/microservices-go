package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vvk17/microservices-go/services/web-blog/validators"
	"github.com/vvk17/microservices-go/services/web-blog/utilities"
	"github.com/vvk17/microservices-go/services/web-blog/models"
	"github.com/vvk17/microservices-go/services/web-blog/database"
	"log"
	"github.com/beego/beego/v2/client/orm"
	"strconv"
	"errors" 
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
			log.Println("routes-AddAuthor ValidateStruct error ", err)
			return c.Status(fiber.StatusInternalServerError).JSON(err)
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
	response := utilities.GetBaseResponseObject()

	qs := database.Database.Orm.QueryTable(models.Authors{}).OrderBy("-created_at")
	var maps []orm.Params

	if count, err := qs.Values(&maps, "id", "title"); err != nil {
		response["error"] = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	} else {
		response["count"] = count
		response["data"] = maps
		response["status"] = "OK"
		delete (response, "message")
		return c.Status(fiber.StatusOK).JSON(response)
	}

}

func GetSingleAuthor (c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	response := utilities.GetBaseResponseObject()

	author := &models.Authors{Id: id}

	if err := database.Database.Orm.Read(author); err != nil {
		response["error"] = err.Error()
		if errors.Is(err, orm.ErrNoRows) {
			response["message"] = "Author not found"
			return c.Status(fiber.StatusNotFound).JSON(response)
		}
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	} else {
		response["data"] = author
		response["status"] = "OK"
		delete(response, "message")
		return c.Status(fiber.StatusOK).JSON(response)
	}
}

func DeleteAuthor (c *fiber.Ctx) error {
	response := utilities.GetBaseResponseObject()
	postBody := &validators.AuthorDeletePostBody{}

	if err := utilities.PostBodyValidation(c, postBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	} else {
		author := &models.Authors{Id: postBody.Id}
		if num, err := database.Database.Orm.Delete(author); err != nil {
			response["error"] = err.Error()
			return c.Status(fiber.StatusInternalServerError).JSON(response)
		} else {
			if num == 0 {
				response["message"] = "No record found"
				return c.Status(fiber.StatusNotFound).JSON(response)
			} else {
				response["mesasge"] = "Deleted successfully"
				response["status"] = "OK"
				return c.Status(fiber.StatusOK).JSON(response)
			}
		}
	}
}

func UpdateAuthor (c *fiber.Ctx) error {
	return c.SendString("Update Author")
}

