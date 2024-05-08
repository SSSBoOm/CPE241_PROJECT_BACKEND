package domain

type Repository struct {
	UserRepository        UserRepository
	SessionRepository     SessionRepository
	RoleRepository        RoleRepository
	RoomTypeRepository    RoomTypeRepository
	PaymentRepository     PaymentRepository
	PaymentTypeRepository PaymentTypeRepository
}

type Usecase struct {
	AuthUsecase        AuthUsecase
	GoogleUsecase      GoogleUsecase
	UserUsecase        UserUsecase
	SessionUsecase     SessionUsecase
	RoleUsecase        RoleUsecase
	RoomTypeUsecase    RoomTypeUsecase
	PaymentUsecase     PaymentUsecase
	PaymentTypeUsecase PaymentTypeUsecase
}
