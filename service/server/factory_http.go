package server

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kyamazawa/kube-factory/components/factory"
	"github.com/kyamazawa/kube-factory/connector"
	"github.com/kyamazawa/kube-factory/provider"
	"github.com/kyamazawa/kube-factory/service/store"
)

// FactoryHTTP is ...
type FactoryHTTP struct {
	router *mux.Router
	sdk    factory.SDK
}

// NewFactoryHTTP is ...
func NewFactoryHTTP(opts ...FactoryHTTPOption) *FactoryHTTP {
	service := &FactoryHTTP{router: mux.NewRouter()}

	for _, opt := range opts {
		opt(service)
	}

	if service.sdk == nil {
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

		service.sdk = controller
	}

	service.config()

	return service
}

// FactoryHTTPOption is ...
type FactoryHTTPOption func(*FactoryHTTP)

// WithSDK is ...
func WithSDK(i factory.SDK) FactoryHTTPOption {
	return func(s *FactoryHTTP) {
		if i != nil {
			s.sdk = i
		}
	}
}

// Start is ...
func (s *FactoryHTTP) Start() net.Listener {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", 8080))
	if err != nil {
		log.Println(err)
		return nil
	}
	go func() {
		http.Serve(l, s.router)
	}()
	// log.Println("Listening on :8080")
	return l
}

func (s *FactoryHTTP) config() {
	s.router.HandleFunc("/", s.rootHandler).Methods("GET")
}

func (s *FactoryHTTP) rootHandler(w http.ResponseWriter, r *http.Request) {
	ret, err := s.sdk.SomeUsecase()
	println(ret)
	println(err)
}
