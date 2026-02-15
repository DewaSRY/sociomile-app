package responsedto

import "time"

type OrganizationResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	OwnerID   uint      `json:"owner_id"`
	Owner     *UserData `json:"owner,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type OrganizationListResponse struct {
	Organizations []OrganizationResponse `json:"organizations"`
	Metadata      PaginateMetaData       `json:"metadata"`
}

// New
type OrganizationStaffRecord struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	RoleName  string `json:"roleName"`
	Email string `json:"email"`
}

type OrganizationStaffPagination struct {
	Data     []OrganizationStaffRecord `json:"data"`
	Metadata PaginateMetaData       `json:"metadata"`
}
