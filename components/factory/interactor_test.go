package factory

import "testing"

func TestNewInteractor(t *testing.T) {
	t.Run("its presenter is injectable", func(t *testing.T) {
		// Arrange
		i := &PresenterSpy{}

		// Act
		interactor := NewInteractor(WithPresenter(i))

		// Assert
		_, ok := interactor.presenter.(Presentation)
		if !ok {
			t.Errorf("got: %v\nwant: %v", ok, true)
		}
	})
}

type PresenterSpy struct {
	PresentCreateClusterCalled bool
	CreateClusterViewModelBox  *CreateClusterViewModel
	CreateClusterResponseBox   *CreateClusterResponse
}

func (s *PresenterSpy) PresentCreateCluster(res *CreateClusterResponse, callback func(*CreateClusterViewModel, error)) {
	s.PresentCreateClusterCalled = true
	s.CreateClusterResponseBox = res
	callback(s.CreateClusterViewModelBox, nil)
}
func TestInteractor_DoCreateCluster(t *testing.T) {
	t.Run("call PresentCreateCluster with CreateClusterResponse", func(t *testing.T) {
		// Arrange
		spy := &PresenterSpy{PresentCreateClusterCalled: false}
		interactor := NewInteractor(WithPresenter(spy))
		req := &CreateClusterRequest{}
		callback := func(*CreateClusterViewModel, error) {}

		// Act
		interactor.DoCreateCluster(req, callback)

		// Assert
		called := spy.PresentCreateClusterCalled
		if called != true {
			t.Errorf("got: %v\nwant: %v", called, true)
		}
		send := spy.CreateClusterResponseBox
		if send == nil {
			t.Errorf("got: %v\nwant: %v", send, "not nil")
		}
	})
}
