package service

import (
	"github.com/kyamazawa/glue-go/connector"
	"github.com/kyamazawa/glue-go/model"
)

// ClusterStoreArango is ...
type ClusterStoreArango struct {
	collection connector.ArangoCollection
}

// NewClusterStoreArango is ...
func NewClusterStoreArango(opts ...ClusterStoreArangoOption) *ClusterStoreArango {
	service := &ClusterStoreArango{}
	for _, opt := range opts {
		opt(service)
	}

	if service.collection == nil {
		endpoints := []string{"http://localhost:8529"}
		dbName := "test"
		collectionName := "clusters"
		service.collection = connector.ConnectArangoCollection(endpoints, dbName, collectionName)
	}

	return service
}

// ClusterStoreArangoOption is ...
type ClusterStoreArangoOption func(*ClusterStoreArango)

// WithClusterCollection is ...
func WithClusterCollection(i connector.ArangoCollection) ClusterStoreArangoOption {
	return func(s *ClusterStoreArango) {
		if i != nil {
			s.collection = i
		}
	}
}

// Save is ...
func (s *ClusterStoreArango) Save(cluster *model.Cluster) (*model.Cluster, error) {
	meta, err := s.collection.CreateDocument(nil, cluster)
	if err != nil {
		return nil, err
	}
	cluster.ID = meta.Key
	return cluster, nil
}

// FetchByID is ...
func (s *ClusterStoreArango) FetchByID(clusterID string) (*model.Cluster, error) {
	var cluster model.Cluster
	s.collection.ReadDocument(nil, clusterID, &cluster)
	cluster.ID = clusterID
	return &cluster, nil
}
