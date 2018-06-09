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
	PresentSomeUsecaseCalled bool
	SomeUsecaseViewModelBox  *SomeUsecaseViewModel
	SomeUsecaseResponseBox   *SomeUsecaseResponse
}

func (s *PresenterSpy) PresentSomeUsecase(res *SomeUsecaseResponse, callback func(*SomeUsecaseViewModel, error)) {
	s.PresentSomeUsecaseCalled = true
	s.SomeUsecaseResponseBox = res
	callback(s.SomeUsecaseViewModelBox, nil)
}
func TestInteractor_DoSomeUsecase(t *testing.T) {
	t.Run("call PresentSomeUsecase with SomeUsecaseResponse", func(t *testing.T) {
		// Arrange
		spy := &PresenterSpy{PresentSomeUsecaseCalled: false}
		interactor := NewInteractor(WithPresenter(spy))
		req := &SomeUsecaseRequest{}
		callback := func(*SomeUsecaseViewModel, error) {}

		// Act
		interactor.DoSomeUsecase(req, callback)

		// Assert
		called := spy.PresentSomeUsecaseCalled
		if called != true {
			t.Errorf("got: %v\nwant: %v", called, true)
		}
		send := spy.SomeUsecaseResponseBox
		if send == nil {
			t.Errorf("got: %v\nwant: %v", send, "not nil")
		}
	})
}
