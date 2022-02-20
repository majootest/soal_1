package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/majoo_test/soal_1/internal/pkg"
	"github.com/majoo_test/soal_1/pkg/entity"
)

type MerchantHandler struct {
	authSrv     entity.AuthService
	merchantSrv entity.MerchantService
}

func NewMerchantHandler(authSrv entity.AuthService, merchantSrv entity.MerchantService) *MerchantHandler {
	return &MerchantHandler{authSrv, merchantSrv}
}

func (handler *MerchantHandler) GetOmzetReport(c *fiber.Ctx) error {
	var err *pkg.Errors

	authToken := c.Get("Authorization")

	payload, err := handler.authSrv.UserAuth(authToken)
	if err != nil {
		response := fiber.Map{
			"result": nil,
			"error":  err.Error(),
		}
		c.JSON(response)
		return c.SendStatus(err.Status())
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

	userId, _ := payload.Get("user_id")
	result := make([]entity.MerchantOmzet, 0)
	if result, err = handler.merchantSrv.FindOmzetReportNovember(int64(userId.(float64)), pagination); err != nil {
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
