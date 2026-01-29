package controllers

import (
	"context"
	"main/db"
	"main/models"
	"main/utils"
	"main/validators"
	"time"

	"github.com/gofiber/fiber/v2"
)

func AccountLogin(c *fiber.Ctx) error {
	var account models.Account
	c.BodyParser(&account)
	if account.Identifier == "" {
		return c.Status(400).JSON(fiber.Map{"msg": "identifier is invalid"})
	}
	if !validators.Password(account.Password) {
		return c.Status(400).JSON(fiber.Map{"msg": "password is too short"})
	}
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		return err
	}
	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
		}
	}()
	ctx := context.Background()
	// Login by Email
	pae, err := client.Account.FindUnique(
		db.Account.Email.Equals(account.Identifier)).Exec(ctx)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"msg": "error at sign-in"})
	}
	if pae != nil {
		if pae.Atempts >= 4 {
			return c.Status(400).JSON(fiber.Map{"msg": "account is locked"})
		}
		password := utils.EncodeString(account.Password)
		if pae.Password != password {
			validators.AccountAtempt(pae.UUID)
			return c.Status(400).JSON(fiber.Map{"msg": "email or password incorrect"})
		}
		token := utils.RandomHash()
		pr, err := client.Account.FindUnique(
			db.Account.ID.Equals(pae.ID),
		).Update(
			db.Account.Token.Set(token),
			db.Account.Handshake.Set(time.Now()),
		).Exec(ctx)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"msg": "error at login"})
		}
		if pr != nil {
			return c.Status(200).JSON(fiber.Map{"token": token})
		}
	}
	// Login by Username
	pau, err := client.Account.FindUnique(
		db.Account.Username.Equals(account.Identifier)).Exec(ctx)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"msg": "error at sign-in"})
	}
	if pau != nil {
		if pae.Atempts >= 4 {
			return c.Status(400).JSON(fiber.Map{"msg": "account is locked"})
		}
		password := utils.EncodeString(account.Password)
		if pau.Password != password {
			validators.AccountAtempt(pae.UUID)
			return c.Status(400).JSON(fiber.Map{"msg": "email or password incorrect"})
		}
		token := utils.RandomHash()
		pr, err := client.Account.FindUnique(
			db.Account.ID.Equals(pau.ID),
		).Update(
			db.Account.Token.Set(token),
			db.Account.Handshake.Set(time.Now()),
		).Exec(ctx)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"msg": "error at login"})
		}
		if pr != nil {
			return c.Status(200).JSON(fiber.Map{"token": token})
		}
	}
	return c.Status(400).JSON(fiber.Map{"msg": "email or password incorrect"})
}
