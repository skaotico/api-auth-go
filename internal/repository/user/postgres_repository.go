// ============================================================
// @file: postgres_repository.go
// @author: Yosemar Andrade
// @date: 2025-11-26
// @lastModified: 2025-11-26
// @description: Implementación del repositorio de usuarios para PostgreSQL.
// ============================================================

package user

import (
	"api-auth/internal/domain/user"
	"api-auth/pkg/logger"
	config "api-auth/pkg/platform/bd"
	"database/sql"
	"errors"

	"go.uber.org/zap"
)

type postgresUserRepository struct {
	db *sql.DB
}

// NewUserRepository crea una nueva instancia del repositorio de usuarios.
//
// Parámetros:
//   - No recibe parámetros.
//
// Retorna:
//   - UserRepository: interfaz del repositorio de usuarios.
//
// Errores:
//   - No retorna errores.
func NewUserRepository() UserRepository {
	return &postgresUserRepository{
		db: config.DB,
	}
}

// FindByEmail busca un usuario por su correo electrónico.
//
// Parámetros:
//   - email: correo electrónico del usuario.
//
// Retorna:
//   - *user.User: el usuario encontrado.
//   - error: error si no se encuentra o hay fallo en BD.
//
// Errores:
//   - Retorna `user not found` si no existe.
//   - Retorna error de BD si falla la consulta.
func (r *postgresUserRepository) FindByEmail(email string) (*user.User, error) {
	var userFind user.User

	query := `SELECT
		id,
		username,
		email,
		password_hash,
		first_name,
		last_name,
		phone,
		birth_date,
		is_active,
		country_id,
		address_line,
		created_at,
		updated_at,
		deleted_at
		FROM users WHERE email = $1`

	logger.Log.Debug("Ejecutando consulta SQL", zap.String("query", query), zap.String("email", email))

	row := r.db.QueryRow(query, email)

	err := row.Scan(
		&userFind.ID,
		&userFind.Username,
		&userFind.Email,
		&userFind.PasswordHash,
		&userFind.FirstName,
		&userFind.LastName,
		&userFind.Phone,
		&userFind.BirthDate,
		&userFind.IsActive,
		&userFind.CountryID,
		&userFind.AddressLine,
		&userFind.CreatedAt,
		&userFind.UpdatedAt,
		&userFind.DeletedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			logger.Log.Warn("Usuario no encontrado", zap.String("email", email))
			return nil, errors.New("user not found")
		}
		logger.Log.Error("Error al buscar usuario por email", zap.Error(err))
		return nil, err
	}

	return &userFind, nil
}

// FindByID busca un usuario por su ID.
//
// Parámetros:
//   - id: identificador del usuario.
//
// Retorna:
//   - *user.User: el usuario encontrado.
//   - error: error si no se encuentra o hay fallo en BD.
func (r *postgresUserRepository) FindByID(id int) (*user.User, error) {
	var userFind user.User

	query := `SELECT
		id,
		username,
		email,
		password_hash,
		first_name,
		last_name,
		phone,
		birth_date,
		is_active,
		country_id,
		address_line,
		created_at,
		updated_at,
		deleted_at
		FROM users WHERE id = $1`

	logger.Log.Debug("Ejecutando consulta SQL", zap.String("query", query), zap.Int("id", id))

	row := r.db.QueryRow(query, id)

	err := row.Scan(
		&userFind.ID,
		&userFind.Username,
		&userFind.Email,
		&userFind.PasswordHash,
		&userFind.FirstName,
		&userFind.LastName,
		&userFind.Phone,
		&userFind.BirthDate,
		&userFind.IsActive,
		&userFind.CountryID,
		&userFind.AddressLine,
		&userFind.CreatedAt,
		&userFind.UpdatedAt,
		&userFind.DeletedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			logger.Log.Warn("Usuario no encontrado", zap.Int("id", id))
			return nil, errors.New("user not found")
		}
		logger.Log.Error("Error al buscar usuario por id", zap.Error(err))
		return nil, err
	}

	return &userFind, nil
}

// FindAll lista todos los usuarios.
//
// Parámetros:
//   - No recibe parámetros.
//
// Retorna:
//   - []*user.User: lista de usuarios.
//   - error: error si falla la consulta.
//
// Errores:
//   - Retorna error de BD si falla la consulta.
func (r *postgresUserRepository) FindAll() ([]*user.User, error) {
	query := `
        SELECT 
            id,
            username,
            email,
            password_hash,
            first_name,
            last_name,
            phone,
            birth_date,
            is_active,
            country_id,
            address_line,
            created_at,
            updated_at,
            deleted_at
        FROM public.users
    `
	logger.Log.Debug("Ejecutando consulta SQL FindAll", zap.String("query", query))

	rows, err := r.db.Query(query)
	if err != nil {
		logger.Log.Error("Error al listar usuarios", zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	var users []*user.User
	for rows.Next() {
		var u user.User
		if err := rows.Scan(
			&u.ID,
			&u.Username,
			&u.Email,
			&u.PasswordHash,
			&u.FirstName,
			&u.LastName,
			&u.Phone,
			&u.BirthDate,
			&u.IsActive,
			&u.CountryID,
			&u.AddressLine,
			&u.CreatedAt,
			&u.UpdatedAt,
			&u.DeletedAt,
		); err != nil {
			logger.Log.Error("Error al escanear usuario", zap.Error(err))
			return nil, err
		}
		users = append(users, &u)
	}

	return users, nil
}

// Save guarda un nuevo usuario en la base de datos.
//
// Parámetros:
//   - u: puntero al usuario a guardar.
//
// Retorna:
//   - error: error si falla la inserción.
//
// Errores:
//   - Retorna error de BD si falla la inserción.
func (r *postgresUserRepository) Save(u *user.User) error {
	query := `
	INSERT INTO users (
		username,
		first_name,
		last_name,
		email,
		password_hash,
		phone,
		birth_date,
		is_active,
		country_id,
		address_line
	) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)
	RETURNING id, created_at, updated_at
	`
	logger.Log.Debug("Ejecutando consulta SQL Save", zap.String("query", query), zap.String("username", u.Username))

	err := r.db.QueryRow(
		query,
		u.Username,
		u.FirstName,
		u.LastName,
		u.Email,
		u.PasswordHash,
		u.Phone,
		u.BirthDate,
		u.IsActive,
		u.CountryID,
		u.AddressLine,
	).Scan(&u.ID, &u.CreatedAt, &u.UpdatedAt)

	if err != nil {
		logger.Log.Error("Error al guardar usuario", zap.Error(err))
		return err
	}
	return nil
}
