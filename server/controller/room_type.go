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

func (controller *RoomTypeController) GetRoomTypeList(ctx *fiber.Ctx) error {
	roomTypeList, err := controller.roomTypeUsecase.GetAll()
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

func (controller *RoomTypeController) GetRoomTypeByID(ctx *fiber.Ctx) error {
	roomTypeID, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
		})
	}

	roomType, err := controller.roomTypeUsecase.GetByID(roomTypeID)
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

func (controller *RoomTypeController) CreateRoomType(ctx *fiber.Ctx) error {
	roomType := new(domain.RoomType)
	if err := ctx.BodyParser(roomType); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_BAD_REQUEST,
		})
	}

	err := controller.roomTypeUsecase.Create(roomType)
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

func (controller *RoomTypeController) UpdateRoomType(ctx *fiber.Ctx) error {
	payload := new(payload.UpdateRoomType)
	if err := ctx.BodyParser(payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_BAD_REQUEST,
		})
	}

	roomType := &domain.RoomType{
		ID:        payload.ID,
		NAME:      payload.NAME,
		IS_ACTIVE: payload.IS_ACTIVE,
	}

	err := controller.roomTypeUsecase.Update(roomType)
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
