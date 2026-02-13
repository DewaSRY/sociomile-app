package responsedto

type AuthResponse struct {
	Token string   `json:"token"`
	User  UserData `json:"user"`
}

type UserData struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

