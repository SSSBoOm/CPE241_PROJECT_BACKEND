package controller

import (
	"github.com/SSSBoOm/CPE241_Project_Backend/domain"
	"github.com/SSSBoOm/CPE241_Project_Backend/internal/constant"
	"github.com/gofiber/fiber/v2"
)

type RoleController struct {
	roleUsecase domain.RoleUsecase
}

func NewRoleController(roleUsecase domain.RoleUsecase) *RoleController {
	return &RoleController{
		roleUsecase: roleUsecase,
	}
}

// GetALL godoc
// @Summary								Get all roles
// @Description						Get all roles
// @Tags									role
// @Accept								json
// @produce								json
// @Router /api/role/all	[get]
func (role *RoleController) GetALL(c *fiber.Ctx) error {
	roles, err := role.roleUsecase.GetAll()
	if err != nil {
		return c.Status(500).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
		})
	}

	return c.Status(200).JSON(domain.Response{
		SUCCESS: true,
		MESSAGE: "Hello World",
		DATA:    roles,
	})
}
