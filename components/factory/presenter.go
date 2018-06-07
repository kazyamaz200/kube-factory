package factory

// Presenter is ...
type Presenter struct {
}

// NewPresenter is ...
func NewPresenter() *Presenter {
	presenter := &Presenter{}
	return presenter
}

// PresenterOption is ...
type PresenterOption func(*Presenter)

// Presentation is ...
type Presentation interface {
	PresentSomeUsecase(res *SomeUsecaseResponse, callback func(*SomeUsecaseViewModel, error))
}

// PresentSomeUsecase is ...
func (s *Presenter) PresentSomeUsecase(res *SomeUsecaseResponse, callback func(*SomeUsecaseViewModel, error)) {
	vm := &SomeUsecaseViewModel{}
	callback(vm, nil)
}
