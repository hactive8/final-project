package controllers

import (
	"hactive/final-project/dto"
	"hactive/final-project/interfaces"
	"hactive/final-project/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type controllers struct {
	Service interfaces.UserService
}

func NewUserController(service interfaces.UserService) interfaces.UserController {
	return &controllers{
		Service: service,
	}
}

func (h *controllers) Register(c *fiber.Ctx) error {
	user := dto.Register{}

	_ = c.BodyParser(&user)

	err := validator.New().Struct(user)
	if err != nil {
		return utils.HandleErrorValidator(err, c)
	}

	result, err := h.Service.Register(&user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"code":    400,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully",
		"code":    201,
		"data":    result,
	})
}

func (h *controllers) Login(c *fiber.Ctx) error {
	user := dto.Login{}

	_ = c.BodyParser(&user)

	err := validator.New().Struct(user)
	if err != nil {
		return utils.HandleErrorValidator(err, c)
	}

	result, err := h.Service.Login(user.Email, user.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"code":    400,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Login successfully",
		"code":    200,
		"data":    result,
	})
}