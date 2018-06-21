package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kyamazawa/kube-factory/components/factory"
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
	s.router.HandleFunc("/api/clusters", s.postClusters).Methods("POST")
}

func (s *FactoryHTTP) postClusters(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	result, err := s.sdk.CreateCluster()

	var body []byte
	if err != nil {
		w.WriteHeader(400)
		body, _ = json.Marshal(err)
	} else {
		body, _ = json.Marshal(result)
	}

	w.Write(body)
}
