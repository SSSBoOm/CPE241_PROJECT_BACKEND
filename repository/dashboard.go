package repository

import (
	"fmt"
	"time"

	"github.com/SSSBoOm/CPE241_Project_Backend/domain"
	"github.com/jmoiron/sqlx"
)

type dashboardRepository struct {
	db *sqlx.DB
}

func NewDashboardRepository(db *sqlx.DB) *dashboardRepository {
	return &dashboardRepository{db}
}

func (r *dashboardRepository) GetDashboardRoomTypeReservation(startDate time.Time, endDate time.Time) (*[]domain.DashboardReservation, error) {
	var dashboard []domain.DashboardReservation
	query := `
	SELECT
    room_type.id, 
    room_type.name, 
    SUM(reservation.price) AS total
	FROM
    reservation
    INNER JOIN
    room
    ON 
        reservation.room_id = room.id
    INNER JOIN
    room_type
    ON 
        room.room_type_id = room_type.id
	WHERE
    reservation.start_date BETWEEN ? AND ?
	GROUP BY
    room_type.id
	HAVING
    SUM(reservation.price) > 0;
	`

	err := r.db.Select(&dashboard, query, startDate, endDate)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &dashboard, nil
}

func (r *dashboardRepository) GetDashboardServiceTypeReservation(startDate time.Time, endDate time.Time) (*[]domain.DashboardReservation, error) {
	var dashboard []domain.DashboardReservation
	query := `
	SELECT
    SUM( reservation.price ) AS total,
    service_type.id,
    service_type.name
	FROM
    reservation
    INNER JOIN service ON reservation.service_id = service.id
    INNER JOIN service_type ON service.service_type_id = service_type.id 
	WHERE
    reservation.start_date BETWEEN ? AND ?
	GROUP BY
    service_type.id 
	HAVING
    SUM( reservation.price ) > 0  
	ORDER BY service_type.id ASC
	`

	err := r.db.Select(&dashboard, query, startDate, endDate)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &dashboard, nil
}

func (r *dashboardRepository) GetDashboardReservationByPaymentType(startDate time.Time, endDate time.Time) (*[]domain.DashboardReservation2, error) {
	var dashboard []domain.DashboardReservation2
	query := `
	SELECT
    COUNT(payment_type.id) AS total, 
    payment_type.payment_type_name AS name
	FROM
    reservation
    INNER JOIN payment ON reservation.payment_info_id = payment.id
    INNER JOIN payment_type ON payment.payment_type_id = payment_type.id
	WHERE
    reservation.start_date BETWEEN ? AND ?
	GROUP BY
    payment_type.payment_type_name
	ORDER BY
    COUNT(payment_type.id) ASC;
	`

	err := r.db.Select(&dashboard, query, startDate, endDate)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &dashboard, nil
}
