package usecase

import (
	"fmt"
	"strings"
	"time"

	"github.com/SSSBoOm/CPE241_Project_Backend/domain"
	"github.com/SSSBoOm/CPE241_Project_Backend/internal/constant"
	"github.com/google/uuid"
)

type UserUsecase struct {
	userRepository domain.UserRepository
	roleRepository domain.RoleRepository
}

func NewUserUsecase(userRepository domain.UserRepository, roleRepository domain.RoleRepository) domain.UserUsecase {
	return &UserUsecase{userRepository: userRepository, roleRepository: roleRepository}
}

func (u *UserUsecase) CreateFromGoogle(name string, email string, picture string) (*domain.User, error) {
	FirstName := ""
	LastName := ""
	result := strings.Split(name, " ")
	if len(result) == 2 {
		FirstName = result[0]
		LastName = result[1]
	}

	user := &domain.User{
		ID:          uuid.NewString(),
		EMAIL:       email,
		FIRST_NAME:  FirstName,
		LAST_NAME:   LastName,
		PROFILE_URL: picture,
		CREATED_AT:  time.Now(),
	}

	err := u.userRepository.CreateFromGoogle(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserUsecase) GetAll() (*[]domain.User, error) {
	return u.userRepository.GetAll()
}

func (u *UserUsecase) FindById(id string) (*domain.User, error) {
	return u.userRepository.FindById(id)
}

func (u *UserUsecase) FindByEmail(email string) (*domain.User, error) {
	user, err := u.userRepository.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserUsecase) UpdateInfomation(user *domain.User) error {
	err := u.userRepository.UpdateInfomation(user)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserUsecase) UpdateRoleById(id string, roleID int) error {
	user, err := u.userRepository.FindById(id)
	if err != nil {
		return err
	} else if user == nil {
		return fmt.Errorf(constant.MESSAGE_NOT_FOUND)
	}
	role, err := u.roleRepository.Get(roleID)
	if err != nil {
		return err
	} else if role == nil {
		return fmt.Errorf("role not found")
	}
	err = u.userRepository.UpdateRoleById(id, roleID)
	if err != nil {
		return err
	}

	return nil
}
