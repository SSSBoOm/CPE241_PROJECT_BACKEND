package controller

import (
	"strconv"

	"github.com/SSSBoOm/CPE241_Project_Backend/domain"
	"github.com/SSSBoOm/CPE241_Project_Backend/internal/constant"
	"github.com/gofiber/fiber/v2"
)

type roomController struct {
	roomUsecase domain.RoomUsecase
}

func NewRoomController(roomUsecase domain.RoomUsecase) *roomController {
	return &roomController{
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
// @Param									room body domain.Room true "Room"
// @Router /api/room	[post]
func (c *roomController) Create(ctx *fiber.Ctx) error {
	var room domain.Room
	if err := ctx.BodyParser(&room); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INVALID_BODY,
		})
	}

	if err := c.roomUsecase.Create(&room); err != nil {
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
