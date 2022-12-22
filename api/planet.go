package api

import (
	"net/http"

	"github.com/andersonribeir0/starfields/internal"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Adapter struct {
	log           *zap.Logger
	planetService internal.PlanetServiceI
}

func NewAdapter(deps *internal.Dependency) *Adapter {
	return &Adapter{
		log:           deps.Components.Log,
		planetService: deps.Services.PlanetService,
	}
}

func (adapter *Adapter) PutPlanet(c *gin.Context) {
	err := adapter.planetService.PullPlanetByID(c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)

		return
	}

	c.AbortWithStatusJSON(http.StatusOK, http.NoBody)
}
