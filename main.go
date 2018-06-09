package main

import (
	"log"

	"github.com/kyamazawa/kube-factory/components/factory"
	"github.com/kyamazawa/kube-factory/provider"
	"github.com/kyamazawa/kube-factory/service/dns"
	"github.com/kyamazawa/kube-factory/service/node"
	"github.com/kyamazawa/kube-factory/service/server"
	"github.com/kyamazawa/kube-factory/service/store"
)

func main() {
	ch := make(chan bool)
	component := setup()
	server := server.NewFactoryHTTP(server.WithSDK(component))
	daemon := provider.NewAPI(provider.WithFactoryServer(server))
	go daemon.Activate()
	log.Println(<-ch)
}

func setup() factory.SDK {
	// arango
	endpoints := []string{"http://localhost:8529"}
	dbName := "test"
	userName := "testuser"
	password := "testpass"
	usersCol := "users"
	clustersCol := "clusters"
	// scaleway
	org := ""
	token := ""
	ug := ""
	region := ""
	// cloudflare
	key := ""
	email := ""

	// services
	users := store.NewUserArango(
		store.WithUserCollection(store.ConnectArangoCollection(endpoints, dbName, userName, password, usersCol)),
	)
	clusters := store.NewClusterArango(
		store.WithClusterCollection(store.ConnectArangoCollection(endpoints, dbName, userName, password, clustersCol)),
	)
	scaleway := node.NewScaleway(
		node.WithSDK(node.ScalewayClient(org, token, ug, region)),
	)
	cloudflare := dns.NewCloudflare(
		dns.WithSDK(dns.CloudflareClient(key, email)),
	)

	// providers
	storeProvider := provider.NewStore(
		provider.WithUserStore(users),
		provider.WithClusterStore(clusters),
	)
	nodeProvider := provider.NewNode(
		provider.WithNodeService(scaleway),
	)
	dnsProvider := provider.NewDNS(
		provider.WithDNSService(cloudflare),
	)

	// component
	factory := factory.NewController(
		factory.WithInteractor(factory.NewInteractor(
			factory.WithPresenter(factory.NewPresenter()),
			factory.WithStore(storeProvider),
			factory.WithNode(nodeProvider),
			factory.WithDNS(dnsProvider),
		)),
	)

	return factory
}
