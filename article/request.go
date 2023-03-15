package article

type ArticleRequest struct {
	Title      string `binding:"required"`
	Content    string `binding:"required"`
	Media      string
	CategoryID int
	Author     string `binding:"required"`
}

type ArticleUpdateRequest struct {
	Title      string
	Content    string
	Media      string
	CategoryID int
	Author     string
}
