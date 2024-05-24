package usecase

import "github.com/SSSBoOm/CPE241_Project_Backend/domain"

type RoomTypePromotionPriceUsecase struct {
	RoomTypePromotionPriceRepository domain.RoomTypePromotionPriceRepository
}

func NewRoomTypePromotionPriceUsecase(RoomTypePromotionPriceRepository domain.RoomTypePromotionPriceRepository) domain.RoomTypePromotionPriceUsecase {
	return &RoomTypePromotionPriceUsecase{
		RoomTypePromotionPriceRepository: RoomTypePromotionPriceRepository,
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
	return uc.RoomTypePromotionPriceRepository.GetByPromotionPriceID(promotionPriceID)
}

func (uc *RoomTypePromotionPriceUsecase) Create(roomTypePromotionPrice *domain.ROOM_TYPE_PROMOTION_PRICE) (*int, error) {
	return uc.RoomTypePromotionPriceRepository.Create(roomTypePromotionPrice)
}

func (uc *RoomTypePromotionPriceUsecase) Update(roomTypePromotionPrice *domain.ROOM_TYPE_PROMOTION_PRICE) error {
	return uc.RoomTypePromotionPriceRepository.Update(roomTypePromotionPrice)
}
