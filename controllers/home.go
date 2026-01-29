package controllers

import (
	"context"
	"main/db"
	"main/models"

	"github.com/gofiber/fiber/v2"
)

func AccountFind(c *fiber.Ctx) error {
	// token := utils.RandomHash()
	var account models.Account
	c.BodyParser(&account)
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		return err
	}
	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
		}
	}()
	ctx := context.Background()
	pa, err := client.Account.FindUnique(
		db.Account.Token.Equals(account.Token)).Exec(ctx)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"msg": "not found"})
	}
	return c.Status(200).JSON(pa)
}
