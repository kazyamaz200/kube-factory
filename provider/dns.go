package provider

// DNS is ...
type DNS struct {
	service DNSService
}

// DNSService is ...
type DNSService interface {
}

// NewDNS is ...
func NewDNS(opts ...DNSOption) *DNS {
	provider := &DNS{}

	for _, opt := range opts {
		opt(provider)
	}

	return provider
}

// DNSOption is ...
type DNSOption func(*DNS)

// WithDNSService is ...
func WithDNSService(i DNSService) DNSOption {
	return func(s *DNS) {
		if i != nil {
			s.service = i
		}
	}
}

// DNSProtocol is ...
type DNSProtocol interface {
}
