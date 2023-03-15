package article

type Service interface {
	GetAll(Category int, Page int, Limit int) ([]Article, error)
	GetById(ID int) (Article, error)
	GetMediaById(ID int) ([]Media, error)
	Create(articleReq ArticleRequest) (Article, error)
	CreateMedia(mediaReq MediapostRequest) (Media, error)
	Update(ID int, articleReq ArticleUpdateRequest) (Article, error)
	Delete(ID int) (Article, error)
	DeleteMedia(ID int) (Media, error)
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

func (s *service) GetMediaById(ID int) ([]Media, error) {
	media, err := s.repository.GetMediaById(ID)
	return media, err
}
func (s *service) Create(articleReq ArticleRequest) (Article, error) {

	payload := Article{
		Title:      articleReq.Title,
		Content:    articleReq.Content,
		Author:     articleReq.Author,
		CategoryID: articleReq.CategoryID,
	}
	article, err := s.repository.Create(payload)
	return article, err
}
func (s *service) CreateMedia(mediaReq MediapostRequest) (Media, error) {

	payload := Media{
		Media:     mediaReq.Media,
		Type:      mediaReq.Type,
		ArticleID: mediaReq.ArticleID,
	}
	article, err := s.repository.CreateMedia(payload)
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
	// if articleReq.Media != "" {
	// 	cst.Media = articleReq.Media
	// }
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

func (s *service) DeleteMedia(ID int) (Media, error) {
	cst, _ := s.repository.GetMediaId(ID)

	media, err := s.repository.DeleteMedia(cst)
	return media, err
}
