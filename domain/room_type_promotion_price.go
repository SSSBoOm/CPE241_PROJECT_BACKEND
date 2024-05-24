package domain

type ROOM_TYPE_PROMOTION_PRICE struct {
	ID                 int    `json:"id" db:"id"`
	PROMOTION_PRICE_ID int    `json:"promotionPriceId" db:"promotion_price_id"`
	ROOM_TYPE_ID       int    `json:"roomTypeId" db:"room_type_id"`
	IS_ACTIVE          bool   `json:"isActive" db:"is_active"`
	CREATED_AT         string `json:"createdAt" db:"created_at"`
	UPDATED_AT         string `json:"updatedAt" db:"updated_at"`
}
