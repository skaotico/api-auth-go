package request

type LoginRequestDto struct {
	Email    string `json:"email" binding:"required,email,trim"`
	Password string `json:"password" binding:"required,min=8,max=20,trim,regexp=^[a-zA-Z0-9_.@-]*$"`
}
