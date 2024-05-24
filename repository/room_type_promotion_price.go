package repository

import (
	"time"

	"github.com/SSSBoOm/CPE241_Project_Backend/domain"
	"github.com/jmoiron/sqlx"
)

type roomTypePromotionPriceRepository struct {
	db *sqlx.DB
}

func NewRoomTypePromotionPriceRepository(db *sqlx.DB) domain.RoomTypePromotionPriceRepository {
	return &roomTypePromotionPriceRepository{
		db: db,
	}
}

func (r *roomTypePromotionPriceRepository) GetAll() (*[]domain.ROOM_TYPE_PROMOTION_PRICE, error) {
	var roomTypePromotionPrices []domain.ROOM_TYPE_PROMOTION_PRICE
	err := r.db.Select(&roomTypePromotionPrices, "SELECT * FROM room_type_promotion_price")
	if err != nil {
		return nil, err
	}
	return &roomTypePromotionPrices, nil
}

func (r *roomTypePromotionPriceRepository) GetByID(id int) (*domain.ROOM_TYPE_PROMOTION_PRICE, error) {
	var roomTypePromotionPrice domain.ROOM_TYPE_PROMOTION_PRICE
	err := r.db.Get(&roomTypePromotionPrice, "SELECT * FROM room_type_promotion_price WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	return &roomTypePromotionPrice, nil
}

func (r *roomTypePromotionPriceRepository) GetByRoomTypeID(roomTypeID int) (*[]domain.ROOM_TYPE_PROMOTION_PRICE, error) {
	var roomTypePromotionPrices []domain.ROOM_TYPE_PROMOTION_PRICE
	err := r.db.Select(&roomTypePromotionPrices, "SELECT * FROM room_type_promotion_price WHERE room_type_id = ?", roomTypeID)
	if err != nil {
		return nil, err
	}
	return &roomTypePromotionPrices, nil
}

func (r *roomTypePromotionPriceRepository) GetByPromotionPriceID(promotionPriceID int) (*[]domain.ROOM_TYPE_PROMOTION_PRICE, error) {
	var roomTypePromotionPrices []domain.ROOM_TYPE_PROMOTION_PRICE
	err := r.db.Select(&roomTypePromotionPrices, "SELECT * FROM room_type_promotion_price WHERE promotion_price_id = ?", promotionPriceID)
	if err != nil {
		return nil, err
	}
	return &roomTypePromotionPrices, nil
}

func (r *roomTypePromotionPriceRepository) Create(roomTypePromotionPrice *domain.ROOM_TYPE_PROMOTION_PRICE) (*int, error) {
	t := r.db.MustBegin()
	roomTypePromotionPrice.CREATED_AT = time.Now()
	roomTypePromotionPrice.UPDATED_AT = time.Now()
	row, err := t.NamedExec("INSERT INTO room_type_promotion_price (promotion_price_id, room_type_id, is_active, created_at, updated_at) VALUES (:promotion_price_id, :room_type_id, :is_active, :created_at, :updated_at)", roomTypePromotionPrice)
	if err != nil {
		t.Rollback()
		return nil, err
	}
	t.Commit()
	LastInsertId, _ := row.LastInsertId()
	Id := int(LastInsertId)
	return &Id, nil
}

func (r *roomTypePromotionPriceRepository) Update(roomTypePromotionPrice *domain.ROOM_TYPE_PROMOTION_PRICE) error {
	t := r.db.MustBegin()
	roomTypePromotionPrice.UPDATED_AT = time.Now()
	_, err := t.NamedExec("UPDATE room_type_promotion_price SET promotion_price_id = :promotion_price_id, room_type_id = :room_type_id, is_active = :is_active, updated_at = :updated_at WHERE id = :id", roomTypePromotionPrice)
	if err != nil {
		t.Rollback()
		return err
	}
	t.Commit()
	return nil
}
