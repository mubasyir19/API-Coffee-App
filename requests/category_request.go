package requests

type CategoryRequest struct {
	Name string `json:"name" binding:"required"`
}
