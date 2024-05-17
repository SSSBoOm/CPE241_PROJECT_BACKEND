package controller

import (
	"strconv"

	"github.com/SSSBoOm/CPE241_Project_Backend/domain"
	"github.com/SSSBoOm/CPE241_Project_Backend/domain/payload"
	"github.com/SSSBoOm/CPE241_Project_Backend/internal/constant"
	"github.com/SSSBoOm/CPE241_Project_Backend/internal/validator"
	"github.com/gofiber/fiber/v2"
)

type maintenanceController struct {
	validator          domain.ValidatorUsecase
	maintenanceUsecase domain.MaintenanceUsecase
}

func NewMaintenanceController(validator domain.ValidatorUsecase, maintenanceUsecase domain.MaintenanceUsecase) *maintenanceController {
	return &maintenanceController{
		validator:          validator,
		maintenanceUsecase: maintenanceUsecase,
	}
}

// GetAll Maintenance	godoc
// @Summary								GetAll Maintenance
// @Description						GetAll Maintenance
// @Tags									Maintenance
// @Accept								json
// @produce								json
// @Security							ApiKeyAuth
// @Router /api/maintenance/all [get]
func (c *maintenanceController) GetAll(ctx *fiber.Ctx) error {
	maintenances, err := c.maintenanceUsecase.GetAll()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(domain.Response{
		SUCCESS: true,
		MESSAGE: constant.MESSAGE_SUCCESS,
		DATA:    maintenances,
	})
}

// GetByID Maintenance	godoc
// @Summary								GetByID Maintenance
// @Description						GetByID Maintenance
// @Tags									Maintenance
// @Accept								json
// @produce								json
// @Security							ApiKeyAuth
// @Param									id path string true "ID"
// @Response 200 {object} domain.Response
// @Router /api/maintenance/{id} [get]
func (c *maintenanceController) GetByID(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INVALID_BODY,
		})
	}
	maintenance, err := c.maintenanceUsecase.GetByID(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(domain.Response{
		SUCCESS: true,
		MESSAGE: constant.MESSAGE_SUCCESS,
		DATA:    maintenance,
	})
}

// Create Maintenance	godoc
// @Summary								Create Maintenance
// @Description						Create Maintenance
// @Tags									Maintenance
// @Accept								json
// @produce								json
// @Security							ApiKeyAuth
// @Param									ssid header string true "Session ID"
// @Param									body body payload.MAINTENANCE_CREATE true "MAINTENANCE_CREATE"
// @Success 201 {object} domain.Response
// @Router /api/maintenance [post]
func (c *maintenanceController) Create(ctx *fiber.Ctx) error {
	userId := ctx.Locals(constant.CTX_USER_ID).(string)
	var body payload.MaintenanceCreateDTO
	if err := validator.NewPayloadValidator().ValidateBody(ctx, &body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INVALID_BODY,
		})
	}

	var MAINTENANCE_LOG = make([]domain.MAINTENANCE_LOG, 0)
	for _, item := range *body.MAINTENANCE_LOG {
		MAINTENANCE_LOG = append(MAINTENANCE_LOG, domain.MAINTENANCE_LOG{
			DESCRIPTION: item.DESCRIPTION,
			STATUS:      item.STATUS,
		})
	}

	data := &domain.MAINTENANCE{
		ROOM_ID:         body.ROOM_ID,
		STAFF_ID:        userId,
		MAINTENANCE_LOG: &MAINTENANCE_LOG,
	}

	err := c.maintenanceUsecase.CreateWithMaintenance_Log(data)
	if err != nil {
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
