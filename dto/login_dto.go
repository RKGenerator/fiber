package dto

type AuthRequest struct {
	Email    string `query:"email"`
	Password string `query:"password"`
}

type AuthResponse struct {
	AccessToken string `json:"access"`
}
