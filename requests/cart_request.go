package requests

type CartInput struct {
	CustomerID string `json:"customerId" binding:"required"`
	ProductID  string `json:"productId" binding:"required"`
	Quantity   int    `json:"quantity"`
}
