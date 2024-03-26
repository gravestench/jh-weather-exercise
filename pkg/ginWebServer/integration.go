package ginWebServer

import (
	"github.com/gin-gonic/gin"
)

// GinRouteInitializer represents anything that knows
// how to initialise routes using a gin route group.
type GinRouteInitializer interface {
	InitializeGinRoutes(g *gin.RouterGroup)
}

// InitializeRoutes will pass the root route group to the delegate
// which implements the GinRouteInitializer integration interface.
// Basically, just make your other service implement this and you
// can pass it into this method during app initialization.
func (s *Service) InitializeRoutes(delegate GinRouteInitializer) {
	if s.router == nil {
		s.router = gin.New()
	}

	delegate.InitializeGinRoutes(s.router.Group(""))
}
