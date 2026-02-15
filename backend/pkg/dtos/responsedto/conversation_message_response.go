package responsedto

import "time"

type ConversationMessageResponse struct {
	ID             uint                  `json:"id"`
	OrganizationID uint                  `json:"organizationId"`
	ConversationID uint                  `json:"conversationId"`
	Conversation   *ConversationResponse `json:"conversation,omitempty"`
	CreatedByID    uint                  `json:"createdById"`
	CreatedBy      *UserData             `json:"createdBy,omitempty"`
	Message        string                `json:"message"`
	CreatedAt      time.Time             `json:"createdAt"`
	UpdatedAt      time.Time             `json:"updatedAt"`
}

type ConversationMessageListResponse struct {
	Messages []ConversationMessageResponse `json:"messages"`
	Metadata PaginateMetaData              `json:"metadata"`
}

type ConversationMessagePaginateResponse struct {
	Data     []ConversationMessageResponse `json:"data"`
	Metadata PaginateMetaData              `json:"metadata"`
}
