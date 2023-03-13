package article

type ArticleRequest struct {
	Title   string `binding:"required"`
	Content string `binding:"required"`
	Media   string
}

type ArticleUpdateRequest struct {
	Title   string
	Content string
	Media   string
}
