package store

import "github.com/kyamazawa/kube-factory/model"

// UserArango is ...
type UserArango struct {
	collection ArangoCollection
}

// NewUserArango is ...
func NewUserArango(opts ...UserArangoOption) *UserArango {
	service := &UserArango{}
	for _, opt := range opts {
		opt(service)
	}
	return service
}

// UserArangoOption is ...
type UserArangoOption func(*UserArango)

// WithUserCollection is ...
func WithUserCollection(i ArangoCollection) UserArangoOption {
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
