package controller

import (
	"fmt"
	"time"

	"github.com/SSSBoOm/CPE241_Project_Backend/domain"
	"github.com/SSSBoOm/CPE241_Project_Backend/internal/config"
	"github.com/SSSBoOm/CPE241_Project_Backend/internal/constant"
	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	config         *config.Config
	authUsecase    domain.AuthUsecase
	googleUsecase  domain.GoogleUsecase
	userUsecase    domain.UserUsecase
	sessionUsecase domain.SessionUsecase
}

func NewAuthController(config *config.Config, authUsecase domain.AuthUsecase, googleUsecase domain.GoogleUsecase, userUsecase domain.UserUsecase, sessionUsecase domain.SessionUsecase) *AuthController {
	return &AuthController{
		config:         config,
		authUsecase:    authUsecase,
		googleUsecase:  googleUsecase,
		userUsecase:    userUsecase,
		sessionUsecase: sessionUsecase,
	}
}

// GetUrl godoc
// @Summary Get Url To Google Login
// @Description Get Url To Google Login
// @Tags Auth
// @Produce json
// @Response 200 {object} domain.Response "OK"
// @Router /api/auth/google [get]
func (auth *AuthController) GetUrl(c *fiber.Ctx) error {
	path := auth.googleUsecase.GoogleConfig()
	url := path.AuthCodeURL("state")

	return c.Status(200).JSON(domain.Response{
		SUCCESS: true,
		MESSAGE: constant.MESSAGE_SUCCESS,
		DATA: map[string]string{
			"url": url,
		},
	})
}

// SignIn godoc
// @Summary Sign In
// @Description Sign In
// @Tags Auth
// @Produce json
// @Param code query string true "Code"
// @Response 200 {object} domain.Response "OK"
// @Router /api/auth/google/callback [post]
func (auth *AuthController) SignInWithGoogle(ctx *fiber.Ctx) error {
	cookie, err := auth.authUsecase.SignInWithGoogle(ctx)
	if err != nil {
		fmt.Println("SignIn :", err)
		return ctx.Status(fiber.StatusInternalServerError).Redirect(auth.config.FRONTEND_URL)
	}

	ctx.Cookie(cookie)
	return ctx.Redirect(auth.config.FRONTEND_URL)
}

// Logout godoc
// @Summary Logout
// @Description Health checking for the service
// @Tags Auth
// @Produce json
// @Response 200 {object} domain.Response "OK"
// @Router /api/auth/logout [get]
func (auth *AuthController) SignOut(ctx *fiber.Ctx) error {
	ssid := ctx.Cookies(constant.SessionCookieName)
	if err := auth.sessionUsecase.Delete(ssid); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
		})
	}

	ctx.Cookie(&fiber.Cookie{Name: constant.SessionCookieName, Expires: time.Unix(0, 0)})
	return ctx.Status(fiber.StatusOK).JSON(domain.Response{
		SUCCESS: true,
		MESSAGE: constant.MESSAGE_SUCCESS,
	})
}
