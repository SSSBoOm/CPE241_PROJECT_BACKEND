package middleware

import (
	"github.com/SSSBoOm/CPE241_Project_Backend/domain"
	"github.com/gofiber/fiber/v2"
)

func NewAuthMiddleware(sessionUsecase domain.SessionUsecase) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		ssid := ctx.Cookies("ssid")
		if ssid == "" {
			return ctx.Status(fiber.StatusUnauthorized).JSON(domain.Response{
				SUCCESS: false,
				MESSAGE: "Unauthorized",
			})
		}

		// session, err := sessionUsecase.Get(ssid)/

		// if session == nil || err != nil {
		// 	return ctx.Status(fiber.StatusUnauthorized).JSON(domain.Response{
		// 		SUCCESS: false,
		// 		MESSAGE: "Unauthorized",
		// 	})
		// }

		// ctx.Locals(constant.USER_ID, session.USER_ID)
		// ctx.Locals(constant.SESSION_TYPE, session.USER_ID)

		return ctx.Next()
	}
}
