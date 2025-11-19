package rules

import (
	"api-auth/internal/domain/user"
	"strings"
)

func ValidateEmail(email string) error {
	if !strings.Contains(email, "@") {
		return user.ErrInvalidEmail
	}
	return nil
}
