package response

type UserServiceResponseDto struct {
	ID        int64   `json:"id"`
	Username  string  `json:"username"`
	Email     string  `json:"email"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Phone     *string `json:"phone,omitempty"`
	CountryID int32   `json:"country_id"`
	Address   *string `json:"address_line,omitempty"`
	Token     string  `json:"token"`
}
