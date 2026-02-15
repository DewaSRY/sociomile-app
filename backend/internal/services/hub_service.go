package services

import (
	"DewaSRY/sociomile-app/pkg/dtos/filtersdto"
	"DewaSRY/sociomile-app/pkg/dtos/requestdto"
	"DewaSRY/sociomile-app/pkg/dtos/responsedto"
)

type HubService interface {
	CreateOrganization(requestdto.RegisterOrganizationRequest) error
	GetOrganizationPagination(filtersdto.FiltersDto) (*responsedto.OrganizationPaginateResponse, error)
}
