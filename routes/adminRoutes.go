package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jaypokale/golang-auth/controllers"
	"github.com/jaypokale/golang-auth/middleware"
)

func AdminRoutes(app *fiber.App) {
	router := app.Group("/admin/users")
	router.Use(middleware.VerifyAdmin)

	router.Get("/", controllers.GetUsers)
	router.Get("/:id", controllers.GetUserByID)
	router.Put("/:id", controllers.UpdateUserByID)
	router.Delete("/:id", controllers.DeleteUserByID)
}
