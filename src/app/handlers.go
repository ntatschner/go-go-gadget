package app

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

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

// getServices responds with the list of all services as JSON.
func getServices(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, services)
}

func postService(c *gin.Context, service string) {
	c.IndentedJSON(http.StatusAccepted, gin.H{"message": "Post Accepted."})
}
