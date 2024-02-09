package domain

type Repository struct {
	User UserRepository
}

type Usecase struct {
	AuthUsecase   AuthUsecase
	GoogleUsecase GoogleUsecase
	UserUsecase   UserUsecase
}
