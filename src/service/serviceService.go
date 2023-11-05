package service

import "github.com/ntatschner/go-go-gadget/src/domain"

type ServiceService interface {
	GetAllService() ([]domain.Service, error)
	GetService(string) (*domain.Service, error)
}

type DefaultServiceService struct {
	repo domain.ServiceRepository
}

func (s DefaultServiceService) GetAllService() ([]domain.Service, error) {
	return s.repo.FindAll()
}

func (s DefaultServiceService) GetService(id string) (*domain.Service, error) {
	return s.repo.ById(id)
}

func NewServiceService(repo domain.ServiceRepository) DefaultServiceService {
	return DefaultServiceService{repo}
}
