package services

import (
	"DewaSRY/sociomile-app/pkg/dtos/responsedto"
)

type OrganizationCrudService interface {
	GetAllOrganizations() (*responsedto.OrganizationListResponse, error)
}
