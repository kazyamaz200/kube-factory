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
	PresentCreateCluster(res *CreateClusterResponse, callback func(*CreateClusterViewModel, error))
}

// PresentCreateCluster is ...
func (s *Presenter) PresentCreateCluster(res *CreateClusterResponse, callback func(*CreateClusterViewModel, error)) {
	vm := &CreateClusterViewModel{}
	callback(vm, nil)
}
