package service

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kyamazawa/glue-go/components/factory"
)

// FactoryServerHTTP is ...
type FactoryServerHTTP struct {
	router  *mux.Router
	factory factory.FactorySDK
}

// NewFactoryServerHTTP is ...
func NewFactoryServerHTTP(opts ...FactoryServerHTTPOption) *FactoryServerHTTP {
	service := &FactoryServerHTTP{router: mux.NewRouter()}

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

// FactoryServerHTTPOption is ...
type FactoryServerHTTPOption func(*FactoryServerHTTP)

// WithFactorySDK is ...
func WithFactorySDK(i factory.FactorySDK) FactoryServerHTTPOption {
	return func(s *FactoryServerHTTP) {
		if i != nil {
			s.factory = i
		}
	}
}

// Start is ...
func (s *FactoryServerHTTP) Start() net.Listener {
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

func (s *FactoryServerHTTP) config() {
	s.router.HandleFunc("/", s.rootHandler).Methods("GET")
}

func (s *FactoryServerHTTP) rootHandler(w http.ResponseWriter, r *http.Request) {
	ret, err := s.factory.SomeUsecase()
	println(ret)
	println(err)
}
