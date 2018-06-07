package factory

import "testing"

func TestNewInteractor(t *testing.T) {
	t.Run("create controller and it has presenter", func(t *testing.T) {
		// Act
		interactor := NewInteractor()

		// Assert
		actual := interactor.presenter
		if actual == nil {
			t.Errorf("got: %v\nwant: %v", actual, "not nil")
		}

	})

	t.Run("its presenter is compatible with Presentation", func(t *testing.T) {
		// Arrange
		expected := true

		// Act
		interactor := NewInteractor()

		// Assert
		_, actual := interactor.presenter.(Presentation)
		if actual != expected {
			t.Errorf("got: %v\nwant: %v", actual, expected)
		}
	})

	t.Run("its presenter is injectable", func(t *testing.T) {
		// Arrange
		i1 := NewPresenter()
		i2 := &PresenterSpy{}

		// Act
		interactor := NewInteractor(WithPresenter(i1))

		// Assert
		actual := interactor.presenter
		if actual != i1 {
			t.Errorf("got: %v\nwant: %v", actual, i1)
		}
		if actual == i2 {
			t.Errorf("got: %v\nwant: %v", actual, i2)
		}
	})

	t.Run("its presenter is not be nil", func(t *testing.T) {
		// Arrange
		var i1 *Presenter // nil

		// Act
		interactor := NewInteractor(WithPresenter(i1))

		// Assert
		actual := interactor.presenter
		if actual == nil {
			t.Errorf("got: %v\nwant: %v", actual, "not nil")
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
