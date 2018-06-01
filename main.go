package main

import (
	"github.com/kyamazawa/glue-go/components/awesome"
	"github.com/kyamazawa/glue-go/provider"
)

func main() {
	// full
	awesomePresenter := awesome.NewPresenter()
	awesomeInteractor := awesome.NewInteractor(awesome.WithPresenter(awesomePresenter))
	awesomeController := awesome.NewController(awesome.WithInteractor(awesomeInteractor))
	awesomeProvider1 := provider.NewAwesome(provider.WithAwesomeService(awesomeController))
	awesomeProvider1.SomeUsecase()

	// simple
	awesomeProvider2 := provider.NewAwesome()
	awesomeProvider2.SomeUsecase()
}
