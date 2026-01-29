package controllers

import (
	"context"
	"main/db"
	"main/models"
	"main/utils"
	"time"

	"github.com/gofiber/fiber/v2"
)

func AccountLogout(c *fiber.Ctx) error {
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
		return c.Status(400).JSON(fiber.Map{"msg": "error at logout"})
	}
	if pa != nil {
		if pa.Token != account.Token {
			return c.Status(400).JSON(fiber.Map{"msg": "error at logout"})
		}
		token := utils.RandomHash()
		pr, err := client.Account.FindUnique(
			db.Account.ID.Equals(pa.ID),
		).Update(
			db.Account.Token.Set(token),
			db.Account.Handshake.Set(time.Now()),
		).Exec(ctx)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"msg": "error at logout"})
		}
		if pr != nil {
			return c.Status(200).JSON(fiber.Map{"msg": "logout"})
		}
	}
	return c.Status(400).JSON(fiber.Map{"msg": "error at logout"})
}
