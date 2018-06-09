package dns

// Cloudflare is ...
type Cloudflare struct {
	sdk CloudflareAPI
}

// NewCloudflare is ...
func NewCloudflare(opts ...CloudflareOption) *Cloudflare {
	service := &Cloudflare{}

	for _, opt := range opts {
		opt(service)
	}

	return service
}

// CloudflareOption is ...
type CloudflareOption func(*Cloudflare)

// WithSDK is ...
func WithSDK(i CloudflareAPI) CloudflareOption {
	return func(s *Cloudflare) {
		if i != nil {
			s.sdk = i
		}
	}
}
