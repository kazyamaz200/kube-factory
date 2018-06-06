package main

import (
	"log"

	"github.com/kyamazawa/glue-go/entity"
	"github.com/kyamazawa/glue-go/provider"
)

func main() {
	// awesomePresenter := awesome.NewPresenter()
	// awesomeInteractor := awesome.NewInteractor(awesome.WithPresenter(awesomePresenter))
	// awesomeController := awesome.NewController(awesome.WithInteractor(awesomeInteractor))
	// awesomeService := service.NewAwesomeServer(service.WithAwesomeSDK(awesomeController))
	// awesomeProvider := provider.NewAwesome(provider.WithAwesomeService(awesomeService))
	// awesomeProvider := provider.NewAwesome()

	userProvider := provider.NewUser()

	user := &entity.User{
		Name: "NewUser",
	}
	userSaved, _ := userProvider.Save(user)
	log.Println(userSaved)

	userFetched, _ := userProvider.FetchByID(userSaved.ID)
	log.Println(userFetched)
	// ch := make(chan bool)
	// go awesomeProvider.Run()
	// log.Println(<-ch)
}
