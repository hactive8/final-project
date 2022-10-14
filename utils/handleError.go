package utils

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func HandleErrorValidator(err error, c *fiber.Ctx) error {
	report, ok := err.(*fiber.Error)
	if !ok {
		report = fiber.NewError(http.StatusBadRequest, err.Error())
	}

	if castedObject, ok := err.(validator.ValidationErrors); ok {
		for _, err := range castedObject {
			switch err.Tag() {
			case "required":
				report.Message = fmt.Sprintf("%s is required", err.Field())
			case "email":
				report.Message = fmt.Sprintf("%s is not valid", err.Field())
			case "min":
				report.Message = fmt.Sprintf("%s min 8 character", err.Field())
			}
		}
	}

	return c.JSON(fiber.Map{
		"message": report.Message,
		"code":    report.Code,
	})
}
