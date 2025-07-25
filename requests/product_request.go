package requests

type ProductInput struct {
	Name        string  `json:"name"`
	CategoryID  string  `json:"category_id"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Image       string  `json:"image_product"`
}
