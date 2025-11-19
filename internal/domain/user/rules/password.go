package rules

import (
	"api-auth/internal/domain/user"
	"strings"
)

func ValidatePasswordNotEmpty(pass string) error {
	if strings.TrimSpace(pass) == "" {
		return user.ErrInvalidPassword
	}
	return nil
}
