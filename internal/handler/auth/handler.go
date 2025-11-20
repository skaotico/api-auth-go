// ============================================================
// @file: handler.go
// @author: Yosemar Andrade
// @created: 2025-11-20
// @description: Handler para autenticación de usuarios.
// ============================================================

// Package auth contiene los handlers para la autenticación.
package auth

import (
	"api-auth/internal/handler/auth/dto/request"
	authService "api-auth/internal/service/auth"
	loginServiceDto "api-auth/internal/service/auth/dto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Handler representa el controlador encargado de manejar
// las solicitudes relacionadas con autenticación.
type Handler struct {
	service authService.ServiceInterface
}

// NewAuthHandler crea una nueva instancia de AuthHandler.
//
// Parámetros:
//   - s: implementación de AuthServiceInterface.
//
// Retorna:
//   - *AuthHandler: instancia inicializada.
//
// Errores:
//   - No aplica.
func NewAuthHandler(s authService.ServiceInterface) *Handler {
	return &Handler{service: s}
}

// Login maneja el proceso de autenticación
// @Summary Iniciar sesión de usuario
// @Description Autentica un usuario mediante email y contraseña
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body request.LoginRequestDto true "Credenciales de acceso"
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /v1/auth/login [post]
func (h *Handler) Login(c *gin.Context) {
	var req request.LoginRequestDto

	if err := c.ShouldBindJSON(&req); err != nil {
		// Guardamos error en response y dejamos que el middleware lo formatee
		c.Set("response_error", map[string]interface{}{
			"message":   err.Error(),
			"errorCode": strconv.Itoa(http.StatusBadGateway),
			"httpCode":  http.StatusBadGateway,
		})
		c.Abort()
		return
	}

	loginDto := &loginServiceDto.LoginServiceDto{
		Email:    req.Email,
		Password: req.Password,
	}

	userResp, err := h.service.Login(loginDto)
	if err != nil {
		c.Set("response_error", map[string]interface{}{
			"message":   err.Error(),
			"errorCode": strconv.Itoa(http.StatusInternalServerError),
			"httpCode":  http.StatusInternalServerError,
		})
		c.Abort()
		return
	}

	// Guardamos el response para que el middleware lo envuelva
	c.Set("response", userResp)
}
