package awesome

import "testing"

func TestPresentSomeUsecase(t *testing.T) {
	t.Run("it should call received callback function", func(t *testing.T) {
		// Arrange
		presenter := NewPresenter()
		res := &SomeUsecaseResponse{}
		called := false
		callback := func(vm *SomeUsecaseViewModel, err error) { called = true }

		// Act
		presenter.PresentSomeUsecase(res, callback)

		// Assert
		if called != true {
			t.Errorf("got: %v\nwant: %v", called, "Should be true")
		}
	})

	t.Run("the call should receive expected vm and error", func(t *testing.T) {
		// Arrange
		presenter := NewPresenter()
		res := &SomeUsecaseResponse{}
		var boxVM *SomeUsecaseViewModel
		var boxError error
		callback := func(vm *SomeUsecaseViewModel, err error) { boxVM = vm; boxError = err }

		// Act
		presenter.PresentSomeUsecase(res, callback)

		// Assert
		if boxVM == nil {
			t.Errorf("got: %v\nwant: %v", boxVM, "Should not be nil")
		}
		if boxError != nil {
			t.Errorf("got: %v\nwant: %v", boxError, "Should be nil")
		}
	})
}
