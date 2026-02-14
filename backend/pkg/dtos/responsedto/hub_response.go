package responsedto

import "time"

type OrganizationRecord struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Owner     *UserData `json:"owner,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type OrganizationPaginateResponse struct {
	Data     []OrganizationRecord `json:"data"`
	Metadata PaginateMetaData     `json:"metadata"`
}
