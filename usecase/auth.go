package usecase

import (
	"github.com/SSSBoOm/github.com/SSSBoOm/CPE241_Project_Backend/domain"
	"github.com/gofiber/fiber/v2"
)

type authUsecase struct {
	googleUsecase domain.GoogleUsecase
	userUsecase   domain.UserUsecase
}

func NewAuthUsecase(
	googleUsecase domain.GoogleUsecase,
	userUsecase domain.UserUsecase,
) domain.AuthUsecase {
	return &authUsecase{
		googleUsecase: googleUsecase,
		userUsecase:   userUsecase,
	}
}

func (u *authUsecase) GoogleLogin(c *fiber.Ctx) (*domain.User, *fiber.Cookie, error) {
	token, err := u.googleUsecase.GetToken(c)
	if err != nil {
		return nil, nil, err
	}
	profile, err := u.googleUsecase.GetProfile(token.AccessToken)
	if err != nil {
		return nil, nil, err
	}

	// Find and Check User
	user, err := u.userUsecase.FindByEmail(profile.Email)
	if err != nil {
		return nil, nil, err
	}
	if user == nil {
		user, err = u.userUsecase.CreateFromGoogle(profile)
		if err != nil {
			return nil, nil, err
		}
	}

	// cookie, err := u.sessionUsecase.Create(user.Id, c.IP())
	// if err != nil {
	// 	return nil, nil, err
	// }

	return user, nil, nil
}
