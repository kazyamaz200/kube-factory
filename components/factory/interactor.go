package factory

import (
	"github.com/kyamazawa/kube-factory/model"
	"github.com/kyamazawa/kube-factory/provider"
)

// Interactor is ...
type Interactor struct {
	presenter Presentation
	store     provider.StoreProtocol
	node      provider.NodeProtocol
	dns       provider.DNSProtocol
}

// NewInteractor is ...
func NewInteractor(opts ...InteractorOption) *Interactor {
	interactor := &Interactor{}
	for _, opt := range opts {
		opt(interactor)
	}
	return interactor
}

// InteractorOption is ...
type InteractorOption func(*Interactor)

// WithPresenter is ...
func WithPresenter(p Presentation) InteractorOption {
	return func(s *Interactor) {
		if p != nil {
			s.presenter = p
		}
	}
}

// WithStore is ...
func WithStore(p provider.StoreProtocol) InteractorOption {
	return func(s *Interactor) {
		if p != nil {
			s.store = p
		}
	}
}

// WithNode is ...
func WithNode(p provider.NodeProtocol) InteractorOption {
	return func(s *Interactor) {
		if p != nil {
			s.node = p
		}
	}
}

// WithDNS is ...
func WithDNS(p provider.DNSProtocol) InteractorOption {
	return func(s *Interactor) {
		if p != nil {
			s.dns = p
		}
	}
}

// Interaction is ...
type Interaction interface {
	DoCreateCluster(req *CreateClusterRequest, callback func(*CreateClusterViewModel, error))
}

// DoCreateCluster is ...
func (s *Interactor) DoCreateCluster(req *CreateClusterRequest, callback func(*CreateClusterViewModel, error)) {
	cluster := &model.Cluster{}
	cluster, err := s.store.SaveCluster(cluster)
	if err != nil {
		callback(nil, err)
	}
	res := &CreateClusterResponse{ID: cluster.ID}
	s.presenter.PresentCreateCluster(res, callback)
}
