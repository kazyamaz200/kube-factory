package main

import (
	"log"

	"github.com/kyamazawa/glue-go/model"
	"github.com/kyamazawa/glue-go/provider"
)

func main() {
	// factoryPresenter := factory.NewPresenter()
	// factoryInteractor := factory.NewInteractor(factory.WithPresenter(factoryPresenter))
	// factoryController := factory.NewController(factory.WithInteractor(factoryInteractor))
	// factoryService := service.NewFactoryServerHTTP(service.WithFactorySDK(factoryController))
	// factoryProvider := provider.NewAPI(provider.WithFactoryServer(factoryService))

	// factoryProvider := provider.NewAPI()

	// ch := make(chan bool)
	// go factoryProvider.Run()
	// log.Println(<-ch)

	store := provider.NewStore()

	user := &model.User{
		Name: "NewUser",
	}

	userSaved, _ := store.SaveUser(user)
	log.Println(userSaved)

	userFetched, _ := store.FetchUserByID(userSaved.ID)
	log.Println(userFetched)
}
