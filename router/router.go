package router

import (
	"main/controllers"

	"github.com/gofiber/fiber/v2"
)

func MainRouter(app *fiber.App) {
	mainGroup := app.Group("/")
	accountGroup := app.Group("/account")

	// Main Routes
	mainGroup.Get("/", controllers.MainRoute)

	// Account Routes
	accountGroup.Post("/login", controllers.AccountLogin)
	accountGroup.Get("/home", controllers.AccountFind)
	accountGroup.Post("/logout", controllers.AccountLogout)
	accountGroup.Post("/register", controllers.AccountRegister)

}
