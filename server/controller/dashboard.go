package controller

import (
	"github.com/SSSBoOm/CPE241_Project_Backend/domain"
	"github.com/SSSBoOm/CPE241_Project_Backend/domain/payload"
	"github.com/SSSBoOm/CPE241_Project_Backend/internal/constant"
	"github.com/gofiber/fiber/v2"
)

type DashboardController struct {
	validate         domain.ValidatorUsecase
	dashboardUsecase domain.DashboardUsecase
}

func NewDashboardController(validate domain.ValidatorUsecase, dashboardUsecase domain.DashboardUsecase) *DashboardController {
	return &DashboardController{
		validate:         validate,
		dashboardUsecase: dashboardUsecase,
	}
}

func (c *DashboardController) GetDashboardRoomTypeReservation(ctx *fiber.Ctx) error {
	var body payload.GetDashboardReservationDTO
	if err := c.validate.ValidateBody(ctx, &body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_BAD_REQUEST,
		})
	}

	dashboard, err := c.dashboardUsecase.GetDashboardRoomTypeReservation(body.START_DATE, body.END_DATE)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(domain.Response{
		SUCCESS: true,
		MESSAGE: constant.MESSAGE_SUCCESS,
		DATA:    dashboard,
	})
}

func (c *DashboardController) GetDashboardServiceTypeReservation(ctx *fiber.Ctx) error {
	var body payload.GetDashboardReservationDTO
	if err := c.validate.ValidateBody(ctx, &body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_BAD_REQUEST,
		})
	}

	dashboard, err := c.dashboardUsecase.GetDashboardServiceTypeReservation(body.START_DATE, body.END_DATE)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(domain.Response{
		SUCCESS: true,
		MESSAGE: constant.MESSAGE_SUCCESS,
		DATA:    dashboard,
	})
}

func (c *DashboardController) GetDashboardReservationByPaymentType(ctx *fiber.Ctx) error {
	var body payload.GetDashboardReservationDTO
	if err := c.validate.ValidateBody(ctx, &body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_BAD_REQUEST,
		})
	}

	dashboard, err := c.dashboardUsecase.GetDashboardReservationByPaymentType(body.START_DATE, body.END_DATE)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(domain.Response{
		SUCCESS: true,
		MESSAGE: constant.MESSAGE_SUCCESS,
		DATA:    dashboard,
	})
}
