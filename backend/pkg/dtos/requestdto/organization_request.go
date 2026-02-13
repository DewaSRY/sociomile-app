package requestdto

type CreateOrganizationRequest struct {
	Name    string `json:"name" validate:"required,min=3,max=100"`
	OwnerID uint   `json:"owner_id" validate:"required"`
}

type UpdateOrganizationRequest struct {
	Name string `json:"name" validate:"required,min=3,max=100"`
}
