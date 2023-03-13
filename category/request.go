package category

type CategoryRequest struct {
	Name        string `binding:"required"`
	Description string
}

type CategoryUpdateRequest struct {
	Name        string
	Description string
}
