package server

import (
	"hactive/final-project/config"
	"hactive/final-project/infrastructure/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func InitServer() *fiber.App {
	app := fiber.New()
	conf := config.Config{}

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	routes.RoutesUser(app, conf)
	routes.RoutesPhoto(app, conf)

	return app
}
