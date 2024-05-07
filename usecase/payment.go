package usecase

import (
	"errors"

	"github.com/SSSBoOm/CPE241_Project_Backend/domain"
)

type paymentUsecase struct {
	paymentRepo        domain.PaymentRepository
	paymentTypeUsecase domain.PaymentTypeUsecase
}

func NewPaymentUsecase(paymentRepo domain.PaymentRepository, paymentTypeUsecase domain.PaymentTypeUsecase) domain.PaymentUsecase {
	return &paymentUsecase{
		paymentRepo:        paymentRepo,
		paymentTypeUsecase: paymentTypeUsecase,
	}
}

func (u *paymentUsecase) GetAll() (*[]domain.Payment, error) {
	return u.paymentRepo.GetAll()
}

func (u *paymentUsecase) GetByID(id int) (*domain.Payment, error) {
	return u.paymentRepo.GetByID(id)
}

func (u *paymentUsecase) Create(payment *domain.Payment) error {
	if paymentType, err := u.paymentTypeUsecase.GetByID(payment.PAYMENT_TYPE_ID); err != nil || paymentType == nil {
		return errors.New("payment type not found")
	}
	return u.paymentRepo.Create(payment)
}
