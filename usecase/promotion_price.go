package usecase

import (
	"sync"

	"github.com/SSSBoOm/CPE241_Project_Backend/domain"
)

type promotionPriceUsecase struct {
	promotionPriceRepo            domain.PromotionPriceRepository
	roomTypePromotionPriceUsecase domain.RoomTypePromotionPriceUsecase
}

func NewPromotionPriceUsecase(promotionPriceRepo domain.PromotionPriceRepository, roomTypePromotionPriceUsecase domain.RoomTypePromotionPriceUsecase) domain.PromotionPriceUsecase {
	return &promotionPriceUsecase{
		promotionPriceRepo:            promotionPriceRepo,
		roomTypePromotionPriceUsecase: roomTypePromotionPriceUsecase,
	}
}

func (u *promotionPriceUsecase) GetAll() (*[]domain.PROMOTION_PRICE, error) {
	promotionPrices, err := u.promotionPriceRepo.GetAll()
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup
	wg.Add(len(*promotionPrices))
	for i, promotionPrice := range *promotionPrices {
		go func(i int, promotionPrice domain.PROMOTION_PRICE) {
			defer wg.Done()
			roomTypePromotionPrices, err := u.roomTypePromotionPriceUsecase.GetByPromotionPriceID(promotionPrice.ID)
			if err != nil {
				return
			}
			(*promotionPrices)[i].ROOM_TYPE_PROMOTION_PRICE = roomTypePromotionPrices
		}(i, promotionPrice)
	}
	wg.Wait()

	return promotionPrices, nil
}

func (u *promotionPriceUsecase) GetByID(id int) (*domain.PROMOTION_PRICE, error) {
	promotionPrice, err := u.promotionPriceRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	roomTypePromotionPrices, err := u.roomTypePromotionPriceUsecase.GetByPromotionPriceID(promotionPrice.ID)
	if err != nil {
		return nil, err
	}
	promotionPrice.ROOM_TYPE_PROMOTION_PRICE = roomTypePromotionPrices
	return promotionPrice, nil

}

func (u *promotionPriceUsecase) Create(promotionPrice *domain.PROMOTION_PRICE) (*int, error) {
	return u.promotionPriceRepo.Create(promotionPrice)
}

func (u *promotionPriceUsecase) Update(promotionPrice *domain.PROMOTION_PRICE) (*int, error) {
	return u.promotionPriceRepo.Update(promotionPrice)
}
