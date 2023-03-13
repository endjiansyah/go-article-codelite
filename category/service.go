package category

type Service interface {
	GetAll() ([]Category, error)
}

type service struct {
	repository CategoryRepo
}

func NewService(repo CategoryRepo) *service {
	return &service{repo}
}

func (s *service) GetAll() ([]Category, error) {
	categories, err := s.repository.GetAll()
	return categories, err
}
