package controllers

import (
	"hactive/final-project/dto"
	"hactive/final-project/interfaces"
	"hactive/final-project/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type PhotoController struct {
	PhotoService interfaces.PhotoService
}

func NewPhotoController(service interfaces.PhotoService) interfaces.PhotoController {
	return &PhotoController{
		PhotoService: service,
	}
}

func (h *PhotoController) CreatePhoto(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["id"].(float64)

	photo := dto.CreatePhoto{
		UserId: uint(id),
	}

	_ = c.BodyParser(&photo)

	err := validator.New().Struct(photo)
	if err != nil {
		return utils.HandleErrorValidator(err, c)
	}

	result, err := h.PhotoService.CreatePhoto(&photo)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"code":    400,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Photo created successfully",
		"code":    201,
		"data":    result,
	})
}
