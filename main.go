package main

import (
	"log"

	"github.com/kyamazawa/glue-go/provider"
)

func main() {
	// awesomePresenter := awesome.NewPresenter()
	// awesomeInteractor := awesome.NewInteractor(awesome.WithPresenter(awesomePresenter))
	// awesomeController := awesome.NewController(awesome.WithInteractor(awesomeInteractor))
	// awesomeService := service.NewAwesomeServer(service.WithAwesomeSDK(awesomeController))
	// awesomeProvider := provider.NewAwesome(provider.WithAwesomeService(awesomeService))
	awesomeProvider := provider.NewAwesome()

	ch := make(chan bool)
	go awesomeProvider.Run()
	log.Println(<-ch)
}
