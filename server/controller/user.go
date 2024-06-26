package controller

import (
	"sync"

	"github.com/SSSBoOm/CPE241_Project_Backend/domain"
	"github.com/SSSBoOm/CPE241_Project_Backend/domain/payload"
	"github.com/SSSBoOm/CPE241_Project_Backend/internal/constant"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	validator          domain.ValidatorUsecase
	userUsecase        domain.UserUsecase
	paymentUsecase     domain.PaymentUsecase
	reservationUsecase domain.ReservationUsecase
}

func NewUserController(validator domain.ValidatorUsecase, userUsecase domain.UserUsecase, paymentUsecase domain.PaymentUsecase, reservationUsecase domain.ReservationUsecase) *UserController {
	return &UserController{
		validator:          validator,
		userUsecase:        userUsecase,
		paymentUsecase:     paymentUsecase,
		reservationUsecase: reservationUsecase,
	}
}

// Me	godoc
// @Summary								Get user information
// @Description						Get user information
// @Tags									user
// @Accept								json
// @produce								json
// @Security							ApiKeyAuth
// @Response 200 {object} domain.User
// @Router /api/user/me	[get]
func (u *UserController) Me(ctx *fiber.Ctx) error {
	id := ctx.Locals(constant.CTX_USER_ID).(string)
	user, err := u.userUsecase.FindById(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
		})
	} else if user == nil {
		return ctx.Status(fiber.StatusNotFound).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_NOT_FOUND,
		})
	}
	user.ID = ""

	return ctx.Status(fiber.StatusOK).JSON(domain.Response{
		SUCCESS: true,
		MESSAGE: constant.MESSAGE_SUCCESS,
		DATA:    user,
	})
}

// GetByID	godoc
// @Summary								Get user by id
// @Description						Get user by id
// @Tags									manage
// @Accept								json
// @produce								json
// @Security							ApiKeyAuth
// @Param									id path	string true	"User ID"
// @Response 200 {object} domain.User
// @Router /api/admin/manage/user/{id}	[get]
func (u *UserController) GetByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	user, err := u.userUsecase.FindById(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
		})
	} else if user == nil {
		return ctx.Status(fiber.StatusNotFound).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_NOT_FOUND,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(domain.Response{
		SUCCESS: true,
		MESSAGE: constant.MESSAGE_SUCCESS,
		DATA:    user,
	})
}

// UpdateByID	godoc
// @Summary								Update user by id
// @Description						Update user by id
// @Tags									manage
// @Accept								json
// @produce								json
// @Security							ApiKeyAuth
// @Param									payload body	payload.UpdateUserDTO true "Payload"
// @Response 200 {object} domain.Response
// @Router /api/admin/manage/user	[put]
func (u *UserController) UpdateByID(ctx *fiber.Ctx) error {
	var body payload.UpdateUserDTO
	err := u.validator.ValidateBody(ctx, &body)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INVALID_BODY,
		})
	}

	user := &domain.User{
		ID:         body.ID,
		PREFIX:     body.PREFIX,
		FIRST_NAME: body.FIRST_NAME,
		LAST_NAME:  body.LAST_NAME,
		DOB:        &body.DOB,
		PHONE:      body.PHONE,
		GENDER:     body.GENDER,
		ADDRESS:    body.ADDRESS,
	}

	err = u.userUsecase.UpdateInfomation(user)
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

// UpdateRoleByID	godoc
// @Summary								Update role by user id
// @Description						Update role by user id
// @Tags									manage
// @Accept								json
// @produce								json
// @Security							ApiKeyAuth
// @Param									payload body	payload.UpdateUserRoleDTO true "Payload"
// @Response 200 {object} domain.Response
// @Router /api/admin/manage/role	[put]
func (u *UserController) UpdateRoleByID(ctx *fiber.Ctx) error {
	var body payload.UpdateUserRoleDTO
	err := u.validator.ValidateBody(ctx, &body)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INVALID_BODY,
		})
	}

	err = u.userUsecase.UpdateRoleById(body.USER_ID, body.ROLE_ID)
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

// UpdateUserInformation	godoc
// @Summary								Update user information
// @Description						Update user information
// @Tags									user
// @Accept								json
// @produce								json
// @Security							ApiKeyAuth
// @Param									payload body	payload.UpdateUserInformationDTO true "Payload"
// @Response 200 {object} domain.Response
// @Router /api/user/			[patch]
func (u *UserController) UpdateInfomationByID(ctx *fiber.Ctx) error {
	var body payload.UpdateUserInformationDTO
	if err := u.validator.ValidateBody(ctx, &body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INVALID_BODY,
		})
	}

	user := &domain.User{
		ID:         ctx.Locals(constant.CTX_USER_ID).(string),
		PREFIX:     body.PREFIX,
		FIRST_NAME: body.FIRST_NAME,
		LAST_NAME:  body.LAST_NAME,
		DOB:        &body.DOB,
		PHONE:      body.PHONE,
		GENDER:     body.GENDER,
		ADDRESS:    body.ADDRESS,
	}

	if err := u.userUsecase.UpdateInfomation(user); err != nil {
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

// GetPaymentByUserID	godoc
// @Summary								Get payment by user id
// @Description						Get payment by user id
// @Tags									user
// @Accept								json
// @produce								json
// @Security							ApiKeyAuth
// @Response 200 {object} []domain.Payment
// @Router /api/user/payment	[get]
func (u *UserController) GetPaymentByUserID(ctx *fiber.Ctx) error {
	id := ctx.Locals(constant.CTX_USER_ID).(string)
	payments, err := u.paymentUsecase.GetByUserID(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(domain.Response{
		SUCCESS: true,
		MESSAGE: constant.MESSAGE_SUCCESS,
		DATA:    payments,
	})
}

// GetALL	godoc
// @Summary								Get all user
// @Description						Get all user
// @Tags									manage
// @Accept								json
// @produce								json
// @Security							ApiKeyAuth
// @Response 200 {object} []domain.User
// @Router /api/admin/manage/user	[get]
func (u *UserController) GetALL(ctx *fiber.Ctx) error {
	users, err := u.userUsecase.GetAll()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
		})
	}

	var wg sync.WaitGroup
	wg.Add(len(*users))
	for i := range *users {
		go func(i int) {
			defer wg.Done()
			reservation, err := u.reservationUsecase.GetByUserID((*users)[i].ID)
			if err != nil {
				return
			}
			(*users)[i].RESERVATIONS = reservation
		}(i)
	}
	wg.Wait()

	return ctx.Status(fiber.StatusOK).JSON(domain.Response{
		SUCCESS: true,
		MESSAGE: constant.MESSAGE_SUCCESS,
		DATA:    users,
	})
}

// GetAllStaff	godoc
// @Summary								Get all staff
// @Description						Get all staff
// @Tags									manage
// @Accept								json
// @produce								json
// @Security							ApiKeyAuth
// @Response 200 {object} []domain.User
// @Router /api/admin/manage/staff	[get]
func (u *UserController) GetAllStaff(ctx *fiber.Ctx) error {
	users, err := u.userUsecase.GetAllByRoleId(2)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(domain.Response{
		SUCCESS: true,
		MESSAGE: constant.MESSAGE_SUCCESS,
		DATA:    users,
	})
}
