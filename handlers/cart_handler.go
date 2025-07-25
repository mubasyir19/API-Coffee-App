package handlers

import (
	"api-coffee-app/requests"
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
