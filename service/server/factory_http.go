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
	api := s.router.PathPrefix("/api").Subrouter()
	api.Use(jsonResponseMiddleware)
	api.HandleFunc("/clusters", s.ClustersPost).Methods("POST")
	api.HandleFunc("/clusters", s.ClustersGet).Methods("GET")
	api.HandleFunc("/clusters/{id}", s.ClustersIdGet).Methods("GET")
	api.HandleFunc("/clusters/{id}", s.ClustersIdPut).Methods("PUT")
	api.HandleFunc("/clusters/{id}", s.ClustersIdDelete).Methods("DELETE")
}

func jsonResponseMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		next.ServeHTTP(w, r)
	})
}

// ClustersPost is handle "POST /api/clusters"
func (s *FactoryHTTP) ClustersPost(w http.ResponseWriter, r *http.Request) {
	result, err := s.sdk.CreateCluster()
	var body []byte
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		body, _ = json.Marshal(err)
	} else {
		w.WriteHeader(http.StatusAccepted)
		body, _ = json.Marshal(result)
	}
	w.Write(body)
}

// ClustersGet is handle "GET /api/clusters"
func (s *FactoryHTTP) ClustersGet(w http.ResponseWriter, r *http.Request) {
	result, err := s.sdk.ListCluster()
	var body []byte
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		body, _ = json.Marshal(err)
	} else {
		w.WriteHeader(http.StatusAccepted)
		body, _ = json.Marshal(result)
	}
	w.Write(body)
}

// ClustersIdGet is handle "GET /api/clusters/{id}"
func (s *FactoryHTTP) ClustersIdGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	result, err := s.sdk.DescribeCluster(vars["id"])
	var body []byte
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		body, _ = json.Marshal(err)
	} else {
		w.WriteHeader(http.StatusAccepted)
		body, _ = json.Marshal(result)
	}
	w.Write(body)
}

// ClustersIdPut is handle "PUT /api/clusters/{id}"
func (s *FactoryHTTP) ClustersIdPut(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	result, err := s.sdk.UpdateCluster(vars["id"])
	var body []byte
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		body, _ = json.Marshal(err)
	} else {
		w.WriteHeader(http.StatusAccepted)
		body, _ = json.Marshal(result)
	}
	w.Write(body)
}

// ClustersIdDelete is handle "DELETE /api/clusters/{id}"
func (s *FactoryHTTP) ClustersIdDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	result, err := s.sdk.DeleteCluster(vars["id"])
	var body []byte
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		body, _ = json.Marshal(err)
	} else {
		w.WriteHeader(http.StatusAccepted)
		body, _ = json.Marshal(result)
	}
	w.Write(body)
}
