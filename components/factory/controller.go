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
	ListCluster() (*ListClusterViewModel, error)
	DescribeCluster(id string) (*DescribeClusterViewModel, error)
	UpdateCluster(id string) (*UpdateClusterViewModel, error)
	DeleteCluster(id string) (*DeleteClusterViewModel, error)
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

// ListCluster is ...
func (s *Controller) ListCluster() (*ListClusterViewModel, error) {
	channel := make(chan func() (*ListClusterViewModel, error))
	defer close(channel)

	req := &ListClusterRequest{}
	go s.interactor.DoListCluster(req, func(vm *ListClusterViewModel, err error) {
		channel <- (func() (*ListClusterViewModel, error) { return vm, err })
	})

	ret, err := (<-channel)()
	return ret, err
}

// DescribeCluster is ...
func (s *Controller) DescribeCluster(id string) (*DescribeClusterViewModel, error) {
	channel := make(chan func() (*DescribeClusterViewModel, error))
	defer close(channel)

	req := &DescribeClusterRequest{}
	go s.interactor.DoDescribeCluster(req, func(vm *DescribeClusterViewModel, err error) {
		channel <- (func() (*DescribeClusterViewModel, error) { return vm, err })
	})

	ret, err := (<-channel)()
	return ret, err
}

// UpdateCluster is ...
func (s *Controller) UpdateCluster(id string) (*UpdateClusterViewModel, error) {
	channel := make(chan func() (*UpdateClusterViewModel, error))
	defer close(channel)

	req := &UpdateClusterRequest{}
	go s.interactor.DoUpdateCluster(req, func(vm *UpdateClusterViewModel, err error) {
		channel <- (func() (*UpdateClusterViewModel, error) { return vm, err })
	})

	ret, err := (<-channel)()
	return ret, err
}

// DeleteCluster is ...
func (s *Controller) DeleteCluster(id string) (*DeleteClusterViewModel, error) {
	channel := make(chan func() (*DeleteClusterViewModel, error))
	defer close(channel)

	req := &DeleteClusterRequest{}
	go s.interactor.DoDeleteCluster(req, func(vm *DeleteClusterViewModel, err error) {
		channel <- (func() (*DeleteClusterViewModel, error) { return vm, err })
	})

	ret, err := (<-channel)()
	return ret, err
}
