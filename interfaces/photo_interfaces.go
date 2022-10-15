package interfaces

import (
	"hactive/final-project/dto"

	"github.com/gofiber/fiber/v2"
)

type PhotoRepository interface {
	CreatePhoto(photo *dto.CreatePhoto) (dto.CreatePhoto, error)
	GetAllPhoto() ([]dto.GetAllPhoto, error)
	UpdatePhoto(id int, photo *dto.UpdatePhoto) (dto.UpdatePhoto, error)
	DeletePhoto(id int) error
}

type PhotoService interface {
	CreatePhoto(photo *dto.CreatePhoto) (dto.GetPhoto, error)
	GetAllPhoto() ([]dto.GetAllPhoto, error)
	UpdatePhoto(id int, photo *dto.UpdatePhoto) (dto.GetUpdatePhoto, error)
	DeletePhoto(id int) error
}

type PhotoController interface {
	CreatePhoto(ctx *fiber.Ctx) error
	GetAllPhoto(ctx *fiber.Ctx) error
	UpdatePhoto(ctx *fiber.Ctx) error
	DeletePhoto(ctx *fiber.Ctx) error
}
