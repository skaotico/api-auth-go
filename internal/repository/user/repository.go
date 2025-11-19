package user

import (
	domain "api-auth/internal/domain/user"
)

type UserRepository interface {
	FindByEmail(email string) (*domain.User, error)
	FindAll() ([]*domain.User, error)
	Save(user *domain.User) error
}
