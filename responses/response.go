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
	PhoneNumber string `json:"phoneNumber"`
	Address     string `json:"address"`
}

func CustomerReponseFromModel(customer *models.Customer) CustomerResponse {
	return CustomerResponse{
		ID:          customer.ID,
		Fullname:    customer.Fullname,
		Username:    customer.Username,
		Email:       customer.Email,
		PhoneNumber: customer.PhoneNumber,
		Address:     customer.Address,
	}

}

type ProductResponse struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Slug        string  `json:"slug"`
	CategoryID  string  `json:"category_id"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Image       string  `json:"image_product"`
}

func ProductReponseFromModel(product *models.Product) ProductResponse {
	return ProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Slug:        product.Slug,
		CategoryID:  product.CategoryID,
		Description: product.Description,
		Price:       product.Price,
		Image:       product.Image,
	}

}

type CategoryResponse struct {
	ID       string            `json:"id"`
	Name     string            `json:"name"`
	Products []ProductResponse `json:"products"`
}

func CategoryReponseFromModel(category *models.Category) CategoryResponse {
	var productResponses []ProductResponse
	for _, p := range category.Products {
		productResponses = append(productResponses, ProductReponseFromModel(&p))
	}

	return CategoryResponse{
		ID:       category.ID,
		Name:     category.Name,
		Products: productResponses,
	}
}

type CartSummaryResponse struct {
	CartID      string                `json:"id"`
	TotalAmount float64               `json:"total_amount"`
	TotalItems  int                   `json:"total_items"`
	Customer    CustomerResponse      `json:"customer"`
	Products    []ProductCartResponse `json:"products"`
}

type ProductCartResponse struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	Slug         string  `json:"slug"`
	Description  string  `json:"description"`
	Price        float64 `json:"price"`
	ImageProduct string  `json:"image_product"`
	Quantity     int     `json:"quantity"`
	TotalPrice   float64 `json:"total_price"`
}
