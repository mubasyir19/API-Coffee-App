package handlers

import (
	"api-coffee-app/requests"
	"api-coffee-app/responses"
	"api-coffee-app/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type customerHandler struct {
	service services.CustomerService
}

func NewCustomerHandler(service *services.CustomerService) *customerHandler {
	return &customerHandler{service: *service}
}

func (h *customerHandler) Register(c *gin.Context) {
	var req requests.CustomerRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, responses.APIResponse{
			Code:    "BAD_REQUEST",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	customer, err := h.service.Register(&req)
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
		Data:    responses.CustomerReponseFromModel(customer),
	})

}

func (h *customerHandler) Login(c *gin.Context) {
	var inputLogin requests.CustomerLogin

	err := c.ShouldBind(&inputLogin)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, responses.APIResponse{
			Code:    "UNPROCESSABLE_CONTENT",
			Message: "Invalid username or password",
			Data:    nil,
		})
		return
	}

	loggedIn, err := h.service.Login(inputLogin)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, responses.APIResponse{
			Code:    "UNPROCESSABLE_CONTENT",
			Message: "Failed login",
			Data:    nil,
		})
		return
	}

	token, err := h.service.GenerateToken(loggedIn)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.APIResponse{
			Code:    "BAD_REQUEST",
			Message: "Login Failed",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, responses.APIResponse{
		Code:    "SUCCESS",
		Message: "Successfully login",
		Data:    token,
	})

}

func (h *customerHandler) VerifyAuth(c *gin.Context) {
	username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Kirim respons dengan username atau data user lain yang dibutuhkan
	c.JSON(http.StatusOK, gin.H{
		"message":  "User verified",
		"username": username,
	})
}
