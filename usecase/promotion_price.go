package usecase

import "github.com/SSSBoOm/CPE241_Project_Backend/domain"

type promotionPriceUsecase struct {
	promotionPriceRepo domain.PromotionPriceRepository
}

func NewPromotionPriceUsecase(promotionPriceRepo domain.PromotionPriceRepository) domain.PromotionPriceUsecase {
	return &promotionPriceUsecase{
		promotionPriceRepo: promotionPriceRepo,
	}
}

func (u *promotionPriceUsecase) GetAll() (*[]domain.PROMOTION_PRICE, error) {
	return u.promotionPriceRepo.GetAll()
}

func (u *promotionPriceUsecase) GetByID(id int) (*domain.PROMOTION_PRICE, error) {
	return u.promotionPriceRepo.GetByID(id)
}

func (u *promotionPriceUsecase) Create(promotionPrice *domain.PROMOTION_PRICE) (*domain.PROMOTION_PRICE, error) {
	return u.promotionPriceRepo.Create(promotionPrice)
}

func (u *promotionPriceUsecase) Update(promotionPrice *domain.PROMOTION_PRICE) (*domain.PROMOTION_PRICE, error) {
	return u.promotionPriceRepo.Update(promotionPrice)
}
