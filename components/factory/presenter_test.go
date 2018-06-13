package factory

import "testing"

func TestPresenter_PresentCreateCluster(t *testing.T) {
	t.Run("call received callback with CreateClusterViewModel and error", func(t *testing.T) {
		// Arrange
		presenter := NewPresenter()
		res := &CreateClusterResponse{}
		called := false
		var boxVM *CreateClusterViewModel
		var boxError error
		callback := func(vm *CreateClusterViewModel, err error) { called = true; boxVM = vm; boxError = err }

		// Act
		presenter.PresentCreateCluster(res, callback)

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
