package repository

import (
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
	row, err := t.NamedExec("INSERT INTO room_type_promotion_price (promotion_price_id, room_type_id, is_active, created_at, updated_at) VALUES (:promotionPriceId, :roomTypeId, :isActive, :createdAt, :updatedAt)", roomTypePromotionPrice)
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
	_, err := t.NamedExec("UPDATE room_type_promotion_price SET promotion_price_id = :promotionPriceId, room_type_id = :roomTypeId, is_active = :isActive, created_at = :createdAt, updated_at = :updatedAt WHERE id = :id", roomTypePromotionPrice)
	if err != nil {
		t.Rollback()
		return err
	}
	t.Commit()
	return nil
}
