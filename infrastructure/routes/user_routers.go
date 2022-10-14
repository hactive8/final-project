package routes

import (
	"hactive/final-project/config"
	"hactive/final-project/controllers"
	"hactive/final-project/repository"
	"hactive/final-project/service"

	"github.com/gofiber/fiber/v2"
)

func RoutesUser(fiber *fiber.App, conf config.Config) {
	db := config.InitDatabase(conf)

	repo := repository.NewUserRepository(db)
	srv := service.NewUserService(conf, repo)
	ctrl := controllers.NewUserController(srv)

	fiber.Post("/users/register", ctrl.Register)
	fiber.Post("/users/login", ctrl.Login)
}
