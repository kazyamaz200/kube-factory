package service

import (
	"github.com/kyamazawa/glue-go/connector"
	"github.com/kyamazawa/glue-go/model"
)

// UserStoreArango is ...
type UserStoreArango struct {
	collection connector.ArangoCollection
}

// NewUserStoreArango is ...
func NewUserStoreArango(opts ...UserStoreArangoOption) *UserStoreArango {
	service := &UserStoreArango{}
	for _, opt := range opts {
		opt(service)
	}

	if service.collection == nil {
		endpoints := []string{"http://localhost:8529"}
		dbName := "test"
		collectionName := "users"
		service.collection = connector.ConnectArangoCollection(endpoints, dbName, collectionName)
	}

	return service
}

// UserStoreArangoOption is ...
type UserStoreArangoOption func(*UserStoreArango)

// WithUserCollection is ...
func WithUserCollection(i connector.ArangoCollection) UserStoreArangoOption {
	return func(s *UserStoreArango) {
		if i != nil {
			s.collection = i
		}
	}
}

// Save is ...
func (s *UserStoreArango) Save(user *model.User) (*model.User, error) {
	meta, err := s.collection.CreateDocument(nil, user)
	if err != nil {
		return nil, err
	}
	user.ID = meta.Key
	return user, nil
}

// FetchByID is ...
func (s *UserStoreArango) FetchByID(userID string) (*model.User, error) {
	var user model.User
	s.collection.ReadDocument(nil, userID, &user)
	user.ID = userID
	return &user, nil
}
