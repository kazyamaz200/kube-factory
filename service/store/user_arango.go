package store

import (
	"github.com/kyamazawa/kube-factory/connector"
	"github.com/kyamazawa/kube-factory/model"
)

// UserArango is ...
type UserArango struct {
	collection connector.ArangoCollection
}

// NewUserArango is ...
func NewUserArango(opts ...UserArangoOption) *UserArango {
	service := &UserArango{}
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

// UserArangoOption is ...
type UserArangoOption func(*UserArango)

// WithUserCollection is ...
func WithUserCollection(i connector.ArangoCollection) UserArangoOption {
	return func(s *UserArango) {
		if i != nil {
			s.collection = i
		}
	}
}

// Save is ...
func (s *UserArango) Save(user *model.User) (*model.User, error) {
	meta, err := s.collection.CreateDocument(nil, user)
	if err != nil {
		return nil, err
	}
	user.ID = meta.Key
	return user, nil
}

// FetchByID is ...
func (s *UserArango) FetchByID(userID string) (*model.User, error) {
	var user model.User
	s.collection.ReadDocument(nil, userID, &user)
	user.ID = userID
	return &user, nil
}
