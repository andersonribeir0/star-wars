package cmd

import (
	"context"

	"github.com/andersonribeir0/starfields/api"
	"github.com/andersonribeir0/starfields/internal"
	"github.com/gin-gonic/gin"
)

func Execute() {
	router := gin.Default()

	deps, err := internal.NewContainer(context.TODO())
	if err != nil {
		panic(err)
	}

	apiAdapter := api.NewAdapter(deps)

	router.GET("/v1/planet", apiAdapter.GetPlanet)

	router.Run()
}
