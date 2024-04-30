package domain

type Repository struct {
	UserRepository    UserRepository
	SessionRepository SessionRepository
	RoleRepository    RoleRepository
}

type Usecase struct {
	AuthUsecase    AuthUsecase
	GoogleUsecase  GoogleUsecase
	UserUsecase    UserUsecase
	SessionUsecase SessionUsecase
	RoleUsecase    RoleUsecase
}
