package auth

type Auth struct {
	// Domain model for authentication
	ID    uint   `json:"id"`
	Token string `json:"token"`
}
