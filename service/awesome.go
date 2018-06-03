package service

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kyamazawa/glue-go/components/awesome"
)

// AwesomeServer is ...
type AwesomeServer struct {
	router  *mux.Router
	awesome awesome.AwesomeSDK
}

// AwesomeServerOption is ...
type AwesomeServerOption func(*AwesomeServer)

// WithAwesomeSDK is ...
func WithAwesomeSDK(i awesome.AwesomeSDK) AwesomeServerOption {
	return func(s *AwesomeServer) {
		if i != nil {
			s.awesome = i
		}
	}
}

// NewAwesomeServer is ...
func NewAwesomeServer(opts ...AwesomeServerOption) *AwesomeServer {
	service := &AwesomeServer{router: mux.NewRouter()}

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

// Start is ...
func (s *AwesomeServer) Start() {
	log.Println("Listening on :8080")
	http.ListenAndServe(":8080", s.router)
}

func (s *AwesomeServer) config() {
	s.router.HandleFunc("/", s.rootHandler).Methods("GET")
}

func (s *AwesomeServer) rootHandler(w http.ResponseWriter, r *http.Request) {
	ret, err := s.awesome.SomeUsecase()
	println(ret)
	println(err)
}
