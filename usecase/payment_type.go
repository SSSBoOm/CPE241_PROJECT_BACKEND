package usecase

import (
	"github.com/SSSBoOm/CPE241_Project_Backend/domain"
)

type paymentTypeUsecase struct {
	paymentTypeRepository domain.PaymentTypeRepository
}

func NewPaymentTypeUsecase(paymentTypeRepository domain.PaymentTypeRepository) domain.PaymentTypeUsecase {
	return &paymentTypeUsecase{
		paymentTypeRepository: paymentTypeRepository,
	}
}

func (u *paymentTypeUsecase) GetAll() (*[]domain.PaymentType, error) {
	return u.paymentTypeRepository.GetAll()
}

func (u *paymentTypeUsecase) GetByID(id int) (*domain.PaymentType, error) {
	return u.paymentTypeRepository.GetByID(id)
}

func (u *paymentTypeUsecase) Create(paymentType *domain.PaymentType) error {
	return u.paymentTypeRepository.Create(paymentType)
}

func (u *paymentTypeUsecase) Update(paymentType *domain.PaymentType) error {
	return u.paymentTypeRepository.Update(paymentType)
}
