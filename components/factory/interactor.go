package factory

import "github.com/kyamazawa/kube-factory/provider"

// Interactor is ...
type Interactor struct {
	presenter Presentation
	store     provider.StoreProtocol
}

// NewInteractor is ...
func NewInteractor(opts ...InteractorOption) *Interactor {
	interactor := &Interactor{}
	for _, opt := range opts {
		opt(interactor)
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
func WithStore(p provider.StoreProtocol) InteractorOption {
	return func(s *Interactor) {
		if p != nil {
			s.store = p
		}
	}
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
