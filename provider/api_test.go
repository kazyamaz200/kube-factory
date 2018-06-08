package provider

import (
	"net"
	"testing"
)

func TestNewAPI(t *testing.T) {
	t.Run("create provider and it has nil factoryServer", func(t *testing.T) {
		// Act
		provider := NewAPI()

		// Assert
		factoryServer := provider.factoryServer
		if factoryServer != nil {
			t.Errorf("got: %v\nwant: %v", factoryServer, "nil")
		}
	})

	t.Run("its factoryServer is compatible with Daemon", func(t *testing.T) {
		// Arrange
		i1 := &DaemonSpy{}
		expected := true

		// Act
		provider := NewAPI(WithFactoryServer(i1))

		// Assert
		_, actual := provider.factoryServer.(Daemon)
		if actual != expected {
			t.Errorf("got: %v\nwant: %v", actual, expected)
		}
	})

	t.Run("its daemon is injectable", func(t *testing.T) {
		// Arrange
		i1 := &DaemonSpy{}
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
