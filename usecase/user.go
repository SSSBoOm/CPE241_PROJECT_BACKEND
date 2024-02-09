package usecase

import (
	"strings"
	"time"

	"github.com/SSSBoOm/CPE241_Project_Backend/domain"
	"github.com/google/uuid"
)

type UserUsecase struct {
	userRepository domain.UserRepository
}

func NewUserUsecase(userRepository domain.UserRepository) domain.UserUsecase {
	return &UserUsecase{userRepository: userRepository}
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
		Id:         uuid.NewString(),
		Email:      email,
		FirstName:  FirstName,
		LastName:   LastName,
		ProfileUrl: picture,
		CreatedAt:  time.Now(),
	}

	err := u.userRepository.CreateFromGoogle(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserUsecase) FindById(id string) (*domain.User, error) {
	user, err := u.userRepository.FindById(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserUsecase) FindByEmail(email string) (*domain.User, error) {
	user, err := u.userRepository.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	return user, nil
}
