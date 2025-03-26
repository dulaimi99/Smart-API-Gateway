package router

import (
	"net/http"
	"smart-api-gateway/handler"
	"smart-api-gateway/utils"

	"github.com/sirupsen/logrus"
)

func NewRouter(cfg *utils.Config, logger *logrus.Logger) http.Handler {
	logger.Info("Starting router setup")

	mux := http.NewServeMux()

	for _, route := range cfg.Routes {
		target := route.Target
		path := route.Path
		logger.Infof("Adding route: %s -> %s", path, target)
		mux.HandleFunc(path, handler.NewProxyHandler(target, logger))
	}

	logger.Info("Completed router setup")
	return mux
}
