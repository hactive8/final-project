package interfaces

import (
	"hactive/final-project/dto"

	"github.com/gofiber/fiber/v2"
)

type SosmedRepository interface {
	CreateSosmed(sosmed *dto.CreateSosmed) (dto.CreateSosmed, error)
	GetSosmed() ([]dto.GetAllSosmed, error)
	GetSosmedByID(id uint) (dto.GetAllSosmed, error)
	UpdateSosmed(id uint, sosmed *dto.UpdateSosmed) (dto.UpdateSosmed, error)
	DeleteSosmed(id uint) error
}

type SosmedService interface {
	CreateSosmed(sosmed *dto.CreateSosmed) (dto.CreateSosmed, error)
	GetSosmed() ([]dto.GetAllSosmed, error)
	GetSosmedByID(id uint) (dto.GetAllSosmed, error)
	UpdateSosmed(id uint, sosmed *dto.UpdateSosmed) (dto.UpdateSosmed, error)
	DeleteSosmed(id uint) error
}

type SosmedController interface {
	CreateSosmed(c *fiber.Ctx) error
	GetSosmed(c *fiber.Ctx) error
	UpdateSosmed(c *fiber.Ctx) error
	DeleteSosmed(c *fiber.Ctx) error
}
