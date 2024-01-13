package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/jaypokale/golang-auth/config"
	"github.com/jaypokale/golang-auth/routes"
)

func main() {
	err := run()
	if err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	if err := config.LoadEnv(); err != nil {
		return err
	}

	if err := config.ConnectDB(); err != nil {
		return err
	}

	defer config.CloseDB()

	app := fiber.New()

	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New())

	routes.AdminRoutes(app)
	routes.AuthRoutes(app)
	routes.UserRoutes(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	app.Listen(":" + port)

	return nil
}
