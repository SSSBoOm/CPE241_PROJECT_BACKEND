package controller

import (
	"fmt"
	"strconv"

	"github.com/SSSBoOm/CPE241_Project_Backend/domain"
	"github.com/SSSBoOm/CPE241_Project_Backend/domain/payload"
	"github.com/SSSBoOm/CPE241_Project_Backend/internal/constant"
	"github.com/SSSBoOm/CPE241_Project_Backend/internal/validator"
	"github.com/gofiber/fiber/v2"
)

type roomController struct {
	validator   domain.ValidatorUsecase
	roomUsecase domain.RoomUsecase
}

func NewRoomController(validator domain.ValidatorUsecase, roomUsecase domain.RoomUsecase) *roomController {
	return &roomController{
		validator:   validator,
		roomUsecase: roomUsecase,
	}
}

// GetAll godoc
// @Summary								Get all rooms
// @Description						Get all rooms
// @Tags									room
// @Accept								json
// @produce								json
// @Router /api/room/all	[get]
func (c *roomController) GetAll(ctx *fiber.Ctx) error {
	rooms, err := c.roomUsecase.GetAll()
	if err != nil {
		fmt.Println(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(domain.Response{
		SUCCESS: true,
		MESSAGE: "OK",
		DATA:    rooms,
	})
}

// GetByID godoc
// @Summary								Get room by id
// @Description						Get room by id
// @Tags									room
// @Accept								json
// @produce								json
// @Param									id path int true "Room ID"
// @Router /api/room/{id}	[get]
func (c *roomController) GetByID(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
		})
	}
	room, err := c.roomUsecase.GetByID(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
		})
	} else if room == nil {
		return ctx.Status(fiber.StatusNotFound).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_NOT_FOUND,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(domain.Response{
		SUCCESS: true,
		MESSAGE: "OK",
		DATA:    room,
	})
}

// Create godoc
// @Summary								Create room
// @Description						Create room
// @Tags									room
// @Accept								json
// @produce								json
// @Param									room body domain.ROOM true "Room"
// @Router /api/room	[post]
func (c *roomController) Create(ctx *fiber.Ctx) error {
	var body payload.RoomCreateDTO
	if err := validator.NewPayloadValidator().ValidateBody(ctx, &body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INVALID_BODY,
		})
	}

	room := &domain.ROOM{
		ROOM_NUMBER:  body.ROOM_NUMBER,
		IS_ACTIVE:    body.IS_ACTIVE,
		ROOM_TYPE_ID: body.ROOM_TYPE_ID,
	}

	if err := c.roomUsecase.Create(room); err != nil {
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
// @Summary								Update room
// @Description						Update room
// @Tags									room
// @Accept								json
// @produce								json
// @Param									id path int true "Room ID"
// @Param									room body domain.ROOM true "Room"
// @Router /api/room/{id}	[put]
func (c *roomController) Update(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
		})
	}

	var body payload.RoomUpdateDTO
	if err := c.validator.ValidateBody(ctx, &body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INVALID_BODY,
		})
	}

	data := &domain.ROOM{
		ID:           id,
		ROOM_NUMBER:  body.ROOM_NUMBER,
		IS_ACTIVE:    body.IS_ACTIVE,
		ROOM_TYPE_ID: body.ROOM_TYPE_ID,
	}

	if err := c.roomUsecase.Update(data); err != nil {
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
// @Summary								Update room is active
// @Description						Update room is active
// @Tags									room
// @Accept								json
// @produce								json
// @Param									id path int true "Room ID"
// @Param									isActive body payload.RoomUpdateRoomIsActiveDTO true "Is Active"
// @Router /api/room/active [post]
func (c *roomController) UpdateIsActive(ctx *fiber.Ctx) error {
	var body payload.RoomUpdateRoomIsActiveDTO
	if err := c.validator.ValidateBody(ctx, &body); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
		})
	}

	if err := c.roomUsecase.UpdateIsActive(body.ID, body.IsActive); err != nil {
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
