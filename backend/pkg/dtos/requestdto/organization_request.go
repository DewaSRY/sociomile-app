package requestdto

type CreateOrganizationRequest struct {
	Name    string `json:"name" validate:"required,min=3,max=100"`
	OwnerID uint   `json:"ownerId" validate:"required"`
}

type UpdateOrganizationRequest struct {
	Name string `json:"name" validate:"required,min=3,max=100"`
}
