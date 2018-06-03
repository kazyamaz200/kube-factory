package main

import (
	"log"

	"github.com/kyamazawa/glue-go/components/awesome"
	"github.com/kyamazawa/glue-go/provider"
	"github.com/kyamazawa/glue-go/service"
)

func main() {
	ch := make(chan bool)
	awesomePresenter := awesome.NewPresenter()
	awesomeInteractor := awesome.NewInteractor(awesome.WithPresenter(awesomePresenter))
	awesomeController := awesome.NewController(awesome.WithInteractor(awesomeInteractor))
	awesomeService := service.NewAwesomeServer(service.WithAwesomeSDK(awesomeController))
	awesomeProvider := provider.NewAwesome(provider.WithAwesomeService(awesomeService))
	go awesomeProvider.Run()
	log.Println(<-ch)
}
