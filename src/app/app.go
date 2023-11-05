package app

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ntatschner/go-go-gadget/src/domain"
	"github.com/ntatschner/go-go-gadget/src/service"
)

func Start() {
	// source, err := localOrDatabase()
	// if source == "local" {
	// 	log.Output(1, "Launching with local test data.")
	// 	listenOnPort = "8080"
	// 	sh := ServiceHandlers{service: service.NewServiceService(domain.NewServiceRepositoryStub())}
	// }
	// if source == "remote" {
	// 	log.Output(1, "Launching with database data.")
	// 	listenOnPort = "8080"
	// 	sh := ServiceHandlers{service: service.NewServiceService(domain.NewServiceRepositoryDB())}
	// }
	// if err != nil {
	// 	log.Fatal("Failed to determine launch environment.")
	// 	os.Exit(2)
	// }

	sh := ServiceHandlers{service: service.NewServiceService(domain.NewServiceRepositoryDB())}

	// prepare and launch http server
	router := gin.Default()
	router.GET("/services", sh.getAllServices)
	router.GET("/services/:service_id", sh.getService)
	// router.GET("/services/", createService)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	var listenOnPort string = "8080"
	log.Fatal(router.Run(":" + listenOnPort))
}
