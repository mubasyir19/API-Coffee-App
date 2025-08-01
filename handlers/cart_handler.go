package handlers

import (
	"api-coffee-app/requests"
	"api-coffee-app/responses"
	"api-coffee-app/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type cartHandler struct {
	service services.CartService
}

func NewCartHandler(service *services.CartService) *cartHandler {
	return &cartHandler{service: *service}
}

func (h *cartHandler) GetCartItems(c *gin.Context) {
	customerID := c.Query("customerId")
	if customerID == "" {
		c.JSON(http.StatusBadRequest, responses.APIResponse{
			Code:    "BAD_REQUEST",
			Message: "User ID is required",
			Data:    nil,
		})
		return
	}

	cartItems, err := h.service.GetItems(customerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.APIResponse{
			Code:    "ERROR",
			Message: "Failed to fetch cart items",
			Data:    nil,
		})
		return
	}

	if len(cartItems) == 0 {
		c.JSON(http.StatusOK, responses.APIResponse{
			Code:    "SUCCESS",
			Message: "Cart is empty",
			Data:    nil,
		})
		return
	}

	var products []responses.ProductCartResponse
	var customerResp responses.CustomerResponse
	var totalAmount float64
	var totalItems int

	for _, item := range cartItems {
		product := item.Product
		customer := item.Customer

		productResp := responses.ProductCartResponse{
			ID:           product.ID,
			Name:         product.Name,
			Description:  product.Description,
			Price:        product.Price,
			ImageProduct: product.Image,
			Quantity:     item.Quantity,
			TotalPrice:   item.TotalPrice,
		}

		products = append(products, productResp)
		totalAmount += item.TotalPrice
		totalItems = len(products)

		if customerResp.ID == "" {
			customerResp = responses.CustomerResponse{
				ID:          customer.ID,
				Fullname:    customer.Fullname,
				Username:    customer.Username,
				Email:       customer.Email,
				PhoneNumber: customer.PhoneNumber,
				Address:     customer.Address,
			}
		}
	}

	c.JSON(http.StatusOK, responses.APIResponse{
		Code:    "SUCCESS",
		Message: "Successfully get data cart",
		Data: responses.CartSummaryResponse{
			TotalAmount: totalAmount,
			TotalItems:  totalItems,
			Customer:    customerResp,
			Products:    products,
		},
	})
}

func (h *cartHandler) AddItemToCart(c *gin.Context) {
	var input requests.CartInput
	if err := c.ShouldBind(&input); err != nil {
		log.Println("Error binding data:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": "Data tidak lengkap",
			"data":    nil,
		})
		return
	}

	addItem, err := h.service.AddItem(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "Gagal tambah item",
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  201,
		"message": "Berhasil tambah item",
		"data":    addItem,
	})
}

func (h *cartHandler) UpdateItemCart(c *gin.Context) {
	var input requests.CartInput
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": "Data tidak lengkap",
			"data":    nil,
		})
		return
	}

	updateItem, err := h.service.UpdateItem(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "Gagal tambah item",
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Berhasil update item cart",
		"data":    updateItem,
	})
}

func (h *cartHandler) RemoveItemCart(c *gin.Context) {
	id := c.Param("id")

	if err := h.service.RemoveItem(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "Gagal hapus item",
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Berhasil hapus item",
		"data":    nil,
	})
}
