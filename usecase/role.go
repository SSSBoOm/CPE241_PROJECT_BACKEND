package usecase

import (
	"github.com/SSSBoOm/CPE241_Project_Backend/domain"
)

type roleUsecase struct {
	roleRepository domain.RoleRepository
}

func NewRoleUsecase(roleRepository domain.RoleRepository) domain.RoleUsecase {
	return &roleUsecase{
		roleRepository: roleRepository,
	}
}

func (u *roleUsecase) Get(id int) (*domain.Role, error) {
	return u.roleRepository.Get(id)
}
