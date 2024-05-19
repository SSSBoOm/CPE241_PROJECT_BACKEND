package controller

import (
	"strconv"

	"github.com/SSSBoOm/CPE241_Project_Backend/domain"
	"github.com/SSSBoOm/CPE241_Project_Backend/domain/payload"
	"github.com/SSSBoOm/CPE241_Project_Backend/internal/constant"
	"github.com/gofiber/fiber/v2"
)

type RoomTypeController struct {
	validator       domain.ValidatorUsecase
	roomTypeUsecase domain.RoomTypeUsecase
}

func NewRoomTypeController(validator domain.ValidatorUsecase, roomTypeUsecase domain.RoomTypeUsecase) *RoomTypeController {
	return &RoomTypeController{
		validator:       validator,
		roomTypeUsecase: roomTypeUsecase,
	}
}

// GetAll godoc
// @Summary								Get all room types
// @Description						Get all room types
// @Tags									room_type
// @Accept								json
// @produce								json
// @Response 200 {object} domain.Response
// @Router /api/room_type/all	[get]
func (c *RoomTypeController) GetRoomTypeList(ctx *fiber.Ctx) error {
	roomTypeList, err := c.roomTypeUsecase.GetAll()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(domain.Response{
		SUCCESS: true,
		MESSAGE: constant.MESSAGE_SUCCESS,
		DATA:    roomTypeList,
	})
}

// GetByID godoc
// @Summary								Get room type by id
// @Description						Get room type by id
// @Tags									room_type
// @Accept								json
// @produce								json
// @Param									id path int true "Room Type ID"
// @Router /api/room_type/{id}	[get]
func (c *RoomTypeController) GetRoomTypeByID(ctx *fiber.Ctx) error {
	roomTypeID, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
		})
	}

	roomType, err := c.roomTypeUsecase.GetByID(roomTypeID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(domain.Response{
		SUCCESS: true,
		MESSAGE: constant.MESSAGE_SUCCESS,
		DATA:    roomType,
	})
}

// CreateRoomType godoc
// @Summary								Create room type
// @Description						Create room type
// @Tags									room_type
// @Accept								json
// @produce								json
// @Param									roomType body domain.RoomType true "Room Type"
// @Router /api/room_type/	[post]
func (c *RoomTypeController) CreateRoomType(ctx *fiber.Ctx) error {
	var body payload.CreateRoomType
	if err := ctx.BodyParser(body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_BAD_REQUEST,
		})
	}

	err := c.roomTypeUsecase.Create(&domain.RoomType{
		NAME:      body.NAME,
		DETAIL:    body.DETAIL,
		IS_ACTIVE: body.IS_ACTIVE,
	})
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

// UpdateRoomType godoc
// @Summary								Update room type
// @Description						Update room type
// @Tags									room_type
// @Accept								json
// @produce								json
// @Param									payload body	payload.UpdateRoomType true "Payload"
// @Router /api/room_type/	[put]
func (c *RoomTypeController) UpdateRoomType(ctx *fiber.Ctx) error {
	var body payload.UpdateRoomType
	if err := c.validator.ValidateBody(ctx, &body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_BAD_REQUEST,
		})
	}

	roomType := &domain.RoomType{
		ID:        body.ID,
		NAME:      body.NAME,
		DETAIL:    body.DETAIL,
		IS_ACTIVE: body.IS_ACTIVE,
	}

	err := c.roomTypeUsecase.Update(roomType)
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

// UpdateIsActive godoc
// @Summary								Update room type is active
// @Description						Update room type is active
// @Tags									room_type
// @Accept								json
// @produce								json
// @Param									payload body	payload.UpdateRoomTypeIsActiveDTO true "Payload"
// @Router /api/room_type/active	[post]
func (c *RoomTypeController) UpdateRoomTypeIsActive(ctx *fiber.Ctx) error {
	var body payload.UpdateRoomTypeIsActiveDTO
	if err := c.validator.ValidateBody(ctx, &body); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
		})
	}

	if err := c.roomTypeUsecase.UpdateIsActive(body.ID, body.IsActive); err != nil {
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
