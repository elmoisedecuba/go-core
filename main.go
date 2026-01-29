//go:generate go run github.com/steebchen/prisma-client-go db push

package main

import (
	"main/common"
	"main/router"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	err := run()

	if err != nil {
		panic(err)
	}
}

func run() error {
	// init env
	err := common.LoadEnv()
	if err != nil {
		return err
	}
	// create app
	app := fiber.New(fiber.Config{
		AppName:      "GoBack",
		ProxyHeader:  "GoBack",
		ServerHeader: "GoBack",
	})
	// add basic middleware
	app.Use(cors.New())

	// add routes
	router.MainRouter(app)

	// start server
	var port string
	if port = os.Getenv("PORT"); port == "" {
		port = "80"
	}
	app.Listen(":" + port)

	return nil
}
