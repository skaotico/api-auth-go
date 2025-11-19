package auth

import (
	config "api-auth/pkg/platform/bd"
	"database/sql"
)

type postgresAuthRepository struct {
	db *sql.DB
}

func NewAuthRepository() AuthRepository {
	return &postgresAuthRepository{
		db: config.DB,
	}
}
