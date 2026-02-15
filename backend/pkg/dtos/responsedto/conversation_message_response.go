package responsedto

import "time"

type ConversationMessageResponse struct {
	ID             uint                  `json:"id"`
	OrganizationID uint                  `json:"organization_id"`
	ConversationID uint                  `json:"conversation_id"`
	Conversation   *ConversationResponse `json:"conversation,omitempty"`
	CreatedByID    uint                  `json:"created_by_id"`
	CreatedBy      *UserData             `json:"created_by,omitempty"`
	Message        string                `json:"message"`
	CreatedAt      time.Time             `json:"created_at"`
	UpdatedAt      time.Time             `json:"updated_at"`
}

type ConversationMessageListResponse struct {
	Messages []ConversationMessageResponse `json:"messages"`
	Metadata PaginateMetaData              `json:"metadata"`
}

type ConversationMessagePaginateResponse struct {
	Data     []ConversationMessageResponse `json:"data"`
	Metadata PaginateMetaData              `json:"metadata"`
}
