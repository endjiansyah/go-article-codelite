package article

import "time"

type ArticleResponse struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Media      string `json:"media"`
	Author     string `json:"author"`
	CategoryID int
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
