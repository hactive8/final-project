package interfaces

import (
	"hactive/final-project/dto"

	"github.com/gofiber/fiber/v2"
)

type CommentRepository interface {
	CreateComment(comment *dto.CreateComment) (dto.CreateComment, error)
}

type CommentService interface {
	CreateComment(comment *dto.CreateComment) (dto.GetComment, error)
}

type CommentController interface {
	CreateComment(c *fiber.Ctx) error
}
