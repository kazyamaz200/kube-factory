package factory

import "testing"

func TestNewController(t *testing.T) {
	t.Run("create controller and it has interactor", func(t *testing.T) {
		// Act
		controller := NewController()

		// Assert
		actual := controller.interactor
		if actual == nil {
			t.Errorf("got: %v\nwant: %v", actual, "not nil")
		}

	})

	t.Run("its interactor is compatible with Interaction", func(t *testing.T) {
		// Arrange
		expected := true

		// Act
		controller := NewController()

		// Assert
		_, actual := controller.interactor.(Interaction)
		if actual != expected {
			t.Errorf("got: %v\nwant: %v", actual, expected)
		}
	})

	t.Run("its interactor is injectable", func(t *testing.T) {
		// Arrange
		i1 := NewInteractor()
		i2 := &InteractorSpy{}

		// Act
		controller := NewController(WithInteractor(i1))

		// Assert
		actual := controller.interactor
		if actual != i1 {
			t.Errorf("got: %v\nwant: %v", actual, i1)
		}
		if actual == i2 {
			t.Errorf("got: %v\nwant: %v", actual, i2)
		}
	})

	t.Run("its interactor is not be nil", func(t *testing.T) {
		// Arrange
		var i1 *Interactor // nil

		// Act
		controller := NewController(WithInteractor(i1))

		// Assert
		actual := controller.interactor
		if actual == nil {
			t.Errorf("got: %v\nwant: %v", actual, "not nil")
		}
	})
}

type InteractorSpy struct {
	DoSomeUsecaseCalled     bool
	SomeUsecaseViewModelBox *SomeUsecaseViewModel
	SomeUsecaseRequestBox   *SomeUsecaseRequest
}

func (s *InteractorSpy) DoSomeUsecase(req *SomeUsecaseRequest, callback func(*SomeUsecaseViewModel, error)) {
	s.DoSomeUsecaseCalled = true
	s.SomeUsecaseRequestBox = req
	callback(s.SomeUsecaseViewModelBox, nil)
}

func TestController_SomeUsecase(t *testing.T) {
	t.Run("return SomeUsecaseViewModel and error", func(t *testing.T) {
		// Arrange
		mockVM := &SomeUsecaseViewModel{}
		spy := &InteractorSpy{SomeUsecaseViewModelBox: mockVM}
		controller := NewController(WithInteractor(spy))

		// Act
		actual, err := controller.SomeUsecase()

		// Assert
		if actual != mockVM {
			t.Errorf("got: %v\nwant: %v", actual, mockVM)
		}
		if err != nil {
			t.Errorf("got: %v\nwant: %v", err, nil)
		}
	})

	t.Run("call DoSomeUsecase with SomeUsecaseRequest", func(t *testing.T) {
		// Arrange
		mockVM := &SomeUsecaseViewModel{}
		spy := &InteractorSpy{DoSomeUsecaseCalled: false, SomeUsecaseViewModelBox: mockVM}
		controller := NewController(WithInteractor(spy))

		// Act
		controller.SomeUsecase()

		// Assert
		called := spy.DoSomeUsecaseCalled
		if called != true {
			t.Errorf("got: %v\nwant: %v", called, true)
		}
		send := spy.SomeUsecaseRequestBox
		if send == nil {
			t.Errorf("got: %v\nwant: %v", send, "not nil")
		}
	})
}
