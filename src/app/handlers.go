package app

import (
	"database/sql"
	"net/http"

	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ntatschner/go-go-gadget/src/service"
)

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
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, services)
}

func (sh *ServiceHandlers) getService(c *gin.Context) {
	serviceID := c.Param("service_id")
	services, err := sh.service.GetAllService()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	for _, a := range services {
		if a.ServiceID == serviceID {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Service not found."})
}

// func createService(c *gin.Context) {
// 	var newService Service

// 	// Call BindJSON to bind the received JSON to newService.
// 	if err := c.BindJSON(&newService); err != nil {
// 		return
// 	}

// 	// Add the new service to the slice.
// 	services = append(services, newService)
// 	c.IndentedJSON(http.StatusCreated, newService)
// }

// Package level variables
var listenOnPort string = os.Getenv("LISTENONPORT")

// Database connection functions
func connectDatabase(URL string, databasename string, username string, password string) (db *sql.DB, exception error) {
	connStr := "user=username dbname=databasename sslmode=verify-full"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func localOrDatabase() (source string, exception error) {
	if os.Getenv("POSTGRESHOST") == "" {
		source = "local"
	} else {
		source = "remote"
	}
	return
}

// Ensure required tables are present in the database and create the table if not
func validateTables(dbconnection *sql.DB, tablename ...string) (servicesTable *sql.Rows, exception error) {
	servicesTable, err := dbconnection.Query(`
		CREATE TABLE IF NOT EXISTS Services (
		Name VARCHAR(50) NOT NULL,
		ServiceName VARCHAR(50) NOT NULL,
		SecurityLevel INT NOT NULL,
		ServiceType VARCHAR(50) NOT NULL
		);`, tablename)
	if err != nil {
		log.Fatal("Failed to run validation, ensure the user has rights to the database")
		os.Exit(2)
	}
	return
}
