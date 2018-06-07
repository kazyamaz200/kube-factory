package provider

import (
	"github.com/kyamazawa/kube-factory/model"
	"github.com/kyamazawa/kube-factory/service"
)

// Store is ...
type Store struct {
	userStore    UserStore
	clusterStore ClusterStore
}

// NewStore is ...
func NewStore(opts ...StoreOption) *Store {
	provider := &Store{}

	for _, opt := range opts {
		opt(provider)
	}

	if provider.userStore == nil {
		userStore := service.NewUserStoreArango()
		provider.userStore = userStore
	}

	if provider.clusterStore == nil {
		clusterStore := service.NewClusterStoreArango()
		provider.clusterStore = clusterStore
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

// UserStore is ...
type UserStore interface {
	Save(user *model.User) (*model.User, error)
	FetchByID(userID string) (*model.User, error)
}

// ClusterStore is ...
type ClusterStore interface {
	Save(cluster *model.Cluster) (*model.Cluster, error)
	FetchByID(clusterID string) (*model.Cluster, error)
}

// SaveUser is ...
func (s *Store) SaveUser(user *model.User) (*model.User, error) {
	return s.userStore.Save(user)
}

// FetchUserByID is ...
func (s *Store) FetchUserByID(userID string) (*model.User, error) {
	return s.userStore.FetchByID(userID)
}

// SaveCluster is ...
func (s *Store) SaveCluster(cluster *model.Cluster) (*model.Cluster, error) {
	return s.clusterStore.Save(cluster)
}

// FetchClusterByID is ...
func (s *Store) FetchClusterByID(clusterID string) (*model.Cluster, error) {
	return s.clusterStore.FetchByID(clusterID)
}
