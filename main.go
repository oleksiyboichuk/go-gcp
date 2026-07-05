package main

import (
	"fmt"
	"go-gcp/config"
	"go-gcp/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()
	router := gin.Default()
	routes.SetupRoutes(router)

	fmt.Printf("Starting server on port %s", cfg.Port)
	router.Run(cfg.Port)
}
