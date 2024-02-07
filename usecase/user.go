package usecase

import (
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

func (u *UserUsecase) CreateFromGoogle(profile *domain.GoogleResponse) (*domain.User, error) {
	id := uuid.NewString()

	user := &domain.User{
		Id:         id,
		Email:      profile.Email,
		Prefix:     "",
		FirstName:  "",
		LastName:   "",
		ProfileUrl: profile.Picture,
		Phone:      "",
		CreatedAt:  time.Now(),
	}

	err := u.userRepository.Create(user)
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
