package service

import (
	"net"
	"testing"

	"github.com/kyamazawa/glue-go/components/awesome"
)

func TestNewAwesomeServer(t *testing.T) {
	t.Run("create service and it has sdk and router", func(t *testing.T) {
		// Act
		service := NewAwesomeServer()

		// Assert
		sdk := service.awesome
		router := service.router
		if sdk == nil {
			t.Errorf("got: %v\nwant: %v", sdk, "not nil")
		}
		if router == nil {
			t.Errorf("got: %v\nwant: %v", router, "not nil")
		}
	})

	t.Run("its sdk is compatible with awesome.AwesomeSDK", func(t *testing.T) {
		// Arrange
		expected := true

		// Act
		service := NewAwesomeServer()

		// Assert
		_, actual := service.awesome.(awesome.AwesomeSDK)
		if actual != expected {
			t.Errorf("got: %v\nwant: %v", actual, expected)
		}
	})

	t.Run("its sdk is injectable", func(t *testing.T) {
		// Arrange
		i1 := awesome.NewController()
		i2 := &AwesomeSDKSpy{}

		// Act
		service := NewAwesomeServer(WithAwesomeSDK(i1))

		// Assert
		actual := service.awesome
		if actual != i1 {
			t.Errorf("got: %v\nwant: %v", actual, i1)
		}
		if actual == i2 {
			t.Errorf("got: %v\nwant: %v", actual, i2)
		}
	})

	t.Run("its sdk is not be nil", func(t *testing.T) {
		// Arrange
		var i1 *awesome.Controller // nil

		// Act
		service := NewAwesomeServer(WithAwesomeSDK(i1))

		// Assert
		actual := service.awesome
		if actual == nil {
			t.Errorf("got: %v\nwant: %v", actual, "not nil")
		}
	})
}

func TestAwesomeServer_Start(t *testing.T) {
	t.Run("listen on :8080", func(t *testing.T) {
		// Arrange
		service := NewAwesomeServer()

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
		service := NewAwesomeServer()

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

type AwesomeSDKSpy struct {
	SomeUsecaseCalled       bool
	SomeUsecaseViewModelBox *awesome.SomeUsecaseViewModel
}

func (s *AwesomeSDKSpy) SomeUsecase() (*awesome.SomeUsecaseViewModel, error) {
	s.SomeUsecaseCalled = true
	return s.SomeUsecaseViewModelBox, nil
}

func TestAwesomeServer_rootHandler(t *testing.T) {
	t.Run("call SomeUsecase", func(t *testing.T) {
		// Arrange
		mockVM := &awesome.SomeUsecaseViewModel{}
		spy := &AwesomeSDKSpy{SomeUsecaseCalled: false, SomeUsecaseViewModelBox: mockVM}
		service := NewAwesomeServer(WithAwesomeSDK(spy))

		// Act
		service.rootHandler(nil, nil)

		// Assert
		called := spy.SomeUsecaseCalled
		if called != true {
			t.Errorf("got: %v\nwant: %v", called, true)
		}
	})
}
