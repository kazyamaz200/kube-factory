package provider

import (
	"net"

	"github.com/kyamazawa/glue-go/service"
)

// Awesome is ...
type Awesome struct {
	daemon AwesomeDaemon
}

// AwesomeOption is ...
type AwesomeOption func(*Awesome)

// WithAwesomeService is ...
func WithAwesomeService(i AwesomeDaemon) AwesomeOption {
	return func(s *Awesome) {
		if i != nil {
			s.daemon = i
		}
	}
}

// NewAwesome is ...
func NewAwesome(opts ...AwesomeOption) *Awesome {
	provider := &Awesome{}

	for _, opt := range opts {
		opt(provider)
	}

	if provider.daemon == nil {
		daemon := service.NewAwesomeServer()
		provider.daemon = daemon
	}

	return provider
}

// AwesomeDaemon is ...
type AwesomeDaemon interface {
	Start() net.Listener
}

// Run is ...
func (s *Awesome) Run() {
	s.daemon.Start()
}
