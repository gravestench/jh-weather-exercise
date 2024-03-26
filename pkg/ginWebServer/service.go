package ginWebServer

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Service is a combined http server with gin web router.
//
// This module provides an integration interface for other
// services to implement to allow them to be the delegate for
// initializing their routes using the root route group.
type Service struct {
	Config
	server *http.Server
	router *gin.Engine
}

// Serve is a blocking call that will init the web server
func (s *Service) Serve() error {
	s.server = &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", s.Port),
		Handler: s.router,
	}

	return s.server.ListenAndServe()
}
