package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/majoo_test/soal_1/internal/pkg"
	"github.com/majoo_test/soal_1/pkg/entity"
)

type UserHandler struct {
	authSrv entity.AuthService
	userSrv entity.UserService
}

func NewUserHandler(authSrv entity.AuthService, userSrv entity.UserService) *UserHandler {
	return &UserHandler{authSrv, userSrv}
}

func (handler *UserHandler) Find(c *fiber.Ctx) error {
	var err *pkg.Errors

	authToken := c.Get("Authorization")
	if _, err := handler.authSrv.UserAuth(authToken); err != nil {
		response := fiber.Map{
			"result": nil,
			"error":  err.Error(),
		}
		c.JSON(response)
		return c.SendStatus(err.Status())
	}

	result := make([]entity.User, 0)

	query := make(map[string]interface{})
	id := c.Query("id")
	name := c.Query("name")

	if id != "" {
		query["id"] = id
	}
	if name != "" {
		query["name"] = name
	}

	limitPerPage, e := strconv.Atoi(c.Query("limit_per_page"))
	if e != nil {
		response := fiber.Map{
			"result": nil,
			"error":  "Invalid Parameter",
		}
		c.JSON(response)
		return c.SendStatus(400)
	}
	pageNo, e := strconv.Atoi(c.Query("page_no"))
	if e != nil {
		response := fiber.Map{
			"result": nil,
			"error":  "Invalid Parameter",
		}
		c.JSON(response)
		return c.SendStatus(400)
	}

	pagination := &entity.Pagination{
		Limit:  limitPerPage,
		PageNo: pageNo,
	}
	if result, err = handler.userSrv.FindUsers(query, pagination); err != nil {
		response := fiber.Map{
			"result": nil,
			"error":  err.Error(),
		}
		c.JSON(response)
		return c.SendStatus(err.Status())
	}

	response := fiber.Map{
		"result": result,
		"error":  nil,
	}
	return c.JSON(response)
}
