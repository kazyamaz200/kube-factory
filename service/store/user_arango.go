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

// User is ...
type User struct {
	Key string `json:"_key,omitempty"`
}

func (s *UserArango) fromEntity(e *model.User) *User {
	ret := &User{
		Key: e.ID,
	}
	return ret
}

func (s *UserArango) toEntity(e *User) *model.User {
	ret := &model.User{
		ID: e.Key,
	}
	return ret
}

// Save is ...
func (s *UserArango) Save(entity *model.User) (*model.User, error) {
	c := s.fromEntity(entity)

	meta, err := s.collection.CreateDocument(nil, c)
	if err != nil {
		return nil, err
	}
	c.Key = meta.Key

	entity = s.toEntity(c)
	return entity, nil
}

// Fetch is ...
func (s *UserArango) Fetch(entity *model.User) (*model.User, error) {
	c := s.fromEntity(entity)

	s.collection.ReadDocument(nil, c.Key, &c)

	entity = s.toEntity(c)
	return entity, nil
}
