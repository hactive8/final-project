package interfaces

import (
	"hactive/final-project/dto"

	"github.com/gofiber/fiber/v2"
)

type UserRepository interface {
	Register(user *dto.Register) (dto.Register, error)
	Login(email, password string) (dto.Login, error)
}

type UserService interface {
	Register(user *dto.Register) (dto.Register, error)
	Login(email, password string) (string, error)
}

type UserController interface {
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
}
