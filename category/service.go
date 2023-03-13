package category

type Service interface {
	GetAll() ([]Category, error)
	GetById(ID int) (Category, error)
	Create(categoryReq CategoryRequest) (Category, error)
	Update(ID int, categoryReq CategoryUpdateRequest) (Category, error)
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

func (s *service) GetById(ID int) (Category, error) {
	category, err := s.repository.GetById(ID)
	return category, err
}

func (s *service) Create(categoryReq CategoryRequest) (Category, error) {

	payload := Category{
		Name:        categoryReq.Name,
		Description: categoryReq.Description,
	}
	category, err := s.repository.Create(payload)
	return category, err
}

func (s *service) Update(ID int, categoryReq CategoryUpdateRequest) (Category, error) {
	cst, _ := s.repository.GetById(ID)

	if categoryReq.Name != "" {
		cst.Name = categoryReq.Name
	}
	if categoryReq.Description != "" {
		cst.Description = categoryReq.Description
	}
	category, err := s.repository.Update(cst)
	return category, err
}
