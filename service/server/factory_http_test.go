package server

import (
	"net"
	"testing"

	"github.com/kyamazawa/kube-factory/components/factory"
)

func TestNewFactoryHTTP(t *testing.T) {
	t.Run("its sdk is injectable", func(t *testing.T) {
		// Arrange
		i := &FactorySDKSpy{}

		// Act
		service := NewFactoryHTTP(WithSDK(i))

		// Assert
		_, ok := service.sdk.(factory.SDK)
		if !ok {
			t.Errorf("got: %v\nwant: %v", ok, true)
		}
	})
}

func TestFactoryHTTP_Start(t *testing.T) {
	t.Run("listen on :8080", func(t *testing.T) {
		// Arrange
		service := NewFactoryHTTP()

		// Act
		l1 := service.Start()

		// Assert
		conn, err := net.Dial("tcp", "0.0.0.0:8080")
		if err != nil {
			t.Errorf("got: %v\nwant: %v", err, "Online")
		}
		defer conn.Close()
		defer l1.Close()
	})

	t.Run("cannot bind address already in use", func(t *testing.T) {
		// Arrange
		service := NewFactoryHTTP()

		// Act
		l1 := service.Start()
		l2 := service.Start()

		// Assert
		if l2 != nil {
			t.Errorf("got: %v\nwant: %v", l2, "nil")
		}
		defer l1.Close()
	})
}

type FactorySDKSpy struct {
	CreateClusterCalled       bool
	CreateClusterViewModelBox *factory.CreateClusterViewModel
}

func (s *FactorySDKSpy) CreateCluster() (*factory.CreateClusterViewModel, error) {
	s.CreateClusterCalled = true
	return s.CreateClusterViewModelBox, nil
}

func TestFactoryHTTP_rootHandler(t *testing.T) {
	t.Run("call CreateCluster", func(t *testing.T) {
		// Arrange
		mockVM := &factory.CreateClusterViewModel{}
		spy := &FactorySDKSpy{CreateClusterCalled: false, CreateClusterViewModelBox: mockVM}
		service := NewFactoryHTTP(WithSDK(spy))

		// Act
		service.rootHandler(nil, nil)

		// Assert
		called := spy.CreateClusterCalled
		if called != true {
			t.Errorf("got: %v\nwant: %v", called, true)
		}
	})
}
