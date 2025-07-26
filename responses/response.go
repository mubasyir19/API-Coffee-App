package responses

import "api-coffee-app/models"

type APIResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type CustomerResponse struct {
	ID          string `json:"id"`
	Fullname    string `json:"fullname"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phoneNumber"`
	Address     string `json:"address"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func CustomerReponseFromModel(customer *models.Customer) CustomerResponse {
	return CustomerResponse{
		ID:          customer.ID,
		Fullname:    customer.Fullname,
		Username:    customer.Username,
		Email:       customer.Email,
		PhoneNumber: customer.PhoneNumber,
		Address:     customer.Address,
		CreatedAt:   customer.CreatedAt.Local().String(),
		UpdatedAt:   customer.UpdatedAt.Local().String(),
	}

}

type ProductResponse struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	CategoryID  string  `json:"category_id"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Image       string  `json:"image_product"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

func ProductReponseFromModel(product *models.Product) ProductResponse {
	return ProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		CategoryID:  product.CategoryID,
		Description: product.Description,
		Price:       product.Price,
		Image:       product.Image,
		CreatedAt:   product.CreatedAt.Local().String(),
		UpdatedAt:   product.UpdatedAt.Local().String(),
	}

}

type CategoryResponse struct {
	ID        string            `json:"id"`
	Name      string            `json:"name"`
	CreatedAt string            `json:"created_at"`
	UpdatedAt string            `json:"updated_at"`
	Products  []ProductResponse `json:"products"`
}

func CategoryReponseFromModel(category *models.Category) CategoryResponse {
	var productResponses []ProductResponse
	for _, p := range category.Products {
		productResponses = append(productResponses, ProductReponseFromModel(&p))
	}

	return CategoryResponse{
		ID:        category.ID,
		Name:      category.Name,
		CreatedAt: category.CreatedAt.Local().String(),
		UpdatedAt: category.UpdatedAt.Local().String(),
		Products:  productResponses,
	}
}
