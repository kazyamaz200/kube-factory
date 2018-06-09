package node

// Scaleway is ...
type Scaleway struct {
	sdk ScalewayAPI
}

// NewScaleway is ...
func NewScaleway(opts ...ScalewayOption) *Scaleway {
	service := &Scaleway{}

	for _, opt := range opts {
		opt(service)
	}

	return service
}

// ScalewayOption is ...
type ScalewayOption func(*Scaleway)

// WithSDK is ...
func WithSDK(i ScalewayAPI) ScalewayOption {
	return func(s *Scaleway) {
		if i != nil {
			s.sdk = i
		}
	}
}
