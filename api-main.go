package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Package level variables
var listenOnIPv4Address string = "localhost"
var listenOnPort string = "8080"

// Definition of available services
type service struct {
	Name          string `json:"name"`
	ServiceName   string `json:"servicename"`
	SecurityLevel int    `json:"securitylevel"`
	ServiceType   string `json:"servicetype"`
}

// For testing purposes i'm creating in-memory data instead of retrieving from database
var services = []service{
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
		ServiceName:   "enterprise-dicker-host",
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

// Main package function
func main() {
	router := gin.Default()
	router.GET("/services", getServices)

	router.Run(listenOnIPv4Address + ":" + listenOnPort)
}

// getServices responds with the list of all services as JSON.
func getServices(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, services)
}
