package provider

import (
	"github.com/kyamazawa/kube-factory/model"
)

// Store is ...
type Store struct {
	userStore    UserStore
	clusterStore ClusterStore
}

// UserStore is ...
type UserStore interface {
	Save(entity *model.User) (*model.User, error)
	Fetch(entity *model.User) (*model.User, error)
}

// ClusterStore is ...
type ClusterStore interface {
	Save(entity *model.Cluster) (*model.Cluster, error)
	Fetch(entity *model.Cluster) (*model.Cluster, error)
}

// NewStore is ...
func NewStore(opts ...StoreOption) *Store {
	provider := &Store{}
	for _, opt := range opts {
		opt(provider)
	}
	return provider
}

// StoreOption is ...
type StoreOption func(*Store)

// WithUserStore is ...
func WithUserStore(i UserStore) StoreOption {
	return func(s *Store) {
		if i != nil {
			s.userStore = i
		}
	}
}

// WithClusterStore is ...
func WithClusterStore(i ClusterStore) StoreOption {
	return func(s *Store) {
		if i != nil {
			s.clusterStore = i
		}
	}
}

// StoreProtocol is ...
type StoreProtocol interface {
	SaveUser(entity *model.User) (*model.User, error)
	FetchUser(entity *model.User) (*model.User, error)
	SaveCluster(entity *model.Cluster) (*model.Cluster, error)
	FetchCluster(entity *model.Cluster) (*model.Cluster, error)
}

// SaveUser is ...
func (s *Store) SaveUser(entity *model.User) (*model.User, error) {
	return s.userStore.Save(entity)
}

// FetchUser is ...
func (s *Store) FetchUser(entity *model.User) (*model.User, error) {
	return s.userStore.Fetch(entity)
}

// SaveCluster is ...
func (s *Store) SaveCluster(entity *model.Cluster) (*model.Cluster, error) {
	return s.clusterStore.Save(entity)
}

// FetchCluster is ...
func (s *Store) FetchCluster(entity *model.Cluster) (*model.Cluster, error) {
	return s.clusterStore.Fetch(entity)
}
