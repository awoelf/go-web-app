package controllers

import (
	"log"

	"github.com/awoelf/go-web-app/services"
	"github.com/gofiber/fiber/v2"
)

var comment services.Comment

// GET/api/comments
func GetAllComments(c *fiber.Ctx) error {
	res, err := comment.GetAllComments()
	if err != nil {
		// Need to write a message logger
		log.Print(err)
		return err
	}
	c.Status(fiber.StatusOK).JSON(res)
	
	return nil
}

// GET/api/comment/{id}
func GetComment(c *fiber.Ctx) error {
	id := c.Params("id")

	res, err := comment.GetComment(id)
	if err != nil {
		log.Print(err)
		return err
	}
	c.Status(fiber.StatusOK).JSON(res)

	return nil
}

// POST/api/comment
func CreateComment(c *fiber.Ctx) error {
	var body services.Comment
	
	if err := c.BodyParser(&body); err != nil {
		log.Print(err)
		return err
	}

	log.Print(body.CommentText)

	res, err := comment.CreateComment(body)
	if err != nil {
		log.Print(err)
		return err
	}

	c.Status(fiber.StatusOK).JSON(res)

	return nil
}

// PUT/api/comment/{id}
func UpdateComment(c *fiber.Ctx) error {
	id := c.Params("id")

	var body services.Comment
	
	if err := c.BodyParser(&body); err != nil {
		log.Print(err)
		return err
	}

	res, err := comment.UpdateComment(id, body)
	if err != nil {
		log.Print(err)
		return err
	}
	c.Status(fiber.StatusOK).JSON(res)

	return nil
}

// DELETE/api/comment/{id}
func DeleteComment(c *fiber.Ctx) error {
	id := c.Params("id")
	err := comment.DeleteComment(id)
	if err != nil {
		log.Print(err)
		return err
	}

	c.Status(fiber.StatusOK).SendString("Message deleted successfully.")
	
	return nil
}
