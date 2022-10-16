package controllers

import (
	"hactive/final-project/dto"
	"hactive/final-project/interfaces"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type SosmedController struct {
	SosmedService interfaces.SosmedService
}

func NewSosmedController(sosmedService interfaces.SosmedService) interfaces.SosmedController {
	return &SosmedController{
		SosmedService: sosmedService,
	}
}

func (s *SosmedController) CreateSosmed(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	uid := claims["id"].(float64)

	sosmed := dto.CreateSosmed{
		CreatedAt: time.Now(),
	}

	_ = c.BodyParser(&sosmed)

	sosmed.UserID = uint(uid)

	data, err := s.SosmedService.CreateSosmed(&sosmed)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"code":    201,
		"message": "Sosmed created successfully",
		"data":    data,
	})
}
