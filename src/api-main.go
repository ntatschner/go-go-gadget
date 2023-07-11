package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

// Package level variables
var listenOnPort string = os.Getenv("LISTENONPORT")

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

// Main package function
func main() {
	if os.Getenv("POSTGRESHOST") == "" {
		log.Fatal("Postgres host not specified, ensure POSTGRESHOST is specified on container run")
		os.Exit(1)
	}
	// establish database connection
	databaseConnection, err := connectDatabase(os.Getenv("POSTGRESHOST"), os.Getenv("DATABASE"), os.Getenv("POSTGRESUSERNAME"), os.Getenv("POSTGRESPASSWORD"))
	if err != nil {
		log.Fatalf("Failed to connect to the database. Error: %p", err)
		os.Exit(1)
	}
	// Test if require tables are present
	validateTables(databaseConnection, "Services")
	// prepare and launch http server
	router := gin.Default()
	router.GET("/services", getServices)

	router.Run(":" + listenOnPort)
}

// Database connection functions
func connectDatabase(URL string, databasename string, username string, password string) (db *sql.DB, exception error) {
	connStr := "user=username dbname=databasename sslmode=verify-full"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
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

// getServices responds with the list of all services as JSON.
func getServices(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, services)
}
