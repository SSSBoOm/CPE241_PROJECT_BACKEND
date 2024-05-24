package domain

type Repository struct {
	UserRepository            UserRepository
	SessionRepository         SessionRepository
	RoleRepository            RoleRepository
	RoomRepository            RoomRepository
	RoomTypeRepository        RoomTypeRepository
	PaymentRepository         PaymentRepository
	PaymentTypeRepository     PaymentTypeRepository
	MaintenanceRepository     MaintenanceRepository
	MaintenanceLogRepository  MaintenanceLogRepository
	ReservationRepository     ReservationRepository
	ReservationTaskRepository ReservationTaskRepository
	ServiceRepository         ServiceRepository
	ServiceTypeRepository     ServiceTypeRepository
}

type Usecase struct {
	AuthUsecase            AuthUsecase
	GoogleUsecase          GoogleUsecase
	UserUsecase            UserUsecase
	SessionUsecase         SessionUsecase
	RoleUsecase            RoleUsecase
	RoomUsecase            RoomUsecase
	RoomTypeUsecase        RoomTypeUsecase
	PaymentUsecase         PaymentUsecase
	PaymentTypeUsecase     PaymentTypeUsecase
	MaintenanceUsecase     MaintenanceUsecase
	MaintenanceLogUsecase  MaintenanceLogUsecase
	ReservationUsecase     ReservationUsecase
	ReservationTaskUsecase ReservationTaskUsecase
	ServiceUsecase         ServiceUsecase
	ServiceTypeUsecase     ServiceTypeUsecase
}
