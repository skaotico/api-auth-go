package user

import (
	"api-auth/internal/domain/user"
	config "api-auth/pkg/platform/bd"
	"database/sql"
	"errors"
	"log"
)

type postgresUserRepository struct {
	db *sql.DB
}

// Constructor que devuelve un UserRepository
func NewUserRepository() UserRepository {
	return &postgresUserRepository{
		db: config.DB,
	}
}

// Buscar usuario por email
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

	log.Printf("SQL -> %s  | args: [%s]", query, email)

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
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &userFind, nil
}

// Listar todos los usuarios
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
	rows, err := r.db.Query(query)
	if err != nil {
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
			return nil, err
		}
		users = append(users, &u)
	}

	return users, nil
}

// Guardar un usuario nuevo
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
		return err
	}
	return nil
}
