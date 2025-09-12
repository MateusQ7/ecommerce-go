package http

import (
	"net/http"

	"github.com/MateusQ7/ecommerce-go/auth-service/internal/domain"
	"github.com/MateusQ7/ecommerce-go/auth-service/internal/usecase"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *usecase.UserService
}

func NewUserHandler(service *usecase.UserService) *UserHandler {
	return &UserHandler{service}
}

func (h *UserHandler) CreateNewUser(c *gin.Context) {
	var u domain.User

	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.CreateNewUser(&u); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, u)
}

func (h *UserHandler) ListUsers(c *gin.Context) {
	users, err := h.service.ListUsers()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}
