package user

import (
	"net/http"

	domain "api-auth/internal/domain/user"
	request "api-auth/internal/handler/user/dto/request"
	service "api-auth/internal/service/user"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service service.UserService
}

// NewUserHandler constructor
func NewUserHandler(s service.UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	users, err := h.service.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req request.CreateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := &domain.User{
		Username:    req.Username,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Email:       req.Email,
		Phone:       &req.Phone,
		BirthDate:   req.BirthDate,
		CountryID:   req.CountryID,
		AddressLine: &req.AddressLine,
		IsActive:    true,
	}

	if err := h.service.CreateUser(user, req.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Limpiar hash antes de devolver
	user.PasswordHash = ""
	c.JSON(http.StatusCreated, user)
}
