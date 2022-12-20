package api

import (
	"net/http"

	"github.com/andersonribeir0/starfields/internal"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

type Adapter struct {
	log *zap.Logger
}

func NewAdapter(deps *internal.Dependency) *Adapter {
	return &Adapter{
		log: deps.Components.Log,
	}
}

func (adapter *Adapter) GetPlanet(c *gin.Context) {
	adapter.log.Info("hello")

	c.JSON(http.StatusOK, gin.H{
		"foo": "bar",
	})
}
