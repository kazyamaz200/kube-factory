package main

import (
	"log"

	"github.com/kyamazawa/glue-go/components/factory"
	"github.com/kyamazawa/glue-go/provider"
	"github.com/kyamazawa/glue-go/service"
)

func main() {
	storeProvider := provider.NewStore()
	factoryPresenter := factory.NewPresenter()
	factoryInteractor := factory.NewInteractor(
		factory.WithPresenter(factoryPresenter),
		factory.WithStore(storeProvider),
	)
	factoryController := factory.NewController(
		factory.WithInteractor(factoryInteractor),
	)
	factoryService := service.NewFactoryServerHTTP(
		service.WithFactorySDK(factoryController),
	)
	factoryProvider := provider.NewAPI(
		provider.WithFactoryServer(factoryService),
	)

	// factoryProvider := provider.NewAPI()

	ch := make(chan bool)
	go factoryProvider.Run()
	log.Println(<-ch)
}
