package handler

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/majoo_test/soal_1/internal/pkg"
	"github.com/majoo_test/soal_1/pkg/entity"
)

type AuthHandler struct {
	authenticationSrv entity.AuthService
}

func NewAuthHandler(authenticationSrv entity.AuthService) *AuthHandler {
	return &AuthHandler{authenticationSrv}
}

func (handler *AuthHandler) Login(c *fiber.Ctx) error {
	var err *pkg.Errors

	body := entity.User{}

	if err := json.Unmarshal(c.Body(), &body); err != nil {
		response := fiber.Map{
			"result": nil,
			"error":  err.Error(),
		}
		c.JSON(response)
		return c.SendStatus(500)
	}

	jwtToken, err := handler.authenticationSrv.UserLogin(&body)
	if err != nil {
		response := fiber.Map{
			"result": nil,
			"error":  err.Error(),
		}
		c.JSON(response)
		return c.SendStatus(err.Status())
	}

	response := fiber.Map{
		"result": jwtToken,
		"error":  nil,
	}
	return c.JSON(response)
}
