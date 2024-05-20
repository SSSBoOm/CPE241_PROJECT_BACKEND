package usecase

import "github.com/SSSBoOm/CPE241_Project_Backend/domain"

type serviceTypeUsecase struct {
	serviceTypeRepo domain.ServiceTypeRepository
}

func NewServiceTypeUsecase(serviceTypeRepo domain.ServiceTypeRepository) domain.ServiceTypeUsecase {
	return &serviceTypeUsecase{serviceTypeRepo}
}

func (u *serviceTypeUsecase) GetAll() (*[]domain.SERVICE_TYPE, error) {
	return u.serviceTypeRepo.GetAll()
}

func (u *serviceTypeUsecase) GetByID(id int) (*domain.SERVICE_TYPE, error) {
	return u.serviceTypeRepo.GetByID(id)
}

func (u *serviceTypeUsecase) Create(serviceType *domain.SERVICE_TYPE) error {
	return u.serviceTypeRepo.Create(serviceType)
}

func (u *serviceTypeUsecase) Update(serviceType *domain.SERVICE_TYPE) error {
	return u.serviceTypeRepo.Update(serviceType)
}

func (u *serviceTypeUsecase) UpdateIsActive(id int, isActive bool) error {
	return u.serviceTypeRepo.UpdateIsActive(id, isActive)
}
