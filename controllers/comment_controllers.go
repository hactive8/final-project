package controllers

import (
	"hactive/final-project/dto"
	"hactive/final-project/interfaces"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type CommentController struct {
	CommentService interfaces.CommentService
}

func NewCommentController(commentService interfaces.CommentService) interfaces.CommentController {
	return &CommentController{
		CommentService: commentService,
	}
}

func (h *CommentController) CreateComment(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	uid := claims["id"].(float64)

	comment := dto.CreateComment{}

	_ = c.BodyParser(&comment)

	comment.UserID = uint(uid)

	data, err := h.CommentService.CreateComment(&comment)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"code":    201,
		"message": "Comment created successfully",
		"data":    data,
	})
}
