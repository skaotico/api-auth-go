package mapper

import (
	domain "api-auth/internal/domain/user"
	resp "api-auth/internal/service/auth/dto/response"
)

func MapUserToResponse(u *domain.User, token string) *resp.UserServiceResponseDto {
	return &resp.UserServiceResponseDto{
		ID:        u.ID,
		Username:  u.Username,
		Email:     u.Email,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Phone:     u.Phone,
		CountryID: u.CountryID,
		Address:   u.AddressLine,
		Token:     token,
	}
}
