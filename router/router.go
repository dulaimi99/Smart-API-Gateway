package router

import (
	"net/http"
	"smart-api-gateway/handler"
	"smart-api-gateway/middleware"
	"smart-api-gateway/utils"

	"github.com/sirupsen/logrus"
)

func NewRouter(cfg *utils.Config, logger *logrus.Logger) http.Handler {
	logger.Info("Starting router setup")

	mux := http.NewServeMux()
	rateLimiter := middleware.NewRateLimiter(logger)

	for _, route := range cfg.Routes {
		target := route.Target
		path := route.Path
		capacity := route.Capacity
		refillRate := route.RefillRate

		logger.Infof("Adding route: %s -> %s (Capacity: %d, RefillRate: %d)", path, target, capacity, refillRate)

		proxyHandler := handler.NewProxyHandler(target, logger)

		// Wrap the proxy handler with the rate-limiting middleware
		mux.Handle(path, rateLimiter.Middleware(proxyHandler, capacity, refillRate))
	}

	logger.Info("Completed router setup")
	return mux
}
