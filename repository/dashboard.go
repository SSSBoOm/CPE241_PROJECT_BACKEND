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

func (r *dashboardRepository) GetDashboardReservation(startDate time.Time, endDate time.Time) (*[]domain.DashboardReservation, error) {
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
