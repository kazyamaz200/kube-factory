package provider

import (
	"github.com/kyamazawa/glue-go/entity"
	"github.com/kyamazawa/glue-go/service"
)

// User is ...
type User struct {
	store UserStore
}

// UserOption is ...
type UserOption func(*User)

// WithUserService is ...
func WithUserService(i UserStore) UserOption {
	return func(s *User) {
		if i != nil {
			s.store = i
		}
	}
}

// NewUser is ...
func NewUser(opts ...UserOption) *User {
	provider := &User{}

	for _, opt := range opts {
		opt(provider)
	}

	if provider.store == nil {
		store := service.NewUserArangoStore()
		provider.store = store
	}

	return provider
}

// UserStore is ...
type UserStore interface {
	Save(user *entity.User) (*entity.User, error)
	FetchByID(userID string) (*entity.User, error)
}

// Save is ...
func (s *User) Save(user *entity.User) (*entity.User, error) {
	return s.store.Save(user)
}

func (s *User) FetchByID(userID string) (*entity.User, error) {
	return s.store.FetchByID(userID)
}
