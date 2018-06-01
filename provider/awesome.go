package provider

import "github.com/kyamazawa/glue-go/components/awesome"

// AwesomeService is ...
type AwesomeService interface {
	SomeUsecase() (*awesome.SomeUsecaseViewModel, error)
}

// Awesome is ...
type Awesome struct {
	awesome AwesomeService
}

// AwesomeOption is ...
type AwesomeOption func(*Awesome)

// WithAwesomeService is ...
func WithAwesomeService(i AwesomeService) AwesomeOption {
	return func(s *Awesome) {
		if i != nil {
			s.awesome = i
		}
	}
}

// NewAwesome is ...
func NewAwesome(opts ...AwesomeOption) *Awesome {
	provider := &Awesome{}

	for _, opt := range opts {
		opt(provider)
	}

	if provider.awesome == nil {
		awesomeController := awesome.NewController()
		provider.awesome = awesomeController
	}

	return provider
}

// SomeHandler is ...
func (s *Awesome) SomeHandler() {
	ret, err := s.awesome.SomeUsecase()
	println(ret)
	println(err)
}
