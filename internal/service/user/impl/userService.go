// ============================================================
// @file: userService.go
// @author: Yosemar Andrade
// @created: 2025-11-20
// @description: Implementación del servicio de usuarios.
// ============================================================

// Package impl implementa la lógica de negocio del servicio de usuarios.
package impl

import (
	"errors"
	"log"

	domain "api-auth/internal/domain/user"
	userRepo "api-auth/internal/repository/user"
	userInterface "api-auth/internal/service/user"

	"golang.org/x/crypto/bcrypt"
)

// UserService representa la implementación concreta del servicio de usuarios.
// Este servicio encapsula las operaciones de negocio y delega persistencia al repositorio.
type UserService struct {
	repo userRepo.Repository
}

// Verifica en tiempo de compilación que UserService implementa la interfaz UserServiceInterface.
var _ userInterface.ServiceInterface = (*UserService)(nil)

// NewUserService crea una nueva instancia de UserService.
func NewUserService(r userRepo.Repository) *UserService {
	log.Printf("Inicializando UserService")
	return &UserService{repo: r}
}

// GetAllUsers obtiene todos los usuarios registrados.
func (s *UserService) GetAllUsers() ([]*domain.User, error) {
	log.Printf("Solicitando listado completo de usuarios")

	users, err := s.repo.FindAll()
	if err != nil {
		log.Printf("Error al obtener usuarios: %v", err)
		return nil, err
	}

	log.Printf("Usuarios obtenidos correctamente. Total: %d", len(users))
	return users, nil
}

// GetUserByEmail obtiene un usuario según su email.
func (s *UserService) GetUserByEmail(email string) (*domain.User, error) {
	log.Printf("Buscando usuario por email: %s", email)

	user, err := s.repo.FindByEmail(email)
	if err != nil {
		log.Printf("No se encontró usuario con email: %s", email)
		return nil, domain.ErrUserNotFound
	}

	log.Printf("Usuario encontrado: ID=%d, Email=%s", user.ID, user.Email)
	return user, nil
}

// Login valida las credenciales de un usuario mediante email y contraseña.
func (s *UserService) Login(email, password string) (*domain.User, error) {
	log.Printf("Intentando autenticar usuario con email: %s", email)

	user, err := s.repo.FindByEmail(email)
	if err != nil {
		log.Printf("Usuario no encontrado para email: %s", email)
		return nil, domain.ErrUserNotFound
	}

	if bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)) != nil {
		log.Printf("Contraseña incorrecta para el usuario con email: %s", email)
		return nil, domain.ErrInvalidPassword
	}

	log.Printf("Autenticación exitosa para usuario ID=%d", user.ID)
	return user, nil
}

// CreateUser crea un nuevo usuario generando automáticamente el hash de su contraseña.
func (s *UserService) CreateUser(u *domain.User, plainPassword string) error {
	log.Printf("Intentando crear un nuevo usuario")

	if u == nil {
		log.Printf("Error: el usuario recibido es nulo")
		return errors.New("user is nil")
	}

	log.Printf("Generando hash de contraseña para el usuario con email: %s", u.Email)
	hash, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error al generar el hash de la contraseña: %v", err)
		return err
	}

	u.PasswordHash = string(hash)

	log.Printf("Guardando usuario en el repositorio: Email=%s", u.Email)
	if err := s.repo.Save(u); err != nil {
		log.Printf("Error al guardar usuario: %v", err)
		return err
	}

	log.Printf("Usuario creado exitosamente: ID=%d, Email=%s", u.ID, u.Email)
	return nil
}
