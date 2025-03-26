package handler

import (
	"io"
	"net/http"
	"net/url"

	"github.com/sirupsen/logrus"
)

func NewProxyHandler(target string, logger *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Infof("Handling request: %s %s", r.Method, r.RequestURI)

		targetURL, err := url.Parse(target)
		if err != nil {
			logger.Errorf("Error parsing target URL: %v", err)
			http.Error(w, "Invalid target URL", http.StatusInternalServerError)
			return
		}

		proxyReq, err := http.NewRequest(r.Method, targetURL.String()+r.RequestURI, r.Body)
		logger.Infof("Forwarding request to: %s", targetURL.String()+r.RequestURI)
		if err != nil {
			logger.Errorf("Error creating request: %v", err)
			http.Error(w, "Failed to create request", http.StatusInternalServerError)
			return
		}

		proxyReq.Header = r.Header

		resp, err := http.DefaultClient.Do(proxyReq)
		if err != nil {
			logger.Errorf("Error forwarding request: %v", err)
			http.Error(w, "Failed to forward request", http.StatusBadGateway)
			return
		}
		defer resp.Body.Close()

		for k, v := range resp.Header {
			for _, vv := range v {
				w.Header().Add(k, vv)
			}
		}
		w.WriteHeader(resp.StatusCode)
		io.Copy(w, resp.Body)

		logger.Infof("Response status: %d", resp.StatusCode)
		logger.Infof("Completed handling request: %s %s", r.Method, r.RequestURI)
	}
}
