// ============================================================
// @file: user_service.go
// @author: Yosemar Andrade
// @date: 2025-11-18
// @lastModified: 2025-11-26
// @description: Implementación del servicio de usuarios, encargado de
// manejar la lógica de negocio relacionada con usuarios, incluyendo
// creación, obtención, autenticación y validaciones.
// ============================================================

package impl

import (
	"errors"

	domain "api-auth/internal/domain/user"
	repo "api-auth/internal/repository/user"
	"api-auth/internal/service/user"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

// UserServiceImpl representa la implementación concreta del servicio de usuarios.
// Este servicio encapsula las operaciones de negocio y delega persistencia al repositorio.
type UserServiceImpl struct {
	repo repo.UserRepository
	log  *zap.Logger
}

// NewUserService crea una nueva instancia de UserService.
//
// Parámetros:
//   - r: repositorio de usuarios.
//   - logger: instancia del logger para trazabilidad.
//
// Retorna:
//   - Una nueva implementación de UserService.
func NewUserService(r repo.UserRepository, logger *zap.Logger) user.UserService {
	logger.Info("Inicializando UserService")
	return &UserServiceImpl{repo: r, log: logger}
}

// GetAllUsers obtiene todos los usuarios registrados.
//
// Retorna:
//   - Lista de usuarios.
//   - Error si ocurre algún problema al obtener los datos.
func (s *UserServiceImpl) GetAllUsers() ([]*domain.User, error) {
	s.log.Info("Solicitando listado completo de usuarios")

	users, err := s.repo.FindAll()
	if err != nil {
		s.log.Error("Error al obtener usuarios", zap.Error(err))
		return nil, err
	}

	s.log.Info("Usuarios obtenidos correctamente", zap.Int("total", len(users)))
	return users, nil
}

// GetUserByEmail obtiene un usuario según su email.
//
// Parámetros:
//   - email: correo del usuario.
//
// Retorna:
//   - Usuario encontrado o error si no existe.
func (s *UserServiceImpl) GetUserByEmail(email string) (*domain.User, error) {
	s.log.Info("Buscando usuario por email", zap.String("email", email))

	user, err := s.repo.FindByEmail(email)
	if err != nil {
		s.log.Warn("Usuario no encontrado", zap.String("email", email))
		return nil, domain.ErrUserNotFound
	}

	s.log.Info("Usuario encontrado",
		zap.Int("id", user.ID),
		zap.String("email", user.Email),
	)
	return user, nil
}

// GetUserByID obtiene un usuario según su ID.
//
// Parámetros:
//   - id: identificador del usuario.
//
// Retorna:
//   - Usuario encontrado o error si no existe.
func (s *UserServiceImpl) GetUserByID(id int) (*domain.User, error) {
	s.log.Info("Buscando usuario por ID", zap.Int("id", id))

	user, err := s.repo.FindByID(id)
	if err != nil {
		s.log.Warn("Usuario no encontrado", zap.Int("id", id))
		return nil, domain.ErrUserNotFound
	}

	s.log.Info("Usuario encontrado",
		zap.Int("id", user.ID),
		zap.String("email", user.Email),
	)
	return user, nil
}

// Login valida las credenciales de un usuario.
//
// Parámetros:
//   - email: correo del usuario.
//   - password: contraseña en texto plano.
//
// Retorna:
//   - Usuario autenticado o error si las credenciales son inválidas.
func (s *UserServiceImpl) Login(email, password string) (*domain.User, error) {
	s.log.Info("Intentando autenticar usuario", zap.String("email", email))

	user, err := s.repo.FindByEmail(email)
	if err != nil {
		s.log.Warn("Usuario no encontrado", zap.String("email", email))
		return nil, domain.ErrUserNotFound
	}

	if bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)) != nil {
		s.log.Warn("Contraseña incorrecta", zap.String("email", email))
		return nil, domain.ErrInvalidPassword
	}

	s.log.Info("Autenticación exitosa",
		zap.Int("id", user.ID),
		zap.String("email", user.Email),
	)

	return user, nil
}

// CreateUser crea un nuevo usuario generando su hash de contraseña.
//
// Parámetros:
//   - u: estructura del usuario.
//   - plainPassword: contraseña sin encriptar.
//
// Retorna:
//   - Error si ocurre algún problema en la creación.
func (s *UserServiceImpl) CreateUser(u *domain.User, plainPassword string) error {
	s.log.Info("Intentando crear un nuevo usuario")

	if u == nil {
		s.log.Error("El usuario recibido es nulo")
		return errors.New("user is nil")
	}

	s.log.Info("Generando hash de contraseña", zap.String("email", u.Email))
	hash, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	if err != nil {
		s.log.Error("Error al generar hash de contraseña", zap.Error(err))
		return err
	}

	u.PasswordHash = string(hash)

	s.log.Info("Guardando usuario", zap.String("email", u.Email))
	if err := s.repo.Save(u); err != nil {
		s.log.Error("Error al guardar usuario", zap.Error(err))
		return err
	}

	s.log.Info("Usuario creado exitosamente",
		zap.Int("id", u.ID),
		zap.String("email", u.Email),
	)

	return nil
}
