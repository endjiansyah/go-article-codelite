package article

import "time"

type ArticleResponse struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Author     string `json:"author"`
	CategoryID int
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
type ArticleAllResponse struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Author     string `json:"author"`
	Media      MediaResponse
	CategoryID int
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
type MediaResponse struct {
	ID        int `json:"id"`
	Media     string
	Type      string
	ArticleID int
}
