package usecase

import (
	"sync"

	"github.com/SSSBoOm/CPE241_Project_Backend/domain"
)

type serviceTypeUsecase struct {
	serviceTypeRepo domain.ServiceTypeRepository
	serviceUsecase  domain.ServiceUsecase
}

func NewServiceTypeUsecase(serviceTypeRepo domain.ServiceTypeRepository, serviceUsecase domain.ServiceUsecase) domain.ServiceTypeUsecase {
	return &serviceTypeUsecase{
		serviceTypeRepo: serviceTypeRepo,
		serviceUsecase:  serviceUsecase,
	}
}

func (u *serviceTypeUsecase) GetAll() (*[]domain.SERVICE_TYPE, error) {
	serviceType, err := u.serviceTypeRepo.GetAll()
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup
	wg.Add(len(*serviceType))
	for i, item := range *serviceType {
		go func(i int, item domain.SERVICE_TYPE) {
			defer wg.Done()
			service, err := u.serviceUsecase.GetByServiceTypeId(item.ID)
			if err != nil {
				return
			}
			(*serviceType)[i].SERVICE = service
		}(i, item)
	}
	wg.Wait()

	return serviceType, nil
}

func (u *serviceTypeUsecase) GetByID(id int) (*domain.SERVICE_TYPE, error) {
	return u.serviceTypeRepo.GetByID(id)
}

func (u *serviceTypeUsecase) Create(serviceType *domain.SERVICE_TYPE) (*int, error) {
	id, err := u.serviceTypeRepo.Create(serviceType)
	if err != nil || id == nil {
		return nil, err
	}

	var wg sync.WaitGroup
	wg.Add(len(*serviceType.SERVICE))
	for _, service := range *serviceType.SERVICE {
		go func(service domain.SERVICE) {
			defer wg.Done()
			service.SERVICE_TYPE_ID = *id
			err := u.serviceUsecase.Create(&service)
			if err != nil {
				return
			}
		}(service)
	}
	wg.Wait()

	return id, nil
}

func (u *serviceTypeUsecase) Update(serviceType *domain.SERVICE_TYPE) error {
	return u.serviceTypeRepo.Update(serviceType)
}

func (u *serviceTypeUsecase) UpdateIsActive(id int, isActive bool) error {
	return u.serviceTypeRepo.UpdateIsActive(id, isActive)
}
