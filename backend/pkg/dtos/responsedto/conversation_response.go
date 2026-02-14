package responsedto

import "time"

type ConversationResponse struct {
	ID                  uint                  `json:"id"`
	OrganizationID      uint                  `json:"organization_id"`
	Organization        *OrganizationResponse `json:"organization,omitempty"`
	GuestID             uint                  `json:"guest_id"`
	Guest               *UserData             `json:"guest,omitempty"`
	OrganizationStaffID *uint                 `json:"organization_staff_id,omitempty"`
	OrganizationStaff   *UserData             `json:"organization_staff,omitempty"`
	Status              string                `json:"status"`
	CreatedAt           time.Time             `json:"created_at"`
	UpdatedAt           time.Time             `json:"updated_at"`
}

type ConversationListResponse struct {
	Conversations []ConversationResponse `json:"conversations"`
	Metadata      PaginateMetaData       `json:"metadata"`
}
