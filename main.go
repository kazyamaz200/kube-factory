package main

import (
	"log"

	"github.com/kyamazawa/glue-go/model"
	"github.com/kyamazawa/glue-go/provider"
)

func main() {
	// awesomePresenter := awesome.NewPresenter()
	// awesomeInteractor := awesome.NewInteractor(awesome.WithPresenter(awesomePresenter))
	// awesomeController := awesome.NewController(awesome.WithInteractor(awesomeInteractor))
	// awesomeService := service.NewAwesomeServerHTTP(service.WithAwesomeSDK(awesomeController))
	// awesomeProvider := provider.NewAPI(provider.WithAwesomeServer(awesomeService))

	// awesomeProvider := provider.NewAPI()

	// ch := make(chan bool)
	// go awesomeProvider.Run()
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
