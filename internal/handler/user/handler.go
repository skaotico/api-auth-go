// ============================================================
// @file: handler.go
// @author: Yosemar Andrade
// @created: 2025-11-20
// @description: Handler para la gestión de usuarios.
// ============================================================

// Package user contiene los handlers para la gestión de usuarios.
package user

import (
	"net/http"

	domain "api-auth/internal/domain/user"
	request "api-auth/internal/handler/user/dto/request"
	userService "api-auth/internal/service/user"

	"github.com/gin-gonic/gin"
)

// Handler representa el controlador encargado de manejar las solicitudes relacionadas con usuarios.
type Handler struct {
	service userService.ServiceInterface
}

// NewHandler crea una nueva instancia de Handler.
//
// Parámetros:
//
//	s: Implementación de ServiceInterface.
//
// Retorna:
//
//	*Handler: Instancia inicializada.
func NewHandler(s userService.ServiceInterface) *Handler {
	return &Handler{service: s}
}

// GetUsers obtiene la lista de usuarios.
//
// @Summary Obtener usuarios
// @Description Obtiene todos los usuarios registrados.
// @Tags User
// @Produce json
// @Success 200 {array} domain.User
// @Failure 500 {object} map[string]string
// @Router /v1/users [get]
func (h *Handler) GetUsers(c *gin.Context) {
	users, err := h.service.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

// CreateUser crea un nuevo usuario.
//
// @Summary Crear usuario
// @Description Crea un nuevo usuario en el sistema.
// @Tags User
// @Accept json
// @Produce json
// @Param request body request.CreateUserRequest true "Datos del usuario"
// @Success 201 {object} domain.User
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /v1/users [post]
func (h *Handler) CreateUser(c *gin.Context) {
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
