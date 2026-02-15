package requestdto

type CreateConversationRequest struct {
	OrganizationID uint `json:"organizationId" validate:"required"`
}

type UpdateConversationRequest struct {
	Status string `json:"status" validate:"required,oneof=pending in_progress done"`
}

type AssignConversationRequest struct {
	OrganizationStaffID uint `json:"organizationStaffId" validate:"required"`
}
