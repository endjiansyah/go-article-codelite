package article

type Service interface {
	GetAll(Category int, Page int, Limit int) ([]Article, error)
	GetById(ID int) (Article, error)
	Create(articleReq ArticleRequest) (Article, error)
	Update(ID int, articleReq ArticleUpdateRequest) (Article, error)
	Delete(ID int) (Article, error)
}

type service struct {
	repository ArticleRepo
}

func NewService(repo ArticleRepo) *service {
	return &service{repo}
}

func (s *service) GetAll(Category int, Page int, Limit int) ([]Article, error) {
	articles, err := s.repository.GetAll(Category, Page, Limit)
	return articles, err
}

func (s *service) GetById(ID int) (Article, error) {
	article, err := s.repository.GetById(ID)
	return article, err
}

func (s *service) Create(articleReq ArticleRequest) (Article, error) {

	payload := Article{
		Title:      articleReq.Title,
		Content:    articleReq.Content,
		Media:      articleReq.Media,
		Author:     articleReq.Author,
		CategoryID: articleReq.CategoryID,
	}
	article, err := s.repository.Create(payload)
	return article, err
}

func (s *service) Update(ID int, articleReq ArticleUpdateRequest) (Article, error) {
	cst, _ := s.repository.GetById(ID)

	if articleReq.Title != "" {
		cst.Title = articleReq.Title
	}
	if articleReq.Content != "" {
		cst.Content = articleReq.Content
	}
	if articleReq.Media != "" {
		cst.Media = articleReq.Media
	}
	if articleReq.Author != "" {
		cst.Author = articleReq.Author
	}
	if articleReq.CategoryID != 0 {
		cst.CategoryID = articleReq.CategoryID
	}
	article, err := s.repository.Update(cst)
	return article, err
}

func (s *service) Delete(ID int) (Article, error) {
	cst, _ := s.repository.GetById(ID)

	article, err := s.repository.Delete(cst)
	return article, err
}
