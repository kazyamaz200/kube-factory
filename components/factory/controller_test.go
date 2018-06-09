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
