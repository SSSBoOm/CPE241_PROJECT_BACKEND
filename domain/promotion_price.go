package domain

type PROMOTION_PRICE struct {
	ID         int     `json:"id" db:"id"`
	NAME       string  `json:"name" db:"name"`
	PRICE      float64 `json:"price" db:"price"`
	START_DATE string  `json:"startDate" db:"start_date"`
	END_DATE   string  `json:"endDate" db:"end_date"`
	CREATED_AT string  `json:"createdAt" db:"created_at"`
	UPDATED_AT string  `json:"updatedAt" db:"updated_at"`
}

type PromotionPriceRepository interface {
	GetAll() ([]PROMOTION_PRICE, error)
	GetByID(id int) (PROMOTION_PRICE, error)
	Create(promotionPrice PROMOTION_PRICE) (PROMOTION_PRICE, error)
	Update(id int, promotionPrice PROMOTION_PRICE) (PROMOTION_PRICE, error)
}

type PromotionPriceUsecase interface {
	GetAll() ([]PROMOTION_PRICE, error)
	GetByID(id int) (PROMOTION_PRICE, error)
	Create(promotionPrice PROMOTION_PRICE) (PROMOTION_PRICE, error)
	Update(id int, promotionPrice PROMOTION_PRICE) (PROMOTION_PRICE, error)
}
