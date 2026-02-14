package responsedto

import "time"

type TicketResponse struct {
	ID             uint                  `json:"id"`
	OrganizationID uint                  `json:"organization_id"`
	Organization   *OrganizationResponse `json:"organization,omitempty"`
	ConversationID uint                  `json:"conversation_id"`
	Conversation   *ConversationResponse `json:"conversation,omitempty"`
	CreatedByID    uint                  `json:"created_by_id"`
	CreatedBy      *UserData             `json:"created_by,omitempty"`
	TicketNumber   string                `json:"ticket_number"`
	Name           string                `json:"name"`
	Status         string                `json:"status"`
	CreatedAt      time.Time             `json:"created_at"`
	UpdatedAt      time.Time             `json:"updated_at"`
}

type TicketListResponse struct {
	Tickets  []TicketResponse `json:"tickets"`
	Metadata PaginateMetaData `json:"metadata"`
}
