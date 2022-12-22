package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/andersonribeir0/starfields/api"
	"github.com/andersonribeir0/starfields/internal"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func JSONMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Next()
	}
}

func ErrorHandler(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		var errs []map[string]interface{}

		for _, ginErr := range c.Errors {
			logger.Error("whoops " + ginErr.Error())

			errs = append(errs, map[string]interface{}{
				"code":    internal.ErrorCode(ginErr),
				"message": ginErr.Error(),
			})
		}

		if errs != nil {
			c.JSON(-1, errs)
		}
	}
}

func Execute() {
	router := gin.Default()

	deps, err := internal.NewContainer(context.TODO())
	if err != nil {
		fmt.Errorf(err.Error())

		panic(err)
	}

	router.Use(ginzap.Ginzap(deps.Components.Log, time.RFC3339, true))
	router.Use(ginzap.RecoveryWithZap(deps.Components.Log, true))
	router.Use(ErrorHandler(deps.Components.Log))
	router.Use(gin.Recovery())
	router.Use(JSONMiddleware())

	apiAdapter := api.NewAdapter(deps)

	router.PUT("/planets/:id", apiAdapter.PutPlanet)

	router.Run(":3081")
}
