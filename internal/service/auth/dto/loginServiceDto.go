// ============================================================
// @file: loginServiceDto.go
// @author: Yosemar Andrade
// @created: 2025-11-20
// @description: DTO utilizado para recibir credenciales de autenticación en el servicio.
// ============================================================

// Package dto contiene los objetos de transferencia de datos para el servicio de autenticación.
package dto

// LoginServiceDto representa las credenciales necesarias para autenticar
// a un usuario dentro del sistema.
type LoginServiceDto struct {
	Email    string // Email del usuario que solicita autenticación.
	Password string // Contraseña asociada al usuario.
}
