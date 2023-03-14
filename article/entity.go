package article

import "time"

type Article struct {
	ID         int
	Title      string
	Media      string
	Content    string
	CategoryID int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
