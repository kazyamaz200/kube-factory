package factory

import "testing"

func TestNewController(t *testing.T) {
	t.Run("its interactor is injectable", func(t *testing.T) {
		// Arrange
		i := &InteractorSpy{}

		// Act
		controller := NewController(WithInteractor(i))

		// Assert
		_, ok := controller.interactor.(Interaction)
		if !ok {
			t.Errorf("got: %v\nwant: %v", ok, true)
		}
	})
}

type InteractorSpy struct {
	DoCreateClusterCalled     bool
	CreateClusterViewModelBox *CreateClusterViewModel
	CreateClusterRequestBox   *CreateClusterRequest
}

func (s *InteractorSpy) DoCreateCluster(req *CreateClusterRequest, callback func(*CreateClusterViewModel, error)) {
	s.DoCreateClusterCalled = true
	s.CreateClusterRequestBox = req
	callback(s.CreateClusterViewModelBox, nil)
}

func TestController_CreateCluster(t *testing.T) {
	t.Run("return CreateClusterViewModel and error", func(t *testing.T) {
		// Arrange
		mockVM := &CreateClusterViewModel{}
		spy := &InteractorSpy{CreateClusterViewModelBox: mockVM}
		controller := NewController(WithInteractor(spy))

		// Act
		actual, err := controller.CreateCluster()

		// Assert
		if actual != mockVM {
			t.Errorf("got: %v\nwant: %v", actual, mockVM)
		}
		if err != nil {
			t.Errorf("got: %v\nwant: %v", err, nil)
		}
	})

	t.Run("call DoCreateCluster with CreateClusterRequest", func(t *testing.T) {
		// Arrange
		mockVM := &CreateClusterViewModel{}
		spy := &InteractorSpy{DoCreateClusterCalled: false, CreateClusterViewModelBox: mockVM}
		controller := NewController(WithInteractor(spy))

		// Act
		controller.CreateCluster()

		// Assert
		called := spy.DoCreateClusterCalled
		if called != true {
			t.Errorf("got: %v\nwant: %v", called, true)
		}
		send := spy.CreateClusterRequestBox
		if send == nil {
			t.Errorf("got: %v\nwant: %v", send, "not nil")
		}
	})
}
