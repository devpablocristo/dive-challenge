package api

import (
	"log"

	"github.com/gin-gonic/gin"

	hdl "github.com/devpablocristo/dive-challenge/cmd/api/handlers"
	ucs "github.com/devpablocristo/dive-challenge/internal/core"
)

func Build(dep *Dependencies) *gin.Engine {
	usecase := ucs.NewUseCase(dep.Repository, dep.ApiClient)
	handler := hdl.NewRestHandler(usecase)

	r := gin.Default()
	r.GET("/api/v1/hello", handler.HelloWorld)
	r.GET("/api/v1/ltp", handler.GetLTP)

	log.Println("Running server on port " + dep.RouterPort + "...")
	if err := r.Run(":" + dep.RouterPort); err != nil {
		log.Fatalf("Failed to run server: %s", err)
	}

	return r
}
