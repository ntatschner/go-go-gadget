package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/ntatschner/go-go-gadget/src/service"
)

// Package level variables
var listenOnPort string = os.Getenv("LISTENONPORT")

// Definition of available services
type Service struct {
	Name          string `json:"name" xml:"name"`
	ServiceName   string `json:"servicen_ame" xml:"servicename"`
	SecurityLevel int    `json:"security_level" xml:"securitylevel"`
	ServiceType   string `json:"service_type" xml:"servicetype"`
	ServiceID     string `json:"service_id" xml:"serviceid"`
}

type ServiceHandlers struct {
	service service.ServiceService
}

func (sh *ServiceHandlers) getAllServices(c *gin.Context) {
	services, err := sh.service.GetAllService()
	if err != nil {
		c.Request.Header.Add("Content-Type", "application/json")
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, services)
}

func (sh *ServiceHandlers) getService(c *gin.Context) {
	serviceID := c.Param("service_id")
	service, err := sh.service.GetService(serviceID)
	log.Output(1, fmt.Sprintf("Service ID: %s", serviceID))
	if err != nil {
		c.Request.Header.Add("Content-Type", "application/json")
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Service not found."})
		return
	} else {
		re := regexp.MustCompile(`^[0-9]{1,4}$`)
		if re.MatchString(serviceID) {
			c.Request.Header.Add("Content-Type", "application/json")
			c.IndentedJSON(http.StatusOK, service)
		} else {
			c.Request.Header.Add("Content-Type", "application/json")
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Invalid Service ID."})
		}
	}

}
