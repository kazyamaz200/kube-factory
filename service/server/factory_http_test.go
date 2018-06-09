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
	SomeUsecaseCalled       bool
	SomeUsecaseViewModelBox *factory.SomeUsecaseViewModel
}

func (s *FactorySDKSpy) SomeUsecase() (*factory.SomeUsecaseViewModel, error) {
	s.SomeUsecaseCalled = true
	return s.SomeUsecaseViewModelBox, nil
}

func TestFactoryHTTP_rootHandler(t *testing.T) {
	t.Run("call SomeUsecase", func(t *testing.T) {
		// Arrange
		mockVM := &factory.SomeUsecaseViewModel{}
		spy := &FactorySDKSpy{SomeUsecaseCalled: false, SomeUsecaseViewModelBox: mockVM}
		service := NewFactoryHTTP(WithSDK(spy))

		// Act
		service.rootHandler(nil, nil)

		// Assert
		called := spy.SomeUsecaseCalled
		if called != true {
			t.Errorf("got: %v\nwant: %v", called, true)
		}
	})
}
