package configManager

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Service) InitializeGinRoutes(g *gin.RouterGroup) {
	g.GET("config/paths", s.handleShowAllConfigs)
}

func (s *Service) handleShowAllConfigs(c *gin.Context) {
	type payload struct {
		Paths []string
	}

	var response payload

	for path := range s.Handles {
		response.Paths = append(response.Paths, path)
	}

	c.JSON(http.StatusOK, response)
}
