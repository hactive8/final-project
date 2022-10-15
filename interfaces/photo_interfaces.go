package interfaces

import (
	"hactive/final-project/dto"

	"github.com/gofiber/fiber/v2"
)

type PhotoRepository interface {
	CreatePhoto(photo *dto.CreatePhoto) (dto.CreatePhoto, error)
}

type PhotoService interface {
	CreatePhoto(photo *dto.CreatePhoto) (dto.GetPhoto, error)
}

type PhotoController interface {
	CreatePhoto(ctx *fiber.Ctx) error
}
