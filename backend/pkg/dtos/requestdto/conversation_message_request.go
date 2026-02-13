package requestdto

type CreateConversationMessageRequest struct {
	ConversationID uint   `json:"conversation_id" validate:"required"`
	Message        string `json:"message" validate:"required,min=1,max=5000"`
}
