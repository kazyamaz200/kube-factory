package service

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kyamazawa/glue-go/components/awesome"
)

// AwesomeServerHTTP is ...
type AwesomeServerHTTP struct {
	router  *mux.Router
	awesome awesome.AwesomeSDK
}

// NewAwesomeServerHTTP is ...
func NewAwesomeServerHTTP(opts ...AwesomeServerHTTPOption) *AwesomeServerHTTP {
	service := &AwesomeServerHTTP{router: mux.NewRouter()}

	for _, opt := range opts {
		opt(service)
	}

	if service.awesome == nil {
		awesomeSDK := awesome.NewController()
		service.awesome = awesomeSDK
	}

	service.config()

	return service
}

// AwesomeServerHTTPOption is ...
type AwesomeServerHTTPOption func(*AwesomeServerHTTP)

// WithAwesomeSDK is ...
func WithAwesomeSDK(i awesome.AwesomeSDK) AwesomeServerHTTPOption {
	return func(s *AwesomeServerHTTP) {
		if i != nil {
			s.awesome = i
		}
	}
}

// Start is ...
func (s *AwesomeServerHTTP) Start() net.Listener {
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

func (s *AwesomeServerHTTP) config() {
	s.router.HandleFunc("/", s.rootHandler).Methods("GET")
}

func (s *AwesomeServerHTTP) rootHandler(w http.ResponseWriter, r *http.Request) {
	ret, err := s.awesome.SomeUsecase()
	println(ret)
	println(err)
}
