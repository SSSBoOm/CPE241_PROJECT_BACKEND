package usecase

import (
	"sync"

	"github.com/SSSBoOm/CPE241_Project_Backend/domain"
)

type serviceUsecase struct {
	serviceRepo        domain.ServiceRepository
	serviceTypeUsecase domain.ServiceTypeUsecase
}

func NewServiceUsecase(serviceRepo domain.ServiceRepository, serviceTypeUsecase domain.ServiceTypeUsecase) domain.ServiceUsecase {
	return &serviceUsecase{
		serviceRepo:        serviceRepo,
		serviceTypeUsecase: serviceTypeUsecase,
	}
}

func (u *serviceUsecase) GetAll() (*[]domain.SERVICE, error) {
	service, err := u.serviceRepo.GetAll()
	if err != nil {
		return nil, err
	}
	var wg sync.WaitGroup
	wg.Add(len(*service))
	for i, item := range *service {
		go func(i int, item domain.SERVICE) {
			defer wg.Done()
			serviceType, err := u.serviceTypeUsecase.GetByID(item.SERVICE_TYPE_ID)
			if err != nil || serviceType == nil {
				return
			}
			(*service)[i].SERVICE_TYPE = serviceType
		}(i, item)
	}
	wg.Wait()
	return service, nil
}

func (u *serviceUsecase) GetById(id int) (*domain.SERVICE, error) {
	service, err := u.serviceRepo.GetById(id)
	if err != nil {
		return nil, err
	}

	serviceType, err := u.serviceTypeUsecase.GetByID(service.SERVICE_TYPE_ID)
	if err != nil || serviceType == nil {
		return nil, err
	}
	service.SERVICE_TYPE = serviceType
	return service, nil
}

func (u *serviceUsecase) GetByServiceTypeId(serviceTypeId int) (*[]domain.SERVICE, error) {
	service, err := u.serviceRepo.GetByServiceTypeId(serviceTypeId)
	if err != nil {
		return nil, err
	}
	return service, nil
}
func (u *serviceUsecase) Create(service *domain.SERVICE) error {
	return u.serviceRepo.Create(service)
}

func (u *serviceUsecase) Update(service *domain.SERVICE) error {
	return u.serviceRepo.Update(service)
}

func (u *serviceUsecase) UpdateIsActive(id int, isActive bool) error {
	return u.serviceRepo.UpdateIsActive(id, isActive)
}
