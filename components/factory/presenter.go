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
	PresentListCluster(res *ListClusterResponse, callback func(*ListClusterViewModel, error))
	PresentDescribeCluster(res *DescribeClusterResponse, callback func(*DescribeClusterViewModel, error))
	PresentUpdateCluster(res *UpdateClusterResponse, callback func(*UpdateClusterViewModel, error))
	PresentDeleteCluster(res *DeleteClusterResponse, callback func(*DeleteClusterViewModel, error))
}

// PresentCreateCluster is ...
func (s *Presenter) PresentCreateCluster(res *CreateClusterResponse, callback func(*CreateClusterViewModel, error)) {
	vm := &CreateClusterViewModel{ID: res.ID}
	callback(vm, nil)
}

// PresentListCluster is ...
func (s *Presenter) PresentListCluster(res *ListClusterResponse, callback func(*ListClusterViewModel, error)) {
	vm := &ListClusterViewModel{ID: res.ID}
	callback(vm, nil)
}

// PresentDescribeCluster is ...
func (s *Presenter) PresentDescribeCluster(res *DescribeClusterResponse, callback func(*DescribeClusterViewModel, error)) {
	vm := &DescribeClusterViewModel{ID: res.ID}
	callback(vm, nil)
}

// PresentUpdateCluster is ...
func (s *Presenter) PresentUpdateCluster(res *UpdateClusterResponse, callback func(*UpdateClusterViewModel, error)) {
	vm := &UpdateClusterViewModel{ID: res.ID}
	callback(vm, nil)
}

// PresentDeleteCluster is ...
func (s *Presenter) PresentDeleteCluster(res *DeleteClusterResponse, callback func(*DeleteClusterViewModel, error)) {
	vm := &DeleteClusterViewModel{ID: res.ID}
	callback(vm, nil)
}
