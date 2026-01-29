package controllers

import (
	"context"
	"main/db"
	"main/models"
	"main/utils"
	"main/validators"

	"github.com/gofiber/fiber/v2"
)

func AccountRegister(c *fiber.Ctx) error {
	var account models.Account
	c.BodyParser(&account)
	if !validators.Email(account.Email) {
		return c.Status(400).JSON(fiber.Map{"msg": "email is invalid"})
	}
	if !validators.Password(account.Password) {
		return c.Status(400).JSON(fiber.Map{"msg": "password is too short"})
	}
	if !validators.Names(account.FirstName) {
		return c.Status(400).JSON(fiber.Map{"msg": "firstName is too short"})
	}
	if !validators.Names(account.LastName) {
		return c.Status(400).JSON(fiber.Map{"msg": "lastName is too short"})
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
	// Find user in db by email
	pae, err := client.Account.FindUnique(
		db.Account.Email.Equals(account.Email),
	).Exec(ctx)
	if err != nil {
	}
	if pae != nil {
		return c.Status(400).JSON(fiber.Map{"msg": "error email alredy exist"})
	}
	// Find user in db by username
	pau, err := client.Account.FindUnique(
		db.Account.Username.Equals(account.Username),
	).Exec(ctx)
	if err != nil {
	}
	if pau != nil {
		return c.Status(400).JSON(fiber.Map{"msg": "error username alredy exist"})
	}
	username := utils.GetUsername(account.Email)
	// Create user in db
	initToken := utils.RandomHash()
	uuid := utils.EncodeUUID(initToken)
	pr, err := client.Account.CreateOne(
		db.Account.UUID.Set(uuid),
		db.Account.Email.Set(account.Email),
		db.Account.Token.Set(initToken),
		db.Account.Password.Set(utils.EncodeString(account.Password)),
		db.Account.FirstName.Set(account.FirstName),
		db.Account.LastName.Set(account.LastName),
		db.Account.Username.Set(username),
		db.Account.Atempts.Set(0),
	).Exec(ctx)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"msg": "error at register"})
	}
	if pr != nil {
		return c.Status(200).JSON(fiber.Map{"token": initToken})
	}
	return c.Status(400).JSON(fiber.Map{"msg": "error at register"})
}
