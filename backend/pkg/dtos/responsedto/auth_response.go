package responsedto

type AuthResponse struct {
	Token string   `json:"token"`
	User  UserProfileData `json:"user"`
}


type UserProfileData struct {
	ID           uint                  `json:"id"`
	Email        string                `json:"email"`
	Name         string                `json:"name"`
	RoleName     string                `json:"roleName"`
	Organization *OrganizationResponse `json:"organization,omitempty"`
}


type UserData struct {
	ID           uint                  `json:"id"`
	Email        string                `json:"email"`
	Name         string                `json:"name"`
}
