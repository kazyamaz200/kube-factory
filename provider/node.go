package provider

// Node is ...
type Node struct {
	service NodeService
}

// NodeService is ...
type NodeService interface {
}

// NewNode is ...
func NewNode(opts ...NodeOption) *Node {
	provider := &Node{}

	for _, opt := range opts {
		opt(provider)
	}

	return provider
}

// NodeOption is ...
type NodeOption func(*Node)

// WithNodeService is ...
func WithNodeService(i NodeService) NodeOption {
	return func(s *Node) {
		if i != nil {
			s.service = i
		}
	}
}

// NodeProtocol is ...
type NodeProtocol interface {
}
