package responsedto

import "time"

type ConversationResponse struct {
	ID                  uint                  `json:"id"`
	OrganizationID      uint                  `json:"organizationId"`
	Organization        *OrganizationResponse `json:"organization,omitempty"`
	GuestID             uint                  `json:"guestId"`
	Guest               *UserData             `json:"guest,omitempty"`
	OrganizationStaffID *uint                 `json:"organizationStaffId,omitempty"`
	OrganizationStaff   *UserData             `json:"organizationStaff,omitempty"`
	Status              string                `json:"status"`
	CreatedAt           time.Time             `json:"createdAt"`
	UpdatedAt           time.Time             `json:"updatedAt"`
}

type ConversationListResponse struct {
	Conversations []ConversationResponse `json:"conversations"`
	Metadata      PaginateMetaData       `json:"metadata"`
}

type ConversationListPaginateResponse struct {
	Data     []ConversationResponse `json:"data"`
	Metadata PaginateMetaData       `json:"metadata"`
}
