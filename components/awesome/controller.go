package awesome

// Controller is ...
type Controller struct {
	interactor Interaction
}

// NewController is ...
func NewController(opts ...ControllerOption) *Controller {
	controller := &Controller{}

	for _, opt := range opts {
		opt(controller)
	}

	if controller.interactor == nil {
		interactor := NewInteractor()
		controller.interactor = interactor
	}

	return controller
}

// ControllerOption is ...
type ControllerOption func(*Controller)

// WithInteractor is ...
func WithInteractor(i Interaction) ControllerOption {
	return func(s *Controller) {
		if i != nil {
			s.interactor = i
		}
	}
}

// SomeUsecase is ...
func (s *Controller) SomeUsecase() (*SomeUsecaseViewModel, error) {
	channel := make(chan func() (*SomeUsecaseViewModel, error))
	defer close(channel)

	req := &SomeUsecaseRequest{}
	go s.interactor.DoSomeUsecase(req, func(vm *SomeUsecaseViewModel, err error) {
		channel <- (func() (*SomeUsecaseViewModel, error) { return vm, err })
	})

	ret, err := (<-channel)()
	return ret, err
}
