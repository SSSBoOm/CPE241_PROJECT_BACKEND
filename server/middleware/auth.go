package middleware

import (
	"github.com/SSSBoOm/CPE241_Project_Backend/domain"
	"github.com/gofiber/fiber/v2"
)

func NewAuthMiddleware(authUsecase domain.AuthUsecase) fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		return ctx.Next()
	}
}