package api

type UserRequestRegister struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRequestLogin struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	UserID    int    `json:"user_id"`
	Token     string `json:"token"`
	ExpiresIn int64  `json:"expires_in"`
}
