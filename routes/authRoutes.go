package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jaypokale/golang-auth/controllers"
)

func AuthRoutes(app *fiber.App) {
	router := app.Group("/auth")

	router.Post("/signup", controllers.CreateUser)
	router.Post("/login", controllers.LoginUser)
}
