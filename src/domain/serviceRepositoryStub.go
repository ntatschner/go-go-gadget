package domain

import (
	"log"
)

type ServiceRepositoryStub struct {
	Services []Service
}

type ServiceRepositoryStubById struct {
	Services *Service
}

func (s ServiceRepositoryStub) FindAll() ([]Service, error) {
	return s.Services, nil
}
func (s ServiceRepositoryStubById) ById(id string) (*Service, error) {
	return s.Services, nil
}

func NewServiceRepositoryStub() ServiceRepositoryStub {
	services := []Service{
		{
			Name:          "Proxmox",
			ServiceName:   "proxmox",
			SecurityLevel: 5,
			ServiceType:   "HyperVisor",
			ServiceID:     "0001",
		},
		{
			Name:          "Voyager Docker Host",
			ServiceName:   "voyager-docker-host",
			SecurityLevel: 5,
			ServiceType:   "ContainerHost",
			ServiceID:     "0002",
		},
		{
			Name:          "Enterprise Docker Host",
			ServiceName:   "enterprise-docker-host",
			SecurityLevel: 5,
			ServiceType:   "ContainerHost",
			ServiceID:     "0003",
		},
		{
			Name:          "UniFi Dream Machine Pro",
			ServiceName:   "unifi-dream-machine-pro",
			SecurityLevel: 5,
			ServiceType:   "Networking",
			ServiceID:     "0004",
		},
	}
	return ServiceRepositoryStub{services}
}

func NewServiceRepositoryStubById(id string) ServiceRepositoryStubById {
	log.Output(1, "ById() called")
	s, _ := NewServiceRepositoryStub().FindAll()

	var targetService Service
	for _, service := range s {
		if service.ServiceID == id {
			targetService = service
			break
		}
	}
	return ServiceRepositoryStubById{&targetService}
}
