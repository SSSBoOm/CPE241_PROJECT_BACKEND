package domain

import "time"

type PROMOTION_PRICE struct {
	ID                        int                          `json:"id" db:"id"`
	NAME                      string                       `json:"name" db:"name"`
	PRICE                     float64                      `json:"price" db:"price"`
	START_DATE                time.Time                    `json:"startDate" db:"start_date"`
	END_DATE                  time.Time                    `json:"endDate" db:"end_date"`
	CREATED_AT                time.Time                    `json:"createdAt" db:"created_at"`
	UPDATED_AT                time.Time                    `json:"updatedAt" db:"updated_at"`
	ROOM_TYPE_PROMOTION_PRICE *[]ROOM_TYPE_PROMOTION_PRICE `json:"roomTypePromotionPrice" db:"-"`
}

type PromotionPriceRepository interface {
	GetAll() (*[]PROMOTION_PRICE, error)
	GetByID(id int) (*PROMOTION_PRICE, error)
	Create(promotionPrice *PROMOTION_PRICE) (*int, error)
	Update(promotionPrice *PROMOTION_PRICE) (*int, error)
}

type PromotionPriceUsecase interface {
	GetAll() (*[]PROMOTION_PRICE, error)
	GetByID(id int) (*PROMOTION_PRICE, error)
	Create(promotionPrice *PROMOTION_PRICE) (*int, error)
	Update(promotionPrice *PROMOTION_PRICE) (*int, error)
}
