package provider

import (
	"net"
	"testing"

	"github.com/kyamazawa/glue-go/service"
)

func TestNewAwesome(t *testing.T) {
	t.Run("create provider and it has daemon", func(t *testing.T) {
		// Act
		provider := NewAwesome()

		// Assert
		daemon := provider.daemon
		if daemon == nil {
			t.Errorf("got: %v\nwant: %v", daemon, "not nil")
		}
	})

	t.Run("its daemon is compatible with AwesomeDaemon", func(t *testing.T) {
		// Arrange
		expected := true

		// Act
		provider := NewAwesome()

		// Assert
		_, actual := provider.daemon.(AwesomeDaemon)
		if actual != expected {
			t.Errorf("got: %v\nwant: %v", actual, expected)
		}
	})

	t.Run("its daemon is injectable", func(t *testing.T) {
		// Arrange
		i1 := service.NewAwesomeServer()
		i2 := &AwesomeDaemonSpy{}

		// Act
		provider := NewAwesome(WithAwesomeService(i1))

		// Assert
		actual := provider.daemon
		if actual != i1 {
			t.Errorf("got: %v\nwant: %v", actual, i1)
		}
		if actual == i2 {
			t.Errorf("got: %v\nwant: %v", actual, i2)
		}
	})

	t.Run("its daemon is not be nil", func(t *testing.T) {
		// Arrange
		var i1 *service.AwesomeServer // nil

		// Act
		provider := NewAwesome(WithAwesomeService(i1))

		// Assert
		actual := provider.daemon
		if actual == nil {
			t.Errorf("got: %v\nwant: %v", actual, "not nil")
		}
	})
}

func TestAwesome_Run(t *testing.T) {
	t.Run("call Start", func(t *testing.T) {
		// Arrange
		spy := &AwesomeDaemonSpy{StartCalled: false}
		provider := NewAwesome(WithAwesomeService(spy))

		// Act
		provider.Run()

		// Assert
		called := spy.StartCalled
		if called != true {
			t.Errorf("got: %v\nwant: %v", called, true)
		}
	})
}

type AwesomeDaemonSpy struct {
	StartCalled bool
}

func (s *AwesomeDaemonSpy) Start() net.Listener {
	s.StartCalled = true
	return nil
}
