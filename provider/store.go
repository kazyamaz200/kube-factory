package provider

import (
	"github.com/kyamazawa/glue-go/model"
	"github.com/kyamazawa/glue-go/service"
)

// Store is ...
type Store struct {
	userStore UserStore
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

// UserStore is ...
type UserStore interface {
	Save(user *model.User) (*model.User, error)
	FetchByID(userID string) (*model.User, error)
}

// SaveUser is ...
func (s *Store) SaveUser(user *model.User) (*model.User, error) {
	return s.userStore.Save(user)
}

// FetchUserByID is ...
func (s *Store) FetchUserByID(userID string) (*model.User, error) {
	return s.userStore.FetchByID(userID)
}
