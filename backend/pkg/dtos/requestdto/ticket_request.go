package requestdto

type CreateTicketRequest struct {
	ConversationID uint   `json:"conversation_id" validate:"required"`
	Name           string `json:"name" validate:"required,min=3,max=200"`
}

type UpdateTicketRequest struct {
	Name   string `json:"name,omitempty" validate:"omitempty,min=3,max=200"`
	Status string `json:"status,omitempty" validate:"omitempty,oneof=pending in_progress done"`
}
