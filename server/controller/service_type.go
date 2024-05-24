package controller

import (
	"github.com/SSSBoOm/CPE241_Project_Backend/domain"
	"github.com/SSSBoOm/CPE241_Project_Backend/domain/payload"
	"github.com/SSSBoOm/CPE241_Project_Backend/internal/constant"
	"github.com/gofiber/fiber/v2"
)

type serviceTypeController struct {
	validate           domain.ValidatorUsecase
	serviceTypeService domain.ServiceTypeUsecase
}

func NewServiceTypeController(validate domain.ValidatorUsecase, serviceTypeService domain.ServiceTypeUsecase) *serviceTypeController {
	return &serviceTypeController{
		validate:           validate,
		serviceTypeService: serviceTypeService,
	}
}

// GetAll godoc
// @Summary Get all service types
// @Tags service_type
// @Accept json
// @Produce json
// @Response 200 {object} domain.Response
// @Router /api/service_type [get]
func (c *serviceTypeController) GetAll(ctx *fiber.Ctx) error {
	services, err := c.serviceTypeService.GetAll()
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
// @Summary Get service type by ID
// @Tags service_type
// @Accept json
// @Produce json
// @Param id path int true "Service Type ID"
// @Response 200 {object} domain.Response
// @Router /api/service_type/{id} [get]
func (c *serviceTypeController) GetByID(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_BAD_REQUEST,
		})
	}

	service, err := c.serviceTypeService.GetByID(id)
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
// @Summary Create service type
// @Tags service_type
// @Accept json
// @Produce json
// @Param serviceType body payload.ServiceTypeCreateDTO true "Service Type"
// @Response 200 {object} domain.Response
// @Router /api/service_type [post]
func (c *serviceTypeController) Create(ctx *fiber.Ctx) error {
	var body payload.ServiceTypeCreateDTO
	if err := c.validate.ValidateBody(ctx, &body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_BAD_REQUEST,
		})
	}

	service := make([]domain.SERVICE, 0)
	for _, item := range *body.SERVICE {
		service = append(service, domain.SERVICE{
			NAME:        item.NAME,
			DESCRIPTION: item.DESCRIPTION,
			INFORMATION: item.INFORMATION,
			PRICE:       &item.PRICE,
			IS_ACTIVE:   item.IS_ACTIVE,
		})
	}

	if _, err := c.serviceTypeService.Create(&domain.SERVICE_TYPE{
		NAME:      body.NAME,
		DETAIL:    body.DETAIL,
		IS_ACTIVE: body.IS_ACTIVE,
		SERVICE:   &service,
	}); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(domain.Response{
		SUCCESS: true,
		MESSAGE: constant.MESSAGE_SUCCESS,
	})
}

// Update godoc
// @Summary Update service type
// @Tags service_type
// @Accept json
// @Produce json
// @Param id path int true "Service Type ID"
// @Param serviceType body payload.ServiceTypeUpdateDTO true "Service Type"
// @Response 200 {object} domain.Response
// @Router /api/service_type [put]
func (c *serviceTypeController) Update(ctx *fiber.Ctx) error {
	var body payload.ServiceTypeUpdateDTO
	if err := c.validate.ValidateBody(ctx, &body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_BAD_REQUEST,
		})
	}

	if err := c.serviceTypeService.Update(
		&domain.SERVICE_TYPE{
			ID:        body.ID,
			NAME:      body.NAME,
			DETAIL:    body.DETAIL,
			IS_ACTIVE: body.IS_ACTIVE,
		},
	); err != nil {
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
// @Summary Update service type is active
// @Tags service_type
// @Accept json
// @Produce json
// @Param isActive body payload.ServiceTypeUpdateIsActiveDTO true "Service Type"
// @Response 200 {object} domain.Response
// @Router /api/service_type/active [post]
func (c *serviceTypeController) UpdateIsActive(ctx *fiber.Ctx) error {
	var body payload.ServiceTypeUpdateIsActiveDTO
	if err := c.validate.ValidateBody(ctx, &body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_BAD_REQUEST,
		})
	}

	err := c.serviceTypeService.UpdateIsActive(body.ID, body.IS_ACTIVE)
	if err != nil {
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
