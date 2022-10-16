package routes

import (
	"hactive/final-project/config"
	"hactive/final-project/controllers"
	"hactive/final-project/infrastructure/middlewares"
	"hactive/final-project/repository"
	"hactive/final-project/service"

	"github.com/gofiber/fiber/v2"
)

func RoutesSosmed(fiber *fiber.App, conf config.Config) {
	db := config.InitDatabase(conf)

	repo := repository.NewSosmedRepository(db)
	svc := service.NewSosmedService(repo, conf)
	ctrl := controllers.NewSosmedController(svc)

	app := fiber.Group("/socialmedias")

	app.Post("/", middlewares.JwtMiddleware(), ctrl.CreateSosmed)
	app.Get("/", middlewares.JwtMiddleware(), ctrl.GetSosmed)
}
