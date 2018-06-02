package awesome

import "testing"

func TestPresenter_PresentSomeUsecase(t *testing.T) {
	t.Run("call received callback with SomeUsecaseViewModel and error", func(t *testing.T) {
		// Arrange
		presenter := NewPresenter()
		res := &SomeUsecaseResponse{}
		called := false
		var boxVM *SomeUsecaseViewModel
		var boxError error
		callback := func(vm *SomeUsecaseViewModel, err error) { called = true; boxVM = vm; boxError = err }

		// Act
		presenter.PresentSomeUsecase(res, callback)

		// Assert
		if called != true {
			t.Errorf("got: %v\nwant: %v", called, true)
		}
		if boxVM == nil {
			t.Errorf("got: %v\nwant: %v", boxVM, "not nil")
		}
		if boxError != nil {
			t.Errorf("got: %v\nwant: %v", boxError, nil)
		}
	})
}
