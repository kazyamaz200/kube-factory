package provider

import (
	"net"
	"testing"

	"github.com/kyamazawa/kube-factory/service"
)

func TestNewAPI(t *testing.T) {
	t.Run("create provider and it has daemon", func(t *testing.T) {
		// Act
		provider := NewAPI()

		// Assert
		daemon := provider.factoryServer
		if daemon == nil {
			t.Errorf("got: %v\nwant: %v", daemon, "not nil")
		}
	})

	t.Run("its daemon is compatible with Daemon", func(t *testing.T) {
		// Arrange
		expected := true

		// Act
		provider := NewAPI()

		// Assert
		_, actual := provider.factoryServer.(Daemon)
		if actual != expected {
			t.Errorf("got: %v\nwant: %v", actual, expected)
		}
	})

	t.Run("its daemon is injectable", func(t *testing.T) {
		// Arrange
		i1 := service.NewFactoryServerHTTP()
		i2 := &DaemonSpy{}

		// Act
		provider := NewAPI(WithFactoryServer(i1))

		// Assert
		actual := provider.factoryServer
		if actual != i1 {
			t.Errorf("got: %v\nwant: %v", actual, i1)
		}
		if actual == i2 {
			t.Errorf("got: %v\nwant: %v", actual, i2)
		}
	})

	t.Run("its daemon is not be nil", func(t *testing.T) {
		// Act
		provider := NewAPI(WithFactoryServer(nil))

		// Assert
		actual := provider.factoryServer
		if actual == nil {
			t.Errorf("got: %v\nwant: %v", actual, "not nil")
		}
	})
}

type DaemonSpy struct {
	StartCalled bool
}

func (s *DaemonSpy) Start() net.Listener {
	s.StartCalled = true
	return nil
}

func TestAPI_Run(t *testing.T) {
	t.Run("call Start", func(t *testing.T) {
		// Arrange
		spy := &DaemonSpy{StartCalled: false}
		provider := NewAPI(WithFactoryServer(spy))

		// Act
		provider.Activate()

		// Assert
		called := spy.StartCalled
		if called != true {
			t.Errorf("got: %v\nwant: %v", called, true)
		}
	})
}
