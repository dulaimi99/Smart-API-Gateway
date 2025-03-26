package main

import (
	"fmt"
	"log"
	"net/http"
	"smart-api-gateway/router"
	"smart-api-gateway/utils"
)

func main() {

	cfg, err := utils.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	logger := utils.NewLogger(cfg)
	logger.Info("Starting Smart API Gateway...")

	reqRouter := router.NewRouter(cfg, logger)

	logger.Infof("Server running on %s:%d", cfg.Server.Host, cfg.Server.Port)
	http.ListenAndServe(fmt.Sprintf(":%d", cfg.Server.Port), reqRouter)
}
