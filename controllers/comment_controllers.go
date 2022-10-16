package controllers

import (
	"hactive/final-project/dto"
	"hactive/final-project/interfaces"
	"strconv"

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

func (h *CommentController) GetComment(c *fiber.Ctx) error {
	data, err := h.CommentService.GetComment()

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    200,
		"message": "Get all comment successfully",
		"data":    data,
	})
}

func (h *CommentController) UpdateComment(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	ids := claims["id"].(float64)

	comment := dto.UpdateComment{}
	id := c.Params("commentId")
	uid, _ := strconv.Atoi(id)

	_ = c.BodyParser(&comment)

	data, err := h.CommentService.UpdateComment(uint(uid), &comment)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if data.UserID != uint(ids) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    200,
		"message": "Comment updated successfully",
		"data":    data,
	})
}
