package category

type CategoryRequest struct {
	Name        string `binding:"required"`
	Description string
}
