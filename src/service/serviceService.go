package service

import "github.com/ntatschner/go-go-gadget/src/domain"

type ServiceService interface {
	GetAllService() ([]domain.Service, error)
}

type DefaultServiceService struct {
	repo domain.ServiceRepository
}

func (s DefaultServiceService) GetAllService() ([]domain.Service, error) {
	return s.repo.FindAll()
}

func NewServiceService(repo domain.ServiceRepository) DefaultServiceService {
	return DefaultServiceService{repo}
}
