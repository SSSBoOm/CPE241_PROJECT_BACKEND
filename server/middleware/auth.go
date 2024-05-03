package middleware

import (
	"time"

	"github.com/SSSBoOm/CPE241_Project_Backend/domain"
	"github.com/SSSBoOm/CPE241_Project_Backend/internal/constant"
	"github.com/gofiber/fiber/v2"
)

func NewAuthMiddleware(sessionUsecase domain.SessionUsecase, userUsecase domain.UserUsecase, roleUsecase domain.RoleUsecase) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		ssid := ctx.Cookies(constant.SessionCookieName)
		if ssid == "" {
			return ctx.Status(fiber.StatusUnauthorized).JSON(domain.Response{
				SUCCESS: false,
				MESSAGE: "Unauthorized",
			})
		}

		session, err := sessionUsecase.Get(ssid)
		if session == nil || err != nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON(domain.Response{
				SUCCESS: false,
				MESSAGE: "Unauthorized",
			})
		} else if !time.Now().Before(session.EXPIRED_AT) {
			ctx.Cookie(&fiber.Cookie{Name: constant.SessionCookieName, Expires: time.Unix(0, 0)})
			return ctx.Status(fiber.StatusUnauthorized).JSON(domain.Response{
				SUCCESS: false,
				MESSAGE: "Unauthorized",
			})
		}

		user, err := userUsecase.FindById(session.USER_ID)
		if user == nil || err != nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON(domain.Response{
				SUCCESS: false,
				MESSAGE: "Unauthorized",
			})
		}

		role, err := roleUsecase.Get(user.ROLE_ID)
		if role == nil || err != nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON(domain.Response{
				SUCCESS: false,
				MESSAGE: "Unauthorized",
			})
		}

		ctx.Locals(constant.CTX_USER_ID, session.USER_ID)
		ctx.Locals(constant.CTX_ROLE, role.NAME)

		return ctx.Next()
	}
}
