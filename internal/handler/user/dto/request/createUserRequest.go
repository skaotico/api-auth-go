package request

import "time"

type CreateUserRequest struct {
	Username    string     `json:"username" binding:"required" example:"yandrade"`
	FirstName   string     `json:"first_name" example:"Yosemar"`
	LastName    string     `json:"last_name" example:"Andrade"`
	Email       string     `json:"email" binding:"required,email" example:"user@example.com"`
	Password    string     `json:"password" binding:"required" example:"secret123"`
	Phone       string     `json:"phone" example:"+56912345678"`
	BirthDate   *time.Time `json:"birth_date,omitempty" example:"1990-01-01T00:00:00Z"`
	CountryID   int32      `json:"country_id" example:"56"`
	AddressLine string     `json:"address_line" example:"Calle Falsa 123"`
}
