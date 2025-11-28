package user

import domain "api-auth/internal/domain/user"

type UserService interface {
	GetAllUsers() ([]*domain.User, error)
	GetUserByEmail(email string) (*domain.User, error)
	GetUserByID(id int) (*domain.User, error)
	Login(email, password string) (*domain.User, error)
	CreateUser(u *domain.User, plainPassword string) error
}
