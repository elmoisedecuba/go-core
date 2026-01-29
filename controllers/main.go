package controllers

import "github.com/gofiber/fiber/v2"

func MainRoute(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{"msg": "Hello Hacker!"})
}
