package store

import "github.com/kyamazawa/kube-factory/model"

// ClusterArango is ...
type ClusterArango struct {
	collection ArangoCollection
}

// NewClusterArango is ...
func NewClusterArango(opts ...ClusterArangoOption) *ClusterArango {
	service := &ClusterArango{}
	for _, opt := range opts {
		opt(service)
	}
	return service
}

// ClusterArangoOption is ...
type ClusterArangoOption func(*ClusterArango)

// WithClusterCollection is ...
func WithClusterCollection(i ArangoCollection) ClusterArangoOption {
	return func(s *ClusterArango) {
		if i != nil {
			s.collection = i
		}
	}
}

// Cluster is ...
type Cluster struct {
	Key string `json:"_key,omitempty"`
}

func (s *ClusterArango) fromEntity(e *model.Cluster) *Cluster {
	ret := &Cluster{
		Key: e.ID,
	}
	return ret
}

func (s *ClusterArango) toEntity(e *Cluster) *model.Cluster {
	ret := &model.Cluster{
		ID: e.Key,
	}
	return ret
}

// Save is ...
func (s *ClusterArango) Save(entity *model.Cluster) (*model.Cluster, error) {
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
func (s *ClusterArango) Fetch(entity *model.Cluster) (*model.Cluster, error) {
	c := s.fromEntity(entity)

	s.collection.ReadDocument(nil, c.Key, &c)

	entity = s.toEntity(c)
	return entity, nil
}
