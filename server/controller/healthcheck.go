package controller

import (
	"github.com/SSSBoOm/CPE241_Project_Backend/domain"
	"github.com/gofiber/fiber/v2"
)

type HealthCheckController struct {
}

func NewHealthCheckController() *HealthCheckController {
	return nil
}

// HealthCheckHandler godoc
// @summary Health Check
// @description Health checking for the service
// @id HealthCheckHandler
// @produce json
// @response 200 {object} domain.Response "OK"
// @Router /api/healthcheck [get]
func (h *HealthCheckController) HealthCheck(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(domain.Response{
		SUCCESS: true,
		MESSAGE: "OK",
	})
}
