package handlers

import (
	"api-coffee-app/requests"
	"api-coffee-app/responses"
	"api-coffee-app/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type productHandler struct {
	service services.ProductService
}

func NewProductHandler(service *services.ProductService) *productHandler {
	return &productHandler{service: *service}
}

func (h *productHandler) GetAllProducts(c *gin.Context) {
	products, err := h.service.FindAllProduct()
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.APIResponse{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	var productResponse []responses.ProductResponse
	for _, product := range products {
		productResponse = append(productResponse, responses.ProductReponseFromModel(&product))
	}

	c.JSON(http.StatusOK, responses.APIResponse{
		Code:    "SUCCESS",
		Message: "Successfully find all products",
		Data:    productResponse,
	})
}

func (h *productHandler) GetProdutcByID(c *gin.Context) {
	id := c.Param("id")

	product, err := h.service.FindProductByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.APIResponse{
			Code:    "BAD_REQUEST",
			Message: "Failed get data products",
			Data:    nil,
		})
		return
	}

	response := responses.ProductReponseFromModel(product)

	c.JSON(http.StatusOK, responses.APIResponse{
		Code:    "SUCCESS",
		Message: "Successfully find all products",
		Data:    response,
	})
}

func (h *productHandler) GetProductByName(c *gin.Context) {
	productName := c.Param("name")

	product, err := h.service.FindProductByName(productName)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.APIResponse{
			Code:    "BAD_REQUEST",
			Message: "Failed get data products",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, responses.APIResponse{
		Code:    "SUCCESS",
		Message: "Successfully find all products",
		Data:    product,
	})
}

func (h *productHandler) GetProductsByCategory(c *gin.Context) {
	categoryID := c.Param("category_id")

	products, err := h.service.FindProductByCategory(categoryID)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.APIResponse{
			Code:    "BAD_REQUEST",
			Message: "Failed get data products",
			Data:    nil,
		})
		return
	}

	var ProductsCategory []responses.ProductResponse
	for _, product := range products {
		ProductsCategory = append(ProductsCategory, responses.ProductReponseFromModel(&product))
	}

	c.JSON(http.StatusOK, responses.APIResponse{
		Code:    "SUCCESS",
		Message: "Successfully find all products",
		Data:    ProductsCategory,
	})
}

func (h *productHandler) AddProduct(c *gin.Context) {
	var input requests.ProductInput

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, responses.APIResponse{
			Code:    "BAD_REQUEST",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	product, err := h.service.AddProduct(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.APIResponse{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusCreated, responses.APIResponse{
		Code:    "CREATED",
		Message: "Customer account created successfully",
		Data:    responses.ProductReponseFromModel(product),
	})
}
