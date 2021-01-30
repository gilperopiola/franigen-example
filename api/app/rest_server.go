package app

import (
	"franigen-example/api/app/rest"
	"franigen-example/api/infrastructure/dependencies"
	"franigen-example/config"
)

func startRestServer(handlers *dependencies.RestHandlerContainer) {
	router := rest.CreateRouter()
	rest.Map(router, handlers)

	if err := router.Run(":" + config.Get().Port); err != nil {
		panic(err)
	}
}
