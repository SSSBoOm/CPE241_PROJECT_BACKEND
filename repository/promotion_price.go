package repository

import (
	"time"

	"github.com/SSSBoOm/CPE241_Project_Backend/domain"
	"github.com/jmoiron/sqlx"
)

type promotionPriceRepo struct {
	db *sqlx.DB
}

func NewPromotionPriceRepository(db *sqlx.DB) domain.PromotionPriceRepository {
	return &promotionPriceRepo{
		db: db,
	}
}

func (r *promotionPriceRepo) GetAll() (*[]domain.PROMOTION_PRICE, error) {
	var promotionPrices []domain.PROMOTION_PRICE
	err := r.db.Select(&promotionPrices, "SELECT * FROM promotion_price")
	if err != nil {
		return nil, err
	}

	return &promotionPrices, nil
}

func (r *promotionPriceRepo) GetByID(id int) (*domain.PROMOTION_PRICE, error) {
	var promotionPrice domain.PROMOTION_PRICE
	err := r.db.Get(&promotionPrice, "SELECT * FROM promotion_price WHERE id = ?", id)
	if err != nil {
		return nil, err
	}

	return &promotionPrice, nil
}

func (r *promotionPriceRepo) Create(promotionPrice *domain.PROMOTION_PRICE) (*int, error) {
	t := r.db.MustBegin()
	data, err := t.NamedExec("INSERT INTO promotion_price (name, price, start_date, end_date) VALUES (:name, :price, :start_date, :end_date)", promotionPrice)
	if err != nil {
		t.Rollback()
		return nil, err
	}
	t.Commit()
	row, _ := data.LastInsertId()
	Id := int(row)
	return &Id, nil
}

func (r *promotionPriceRepo) Update(promotionPrice *domain.PROMOTION_PRICE) (*int, error) {
	t := r.db.MustBegin()
	promotionPrice.UPDATED_AT = time.Now()
	data, err := t.NamedExec("UPDATE promotion_price SET name = :name, price = :price, start_date = :startDate, end_date = :endDate, updated_at = :updated_at WHERE id = :id", promotionPrice)
	if err != nil {
		t.Rollback()
		return nil, err
	}
	t.Commit()
	row, _ := data.LastInsertId()
	Id := int(row)
	return &Id, nil
}
