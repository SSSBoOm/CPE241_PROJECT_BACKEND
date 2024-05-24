package usecase

import (
	"sync"

	"github.com/SSSBoOm/CPE241_Project_Backend/domain"
)

type RoomTypePromotionPriceUsecase struct {
	RoomTypePromotionPriceRepository domain.RoomTypePromotionPriceRepository
	roomTypeUsecase                  domain.RoomTypeUsecase
}

func NewRoomTypePromotionPriceUsecase(RoomTypePromotionPriceRepository domain.RoomTypePromotionPriceRepository, roomTypeUsecase domain.RoomTypeUsecase) domain.RoomTypePromotionPriceUsecase {
	return &RoomTypePromotionPriceUsecase{
		RoomTypePromotionPriceRepository: RoomTypePromotionPriceRepository,
		roomTypeUsecase:                  roomTypeUsecase,
	}
}

func (uc *RoomTypePromotionPriceUsecase) GetAll() (*[]domain.ROOM_TYPE_PROMOTION_PRICE, error) {
	return uc.RoomTypePromotionPriceRepository.GetAll()
}

func (uc *RoomTypePromotionPriceUsecase) GetByID(id int) (*domain.ROOM_TYPE_PROMOTION_PRICE, error) {
	return uc.RoomTypePromotionPriceRepository.GetByID(id)
}

func (uc *RoomTypePromotionPriceUsecase) GetByRoomTypeID(roomTypeID int) (*[]domain.ROOM_TYPE_PROMOTION_PRICE, error) {
	return uc.RoomTypePromotionPriceRepository.GetByRoomTypeID(roomTypeID)
}

func (uc *RoomTypePromotionPriceUsecase) GetByPromotionPriceID(promotionPriceID int) (*[]domain.ROOM_TYPE_PROMOTION_PRICE, error) {
	roomTypePromotionPrices, err := uc.RoomTypePromotionPriceRepository.GetByPromotionPriceID(promotionPriceID)
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup
	wg.Add(len(*roomTypePromotionPrices))
	for i, roomTypePromotionPrice := range *roomTypePromotionPrices {
		go func(i int, roomTypePromotionPrice domain.ROOM_TYPE_PROMOTION_PRICE) {
			defer wg.Done()
			roomType, err := uc.roomTypeUsecase.GetByID(roomTypePromotionPrice.ROOM_TYPE_ID)
			if err != nil {
				return
			}
			(*roomTypePromotionPrices)[i].ROOM_TYPE = roomType
		}(i, roomTypePromotionPrice)
	}
	wg.Wait()

	return roomTypePromotionPrices, nil
}

func (uc *RoomTypePromotionPriceUsecase) Create(roomTypePromotionPrice *domain.ROOM_TYPE_PROMOTION_PRICE) (*int, error) {
	return uc.RoomTypePromotionPriceRepository.Create(roomTypePromotionPrice)
}

func (uc *RoomTypePromotionPriceUsecase) Update(roomTypePromotionPrice *domain.ROOM_TYPE_PROMOTION_PRICE) error {
	return uc.RoomTypePromotionPriceRepository.Update(roomTypePromotionPrice)
}
