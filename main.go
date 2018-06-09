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
	users := "users"
	clusters := "clusters"

	userCollection := connector.ConnectArangoCollection(endpoints, dbName, users)
	clusterCollection := connector.ConnectArangoCollection(endpoints, dbName, clusters)
	user := store.NewUserArango(store.WithUserCollection(userCollection))
	cluster := store.NewClusterArango(store.WithClusterCollection(clusterCollection))
	storeProvider := provider.NewStore(provider.WithUserStore(user), provider.WithClusterStore(cluster))
	presenter := factory.NewPresenter()
	interactor := factory.NewInteractor(factory.WithPresenter(presenter), factory.WithStore(storeProvider))
	controller := factory.NewController(factory.WithInteractor(interactor))

	ch := make(chan bool)
	server := server.NewFactoryHTTP(server.WithSDK(controller))
	api := provider.NewAPI(provider.WithFactoryServer(server))
	go api.Activate()
	log.Println(<-ch)
}
