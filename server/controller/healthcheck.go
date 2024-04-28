package controller

import "github.com/gofiber/fiber/v2"

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
func (h *HealthCheckController) HealthCheck(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  fiber.StatusOK,
		"message": "OK"})
}
