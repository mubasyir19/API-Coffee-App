package handlers

import (
	"api-coffee-app/requests"
	"api-coffee-app/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type orderHandler struct {
	service services.OrderService
}

func NewOrderHandler(service *services.OrderService) *orderHandler {
	return &orderHandler{service: *service}
}

func (h *orderHandler) Checkout(c *gin.Context) {
	var input requests.OrderRequest
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": "Data tidak lengkap",
			"data":    nil,
		})
		return
	}

	order, err := h.service.Checkout(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "Gagal Checkout",
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  201,
		"message": "Berhasil membuat order",
		"data":    order,
	})
}
