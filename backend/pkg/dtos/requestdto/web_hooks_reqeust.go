package requestdto



type WebHooksRequest struct{
	OrganizationID uint `json:"organizationId" validate:"required"`
	Message        string `json:"message" validate:"required,min=1,max=5000"`
	Email    string `json:"email" validate:"required,email"`
}