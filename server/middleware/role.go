package middleware

import (
	"slices"

	"github.com/SSSBoOm/CPE241_Project_Backend/domain"
	"github.com/SSSBoOm/CPE241_Project_Backend/internal/constant"
	"github.com/gofiber/fiber/v2"
)

func NewRoleAuthMiddleware(rolesAuth []string) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		role := ctx.Locals(constant.CTX_ROLE).(string)

		if !slices.Contains(rolesAuth, role) {
			return ctx.Status(fiber.StatusForbidden).JSON(domain.Response{
				SUCCESS: false,
				MESSAGE: "Forbidden",
			})
		}
		return ctx.Next()
	}
}
