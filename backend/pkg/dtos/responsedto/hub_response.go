package responsedto

import "time"

type OrganizationRecord struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Owner     *UserData `json:"owner,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type OrganizationPaginateResponse struct {
	Data     []OrganizationRecord `json:"data"`
	Metadata PaginateMetaData     `json:"metadata"`
}
