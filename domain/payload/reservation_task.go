package payload

type ReservationTaskCreateDTO struct {
	RESERVATION_ID string `json:"reservation_id" validate:"required"`
	
}
