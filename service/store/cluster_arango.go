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

// Save is ...
func (s *ClusterArango) Save(cluster *model.Cluster) (*model.Cluster, error) {
	meta, err := s.collection.CreateDocument(nil, cluster)
	if err != nil {
		return nil, err
	}
	cluster.ID = meta.Key
	return cluster, nil
}

// FetchByID is ...
func (s *ClusterArango) FetchByID(clusterID string) (*model.Cluster, error) {
	var cluster model.Cluster
	s.collection.ReadDocument(nil, clusterID, &cluster)
	cluster.ID = clusterID
	return &cluster, nil
}
