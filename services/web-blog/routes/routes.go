package routes

import (
	"github.com/gofiber/fiber/v2"
)

func HelloMate (c *fiber.Ctx) error {
	return c.SendString("Hello Mate, 1")
}

func AddAuthor (c *fiber.Ctx) error {
	return c.SendString ("AddAuthor")
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

