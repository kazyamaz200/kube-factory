package provider

import (
	"net"

	"github.com/kyamazawa/glue-go/service"
)

// API is ...
type API struct {
	awesomeServer Daemon
}

// NewAPI is ...
func NewAPI(opts ...APIOption) *API {
	provider := &API{}

	for _, opt := range opts {
		opt(provider)
	}

	if provider.awesomeServer == nil {
		awesomeServer := service.NewAwesomeServerHTTP()
		provider.awesomeServer = awesomeServer
	}

	return provider
}

// APIOption is ...
type APIOption func(*API)

// WithAwesomeServer is ...
func WithAwesomeServer(i Daemon) APIOption {
	return func(s *API) {
		if i != nil {
			s.awesomeServer = i
		}
	}
}

// Daemon is ...
type Daemon interface {
	Start() net.Listener
}

// Run is ...
func (s *API) Run() {
	s.awesomeServer.Start()
}
