package service

import (
	driver "github.com/arangodb/go-driver"
	"github.com/kyamazawa/glue-go/entity"
	"github.com/kyamazawa/glue-go/utils"
)

// UserArangoStore is ...
type UserArangoStore struct {
	collection driver.Collection
}

// UserArangoStoreOption is ...
type UserArangoStoreOption func(*UserArangoStore)

// WithDriver is ...
func WithDriver(i driver.Collection) UserArangoStoreOption {
	return func(s *UserArangoStore) {
		if i != nil {
			s.collection = i
		}
	}
}

// NewUserArangoStore is ...
func NewUserArangoStore(opts ...UserArangoStoreOption) *UserArangoStore {
	service := &UserArangoStore{}
	for _, opt := range opts {
		opt(service)
	}

	if service.collection == nil {
		endpoints := []string{"http://localhost:8529"}
		dbName := "test"
		collectionName := "users"
		service.collection = utils.NewArangoCollection(endpoints, dbName, collectionName)
	}

	return service
}

// Save is ...
func (s *UserArangoStore) Save(user *entity.User) (*entity.User, error) {
	meta, err := s.collection.CreateDocument(nil, user)
	if err != nil {
		return nil, err
	}
	user.ID = meta.Key
	return user, nil
}

// FetchByID is ...
func (s *UserArangoStore) FetchByID(userID string) (*entity.User, error) {
	var user entity.User
	s.collection.ReadDocument(nil, userID, &user)
	user.ID = userID
	return &user, nil
}
