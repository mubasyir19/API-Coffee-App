package handlers

import (
	"api-coffee-app/requests"
	"api-coffee-app/responses"
	"api-coffee-app/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type categoryHandler struct {
	service services.CategoryService
}

func NewCategoryHandler(service *services.CategoryService) *categoryHandler {
	return &categoryHandler{service: *service}
}

func (h *categoryHandler) GetAllCategories(c *gin.Context) {
	categories, err := h.service.GetAllCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.APIResponse{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	var categoriesResponse []responses.CategoryResponse
	for _, category := range categories {
		categoriesResponse = append(categoriesResponse, responses.CategoryReponseFromModel(&category))
	}

	c.JSON(http.StatusOK, responses.APIResponse{
		Code:    "SUCCESS",
		Message: "Successfully find all categories",
		Data:    categoriesResponse,
	})
}

func (h *categoryHandler) GetCategoryByID(c *gin.Context) {
	id := c.Param("id")

	category, err := h.service.GetCategory(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.APIResponse{
			Code:    "BAD_REQUEST",
			Message: "Failed get data category",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, responses.APIResponse{
		Code:    "SUCCESS",
		Message: "Successfully get data category",
		Data:    category,
	})
}

func (h *categoryHandler) AddCategory(c *gin.Context) {
	var req requests.CategoryRequest

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, responses.APIResponse{
			Code:    "BAD_REQUEST",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	category, err := h.service.CreateCategory(&req)
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
		Data:    responses.CategoryReponseFromModel(category),
	})
}
