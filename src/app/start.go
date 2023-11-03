package app

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func Start() {
	source, err := localOrDatabase()
	if source == "local" {
		log.Output(1, "Launching with local test data.")
		listenOnPort = "8080"
	}
	if source == "remote" {
		// establish database connection
		databaseConnection, err := connectDatabase(os.Getenv("POSTGRESHOST"), os.Getenv("DATABASE"), os.Getenv("POSTGRESUSERNAME"), os.Getenv("POSTGRESPASSWORD"))
		if err != nil {
			log.Fatalf("Failed to connect to the database. Error: %p", err)
			os.Exit(1)
		}
		// Test if require tables are present
		validateTables(databaseConnection, "Services")
	}
	if err != nil {
		log.Fatal("Failed to determine launch environment.")
		os.Exit(2)
	}

	// prepare and launch http server
	router := gin.Default()
	router.GET("/services", getServices)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.POST("/services")
	router.Run(":" + listenOnPort)
}
