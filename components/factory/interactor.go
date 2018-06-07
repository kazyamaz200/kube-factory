package factory

import "github.com/kyamazawa/glue-go/model"

// Interactor is ...
type Interactor struct {
	presenter Presentation
	store     Store
}

// NewInteractor is ...
func NewInteractor(opts ...InteractorOption) *Interactor {
	interactor := &Interactor{}

	for _, opt := range opts {
		opt(interactor)
	}

	if interactor.presenter == nil {
		presenter := NewPresenter()
		interactor.presenter = presenter
	}

	return interactor
}

// InteractorOption is ...
type InteractorOption func(*Interactor)

// WithPresenter is ...
func WithPresenter(p Presentation) InteractorOption {
	return func(s *Interactor) {
		if p != nil {
			s.presenter = p
		}
	}
}

// WithStore is ...
func WithStore(p Store) InteractorOption {
	return func(s *Interactor) {
		if p != nil {
			s.store = p
		}
	}
}

// Store is ...
type Store interface {
	SaveUser(user *model.User) (*model.User, error)
	FetchUserByID(userID string) (*model.User, error)
	SaveCluster(cluster *model.Cluster) (*model.Cluster, error)
	FetchClusterByID(clusterID string) (*model.Cluster, error)
}

// Interaction is ...
type Interaction interface {
	DoSomeUsecase(req *SomeUsecaseRequest, callback func(*SomeUsecaseViewModel, error))
}

// DoSomeUsecase is ...
func (s *Interactor) DoSomeUsecase(req *SomeUsecaseRequest, callback func(*SomeUsecaseViewModel, error)) {
	res := &SomeUsecaseResponse{}
	s.presenter.PresentSomeUsecase(res, callback)
}
