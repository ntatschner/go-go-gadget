package app

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ntatschner/gogogadget/src/domain"
	"github.com/ntatschner/gogogadget/src/service"
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

	sh := ServiceHandlers{service: service.NewServiceService(domain.NewServiceRepositoryStub())}

	// prepare and launch http server
	router := gin.Default()
	router.GET("/services", sh.getAllServices)
	// router.GET("/services/{service_id:[0-9]{4}}", getService)
	// router.GET("/services/", createService)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	log.Fatal(router.Run(":" + listenOnPort))
}
