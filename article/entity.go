package article

import "time"

type Article struct {
	ID         int
	Title      string
	Content    string
	CategoryID int
	Author     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
type Media struct {
	ID        int
	Media     string
	Type      string
	ArticleID int
	CreatedAt time.Time
	UpdatedAt time.Time
}
