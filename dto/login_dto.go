package dto

type AuthRequest struct {
	User     string `query:"user"`
	Password string `query:"password"`
}
