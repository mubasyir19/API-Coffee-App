package requests

type OrderDetails struct {
	ProductID string  `json:"productId" binding:"required"`
	Quantity  int     `json:"quantity" binding:"required"`
	SubTotal  float64 `json:"subTotal" binding:"required"`
}

type OrderRequest struct {
	CustomerID    string         `json:"customerId" binding:"required"`
	Address       string         `json:"address" `
	PhoneNumber   string         `json:"phoneNumber" `
	PaymentMethod string         `json:"paymentMethod" binding:"required"`
	Items         []OrderDetails `json:"items" binding:"required"`
	TotalPrice    float64        `json:"totalPrice" binding:"required"`
	AdminFee      float64        `json:"adminFee"`
}
