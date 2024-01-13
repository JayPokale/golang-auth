package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jaypokale/golang-auth/controllers"
	"github.com/jaypokale/golang-auth/middleware"
)

func UserRoutes(app *fiber.App) {
	router := app.Group("/users")
	router.Use(middleware.VerifyUser)

	router.Get("/", controllers.GetUserByID)
	router.Put("/", controllers.UpdateUserByID)
	router.Delete("/", controllers.DeleteUserByID)
}
