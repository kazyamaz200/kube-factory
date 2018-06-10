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

// CreateNode is ...
func (s *Scaleway) CreateNode(name string, clusterID string) error {
	image := "image:ubuntu-mini-xenial-25g"
	arch := "x86_64"
	commercialType := "START1-XS"

	imageID, err := s.sdk.GetImageID(image, arch)
	if err != nil {
		return err
	}

	req := ScalewayServerDefinition{
		Name:           name,
		CommercialType: commercialType,
		Image:          &imageID.Identifier,
		Tags:           []string{clusterID},
	}

	serverID, err := s.sdk.PostServer(req)
	if err != nil {
		return err
	}

	server, err := s.sdk.GetServerID(serverID)
	if err != nil {
		return err
	}

	if err = s.sdk.PostServerAction(server, "poweron"); err != nil {
		return err
	}

	return nil
}
