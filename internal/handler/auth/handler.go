// ============================================================
// @file: auth_handler.go
// @author: Yosemar Andrade
// @date: 2025-11-18
// @lastModified: 2025-11-18
// @description: Handler para autenticación de usuarios.
// ============================================================

package auth

import (
	"api-auth/internal/handler/auth/dto/request"
	service "api-auth/internal/service/auth"
	loginServiceDto "api-auth/internal/service/auth/dto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AuthHandler representa el controlador encargado de manejar
// las solicitudes relacionadas con autenticación.
type AuthHandler struct {
	service service.AuthServiceInterface
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
func NewAuthHandler(s service.AuthServiceInterface) *AuthHandler {
	return &AuthHandler{service: s}
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
func (h *AuthHandler) Login(c *gin.Context) {
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
