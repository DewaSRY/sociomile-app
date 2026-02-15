package responsedto

import "time"

type TicketResponse struct {
	ID             uint                  `json:"id"`
	OrganizationID uint                  `json:"organizationId"`
	Organization   *OrganizationResponse `json:"organization,omitempty"`
	ConversationID uint                  `json:"conversationId"`
	Conversation   *ConversationResponse `json:"conversation,omitempty"`
	CreatedByID    uint                  `json:"createdById"`
	CreatedBy      *UserData             `json:"createdBy,omitempty"`
	TicketNumber   string                `json:"ticketNumber"`
	Name           string                `json:"name"`
	Status         string                `json:"status"`
	CreatedAt      time.Time             `json:"createdAt"`
	UpdatedAt      time.Time             `json:"updatedAt"`
}

type TicketListResponse struct {
	Tickets  []TicketResponse `json:"tickets"`
	Metadata PaginateMetaData `json:"metadata"`
}
