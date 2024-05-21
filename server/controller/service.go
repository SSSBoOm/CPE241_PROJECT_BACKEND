package controller

import (
	"github.com/SSSBoOm/CPE241_Project_Backend/domain"
	"github.com/SSSBoOm/CPE241_Project_Backend/domain/payload"
	"github.com/SSSBoOm/CPE241_Project_Backend/internal/constant"
	"github.com/gofiber/fiber/v2"
)

type serviceController struct {
	validate           domain.ValidatorUsecase
	serviceUsecase     domain.ServiceUsecase
	serviceTypeUsecase domain.ServiceTypeUsecase
}

func NewServiceController(validate domain.ValidatorUsecase, serviceUsecase domain.ServiceUsecase, serviceTypeUsecase domain.ServiceTypeUsecase) *serviceController {
	return &serviceController{
		validate:           validate,
		serviceUsecase:     serviceUsecase,
		serviceTypeUsecase: serviceTypeUsecase,
	}
}

// GetAll godoc
// @Summary Get all services
// @Tags service
// @Accept json
// @Produce json
// @Response 200 {object} domain.Response
// @Router /api/service [get]
func (c *serviceController) GetAll(ctx *fiber.Ctx) error {
	services, err := c.serviceUsecase.GetAll()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(domain.Response{
		SUCCESS: true,
		MESSAGE: constant.MESSAGE_SUCCESS,
		DATA:    services,
	})
}

// GetByID godoc
// @Summary Get service by ID
// @Tags service
// @Accept json
// @Produce json
// @Param id path int true "Service ID"
// @Response 200 {object} domain.Response
// @Router /api/service/{id} [get]
func (c *serviceController) GetByID(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_BAD_REQUEST,
		})
	}

	service, err := c.serviceUsecase.GetById(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(domain.Response{
		SUCCESS: true,
		MESSAGE: constant.MESSAGE_SUCCESS,
		DATA:    service,
	})
}

// Create godoc
// @Summary Create service
// @Tags service
// @Accept json
// @Produce json
// @Param service body payload.ServiceCreateDTO true "Service"
// @Response 200 {object} domain.Response
// @Router /api/service [post]
func (c *serviceController) Create(ctx *fiber.Ctx) error {
	var body payload.ServiceCreateDTO
	if err := c.validate.ValidateBody(ctx, &body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_BAD_REQUEST,
		})
	}

	if err := c.serviceUsecase.Create(&domain.SERVICE{
		NAME:            body.NAME,
		PRICE:           &body.PRICE,
		IS_ACTIVE:       body.IS_ACTIVE,
		SERVICE_TYPE_ID: body.SERVICE_TYPE_ID,
	}); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(domain.Response{
		SUCCESS: true,
		MESSAGE: constant.MESSAGE_SUCCESS,
	})
}

// Update godoc
// @Summary Update service
// @Tags service
// @Accept json
// @Produce json
// @Param service body payload.ServiceUpdateDTO true "Service"
// @Response 200 {object} domain.Response
// @Router /api/service [put]
func (c *serviceController) Update(ctx *fiber.Ctx) error {
	var body payload.ServiceUpdateDTO
	if err := c.validate.ValidateBody(ctx, &body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_BAD_REQUEST,
		})
	}

	if err := c.serviceUsecase.Update(&domain.SERVICE{
		ID:              body.ID,
		NAME:            body.NAME,
		PRICE:           &body.PRICE,
		IS_ACTIVE:       body.IS_ACTIVE,
		SERVICE_TYPE_ID: body.SERVICE_TYPE_ID,
	}); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(domain.Response{
		SUCCESS: true,
		MESSAGE: constant.MESSAGE_SUCCESS,
	})
}

// UpdateIsActive godoc
// @Summary Update service is active
// @Tags service
// @Accept json
// @Produce json
// @Param service body payload.ServiceUpdateIsActiveDTO true "Service"
// @Response 200 {object} domain.Response
// @Router /api/service/active [put]
func (c *serviceController) UpdateIsActive(ctx *fiber.Ctx) error {
	var body payload.ServiceUpdateIsActiveDTO
	if err := c.validate.ValidateBody(ctx, &body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_BAD_REQUEST,
		})
	}

	if err := c.serviceUsecase.UpdateIsActive(body.ID, body.IS_ACTIVE); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(domain.Response{
		SUCCESS: true,
		MESSAGE: constant.MESSAGE_SUCCESS,
	})
}
