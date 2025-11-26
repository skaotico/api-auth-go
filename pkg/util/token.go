package utils

import (
	"crypto/rand"
	"encoding/base64"
)

// NewRefreshToken genera un refresh token seguro, aleatorio y opaco.
func NewRefreshToken() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}
