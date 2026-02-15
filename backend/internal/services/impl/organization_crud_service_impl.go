package impl

import (
	"DewaSRY/sociomile-app/internal/services"
	"DewaSRY/sociomile-app/pkg/dtos/responsedto"
	"DewaSRY/sociomile-app/pkg/models"
	"errors"

	"gorm.io/gorm"
)

type organizationCrudServiceImpl struct {
	db *gorm.DB
}


// GetAllOrganizations implements services.OrganizationService.
func (t *organizationCrudServiceImpl) GetAllOrganizations() (*responsedto.OrganizationListResponse, error) {
	var organizations []models.OrganizationModel
	if err := t.db.Preload("Owner").Find(&organizations).Error; err != nil {
		return nil, errors.New("failed to fetch organizations")
	}

	var organizationResponses []responsedto.OrganizationResponse
	for _, org := range organizations {
		organizationResponses = append(organizationResponses, *t.mapToOrganizationResponse(&org))
	}

	return &responsedto.OrganizationListResponse{
		Organizations: organizationResponses,
		Metadata: responsedto.PaginateMetaData{
			Total: len(organizationResponses),
			Page:  1,
			Limit: len(organizationResponses),
		},
	}, nil
}



func (t *organizationCrudServiceImpl) mapToOrganizationResponse(org *models.OrganizationModel) *responsedto.OrganizationResponse {
	response := &responsedto.OrganizationResponse{
		ID:        org.ID,
		Name:      org.Name,
		OwnerID:   org.OwnerID,
		CreatedAt: org.CreatedAt,
		UpdatedAt: org.UpdatedAt,
	}

	if org.Owner != nil {
		response.Owner = &responsedto.UserData{
			ID:    org.Owner.ID,
			Email: org.Owner.Email,
			Name:  org.Owner.Name,
		}
	}

	return response
}

func NewOrganizationCrudService(db *gorm.DB) services.OrganizationCrudService {
	return &organizationCrudServiceImpl{db:db}
}
