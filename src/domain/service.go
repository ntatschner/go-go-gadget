package domain

// Definition of available services
type Service struct {
	Name          string `json:"name" xml:"name"`
	ServiceName   string `json:"servicen_ame" xml:"servicename"`
	SecurityLevel int    `json:"security_level" xml:"securitylevel"`
	ServiceType   string `json:"service_type" xml:"servicetype"`
}

type ServiceRepository interface {
	FindAll() ([]Service, error)
}

type ServiceRepositoryStub struct {
	services []Service
}

func (s ServiceRepositoryStub) FindAll() ([]Service, error) {
	return s.services, nil
}

func NewServiceRepositoryStub() ServiceRepository {
	services := []Service{
		{
			Name:          "Proxmox",
			ServiceName:   "proxmox",
			SecurityLevel: 5,
			ServiceType:   "HyperVisor",
		},
		{
			Name:          "Voyager Docker Host",
			ServiceName:   "voyager-docker-host",
			SecurityLevel: 5,
			ServiceType:   "ContainerHost",
		},
		{
			Name:          "Enterprise Docker Host",
			ServiceName:   "enterprise-docker-host",
			SecurityLevel: 5,
			ServiceType:   "ContainerHost",
		},
		{
			Name:          "UniFi Dream Machine Pro",
			ServiceName:   "unifi-dream-machine-pro",
			SecurityLevel: 5,
			ServiceType:   "Networking",
		},
	}
}
