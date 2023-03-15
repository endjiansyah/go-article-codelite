package article

type ArticleRequest struct {
	Title      string `binding:"required"`
	Content    string `binding:"required"`
	CategoryID int
	Author     string `binding:"required"`
}

type ArticleUpdateRequest struct {
	Title      string
	Content    string
	CategoryID int
	Author     string
}

type MediapostRequest struct {
	Media     string
	Type      string
	ArticleID int
}
