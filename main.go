package main

import (
	"log"

	"github.com/kyamazawa/kube-factory/provider"
	"github.com/kyamazawa/kube-factory/service/server"
)

func main() {
	ch := make(chan bool)
	server := server.NewFactoryHTTP()
	api := provider.NewAPI(provider.WithFactoryServer(server))
	go api.Activate()
	log.Println(<-ch)
}
