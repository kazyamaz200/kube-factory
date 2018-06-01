package awesome

import "testing"

func TestNewInteractor(t *testing.T) {
	t.Run("Interactor has default presenter", func(t *testing.T) {
		// Act
		interactor := NewInteractor()

		// Assert
		actual := interactor.presenter
		if actual == nil {
			t.Errorf("got: %v\nwant: %v", actual, "NotNil")
		}

	})

	t.Run("The presenter is Presentation interface compatible", func(t *testing.T) {
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

	t.Run("The presenter is injectable", func(t *testing.T) {
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

	t.Run("The presenter should not be nil with wrong usage", func(t *testing.T) {
		// Arrange
		var i1 *Presenter // nil

		// Act
		interactor := NewInteractor(WithPresenter(i1))

		// Assert
		actual := interactor.presenter
		if actual == nil {
			t.Errorf("got: %v\nwant: %v", actual, "NotNil")
		}
	})
}

type PresenterSpy struct {
}

func (s *PresenterSpy) PresentSomeUsecase(res *SomeUsecaseResponse, callback func(*SomeUsecaseViewModel, error)) {
	vm := &SomeUsecaseViewModel{}
	callback(vm, nil)
}
