package controllers

import (
	"hactive/final-project/dto"
	"hactive/final-project/interfaces"
	"strconv"
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

func (s *SosmedController) GetSosmed(c *fiber.Ctx) error {
	data, err := s.SosmedService.GetSosmed()

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    200,
		"message": "Sosmed fetched successfully",
		"data":    data,
	})
}

func (s *SosmedController) UpdateSosmed(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	ids := claims["id"].(float64)

	id := c.Params("socialMediaId")
	uid, _ := strconv.Atoi(id)

	sosmed := dto.UpdateSosmed{
		UpdatedAt: time.Now(),
	}

	_ = c.BodyParser(&sosmed)

	data, err := s.SosmedService.UpdateSosmed(uint(uid), &sosmed)

	data.UserID = uint(ids)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    200,
		"message": "Sosmed updated successfully",
		"data":    data,
	})
}
