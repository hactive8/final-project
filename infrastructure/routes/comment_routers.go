package routes

import (
	"hactive/final-project/config"
	"hactive/final-project/controllers"
	"hactive/final-project/infrastructure/middlewares"
	"hactive/final-project/repository"
	"hactive/final-project/service"

	"github.com/gofiber/fiber/v2"
)

func RoutesComment(fiber *fiber.App, conf config.Config) {
	db := config.InitDatabase(conf)

	repo := repository.NewCommentRepository(db)
	svc := service.NewCommentService(repo, conf)
	ctrl := controllers.NewCommentController(svc)

	app := fiber.Group("/comments")

	app.Post("/", middlewares.JwtMiddleware(), ctrl.CreateComment)
	app.Get("/", middlewares.JwtMiddleware(), ctrl.GetComment)
}
