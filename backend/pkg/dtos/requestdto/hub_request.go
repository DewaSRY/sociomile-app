package requestdto

type RegisterOrganizationRequest struct {
	Name      string `json:"name" validate:"required,min=3,max=100"`
	Email     string `json:"email" validate:"required,email"`
	OwnerName string `json:"ownerName" validate:"required,min=3,max=100"`
	Password  string `json:"password" validate:"required,min=6"`
}

type PutOrganizationRequest struct {
	Name string `json:"name" validate:"required,min=3,max=100"`
}
