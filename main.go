package main

import (
	"log"

	"github.com/kyamazawa/kube-factory/components/factory"
	"github.com/kyamazawa/kube-factory/connector"
	"github.com/kyamazawa/kube-factory/provider"
	"github.com/kyamazawa/kube-factory/service/server"
	"github.com/kyamazawa/kube-factory/service/store"
)

func main() {
	endpoints := []string{"http://localhost:8529"}
	dbName := "test"

	u := store.NewUserArango(
		store.WithUserCollection(connector.ConnectArangoCollection(endpoints, dbName, "users")),
	)

	c := store.NewClusterArango(
		store.WithClusterCollection(connector.ConnectArangoCollection(endpoints, dbName, "clusters")),
	)

	s := provider.NewStore(
		provider.WithUserStore(u),
		provider.WithClusterStore(c),
	)

	p := factory.NewPresenter()

	i := factory.NewInteractor(
		factory.WithPresenter(p),
		factory.WithStore(s),
	)

	sdk := factory.NewController(factory.WithInteractor(i))
	server := server.NewFactoryHTTP(server.WithSDK(sdk))
	api := provider.NewAPI(provider.WithFactoryServer(server))

	ch := make(chan bool)
	go api.Activate()
	log.Println(<-ch)
}
