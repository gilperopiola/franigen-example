package app

import "franigen-example/api/infrastructure/dependencies"

func Start() {
	restHandlers := dependencies.SetUp()
	startRestServer(restHandlers)
}
