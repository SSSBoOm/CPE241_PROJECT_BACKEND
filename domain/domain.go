package domain

type Repository struct {
	UserRepository        UserRepository
	SessionRepository     SessionRepository
	RoleRepository        RoleRepository
	PaymentRepository     PaymentRepository
	PaymentTypeRepository PaymentTypeRepository
}

type Usecase struct {
	AuthUsecase        AuthUsecase
	GoogleUsecase      GoogleUsecase
	UserUsecase        UserUsecase
	SessionUsecase     SessionUsecase
	RoleUsecase        RoleUsecase
	PaymentUsecase     PaymentUsecase
	PaymentTypeUsecase PaymentTypeUsecase
}
