package controllers

import (
	"hactive/final-project/dto"
	"hactive/final-project/interfaces"
	"hactive/final-project/utils"
	"strconv"

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

func (h *PhotoController) GetAllPhoto(c *fiber.Ctx) error {
	result, err := h.PhotoService.GetAllPhoto()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"code":    400,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Get all photo successfully",
		"code":    200,
		"data":    result,
	})
}

func (h *PhotoController) UpdatePhoto(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	uid := claims["id"].(float64)

	id, _ := strconv.Atoi(c.Params("photoId"))

	photo := dto.UpdatePhoto{}

	_ = c.BodyParser(&photo)

	err := validator.New().Struct(photo)
	if err != nil {
		return utils.HandleErrorValidator(err, c)
	}

	result, err := h.PhotoService.UpdatePhoto(id, &photo)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"code":    400,
		})
	}

	result.UserID = uint(uid)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Photo updated successfully",
		"code":    200,
		"data":    result,
	})
}

func (h *PhotoController) DeletePhoto(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("photoId"))

	err := h.PhotoService.DeletePhoto(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"code":    400,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Photo deleted successfully",
		"code":    200,
	})
}
