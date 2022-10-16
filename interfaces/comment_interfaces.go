package interfaces

import (
	"hactive/final-project/dto"

	"github.com/gofiber/fiber/v2"
)

type CommentRepository interface {
	CreateComment(comment *dto.CreateComment) (dto.CreateComment, error)
	GetComment() ([]dto.GetAllComment, error)
	UpdateComment(id uint, comment *dto.UpdateComment) (dto.UpdateComment, error)
	GetPhotoId(id uint) (dto.GetPhotoUser, error)
}

type CommentService interface {
	CreateComment(comment *dto.CreateComment) (dto.GetComment, error)
	GetComment() ([]dto.GetAllComment, error)
	UpdateComment(id uint, comment *dto.UpdateComment) (*dto.GetPhotoUser, error)
	GetPhotoId(id uint) (dto.GetPhotoUser, error)
}

type CommentController interface {
	CreateComment(c *fiber.Ctx) error
	GetComment(c *fiber.Ctx) error
	UpdateComment(c *fiber.Ctx) error
}
