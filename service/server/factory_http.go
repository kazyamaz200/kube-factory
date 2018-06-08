package server

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kyamazawa/kube-factory/components/factory"
)

// FactoryHTTP is ...
type FactoryHTTP struct {
	router  *mux.Router
	factory factory.FactorySDK
}

// NewFactoryHTTP is ...
func NewFactoryHTTP(opts ...FactoryHTTPOption) *FactoryHTTP {
	service := &FactoryHTTP{router: mux.NewRouter()}

	for _, opt := range opts {
		opt(service)
	}

	if service.factory == nil {
		factorySDK := factory.NewController()
		service.factory = factorySDK
	}

	service.config()

	return service
}

// FactoryHTTPOption is ...
type FactoryHTTPOption func(*FactoryHTTP)

// WithFactorySDK is ...
func WithFactorySDK(i factory.FactorySDK) FactoryHTTPOption {
	return func(s *FactoryHTTP) {
		if i != nil {
			s.factory = i
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
	ret, err := s.factory.SomeUsecase()
	println(ret)
	println(err)
}
