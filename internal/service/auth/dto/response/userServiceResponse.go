// ============================================================
// @file: userServiceResponse.go
// @author: Yosemar Andrade
// @created: 2025-11-20
// @description: DTO de respuesta utilizado para retornar información de usuario desde el servicio.
// ============================================================

// Package response contiene los DTOs de respuesta del servicio de autenticación.
package response

// UserServiceResponseDto representa la estructura de información entregada
// por el servicio de usuarios al momento de autenticar o recuperar datos
// del usuario.
//
// Este DTO incluye información personal y el token de acceso generado.
type UserServiceResponseDto struct {
	ID        int64   `json:"id"`
	Username  string  `json:"username"`
	Email     string  `json:"email"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Phone     *string `json:"phone,omitempty"`
	CountryID int32   `json:"country_id"`
	Address   *string `json:"address_line,omitempty"`
	Token     string  `json:"token"`
}
