package interfaces

import (
	"hactive/final-project/dto"

	"github.com/gofiber/fiber/v2"
)

type SosmedRepository interface {
	CreateSosmed(sosmed *dto.CreateSosmed) (dto.CreateSosmed, error)
	GetSosmed() ([]dto.GetAllSosmed, error)
}

type SosmedService interface {
	CreateSosmed(sosmed *dto.CreateSosmed) (dto.CreateSosmed, error)
	GetSosmed() ([]dto.GetAllSosmed, error)
}

type SosmedController interface {
	CreateSosmed(c *fiber.Ctx) error
	GetSosmed(c *fiber.Ctx) error
}
