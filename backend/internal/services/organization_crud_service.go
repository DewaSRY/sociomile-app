package services

import (
	"DewaSRY/sociomile-app/pkg/dtos/requestdto"
	"DewaSRY/sociomile-app/pkg/dtos/responsedto"
	"DewaSRY/sociomile-app/pkg/models"
)

type OrganizationCrudService interface {
	CreateOrganization(req requestdto.CreateOrganizationRequest) (*responsedto.OrganizationResponse, error)
	GetOrganizationByID(id uint) (*responsedto.OrganizationResponse, error)
	GetAllOrganizations() (*responsedto.OrganizationListResponse, error)
	UpdateOrganization(id uint, req requestdto.UpdateOrganizationRequest) (*responsedto.OrganizationResponse, error)
	DeleteOrganization(id uint) error
	GetOrganizationStats(organizationID uint) (map[string]interface{}, error)
	CreateOwnerUser(email, name, password string) (*models.UserModel, error)
}
