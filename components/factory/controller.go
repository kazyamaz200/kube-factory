package factory

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

// SDK is ...
type SDK interface {
	CreateCluster() (*CreateClusterViewModel, error)
}

// CreateCluster is ...
func (s *Controller) CreateCluster() (*CreateClusterViewModel, error) {
	channel := make(chan func() (*CreateClusterViewModel, error))
	defer close(channel)

	req := &CreateClusterRequest{}
	go s.interactor.DoCreateCluster(req, func(vm *CreateClusterViewModel, err error) {
		channel <- (func() (*CreateClusterViewModel, error) { return vm, err })
	})

	ret, err := (<-channel)()
	return ret, err
}
