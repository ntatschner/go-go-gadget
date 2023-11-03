package domain

// Definition of available services
type Service struct {
	Name          string `json:"name" xml:"name"`
	ServiceName   string `json:"servicen_ame" xml:"servicename"`
	SecurityLevel int    `json:"security_level" xml:"securitylevel"`
	ServiceType   string `json:"service_type" xml:"servicetype"`
	ServiceID     string `json:"service_id" xml:"serviceid"`
}

type ServiceRepository interface {
	FindAll() ([]Service, error)
}
