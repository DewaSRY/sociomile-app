package impl

import (
	"DewaSRY/sociomile-app/internal/services"
	"DewaSRY/sociomile-app/pkg/dtos/filtersdto"
	"DewaSRY/sociomile-app/pkg/dtos/requestdto"
	"DewaSRY/sociomile-app/pkg/dtos/responsedto"
	"DewaSRY/sociomile-app/pkg/models"
	"errors"

	"gorm.io/gorm"
)

type hubServiceImpl struct {
	db *gorm.DB
}

// CreateOrganization implements services.HubService.
func (t *hubServiceImpl) CreateOrganization(req requestdto.RegisterOrganizationRequest) error {
	return t.db.Transaction(func(tx *gorm.DB) error {

		var existingUser models.UserModel
		if err := tx.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
			return errors.New("user with this email already exists")
		}

		var ownerRole models.UserRoleModel
		if err := tx.Where("name = ?", models.RoleOrganizationOwner).
			First(&ownerRole).Error; err != nil {
			return errors.New("organization owner role not found")
		}

		userOwner := models.UserModel{
			Email:    req.Email,
			Name:     req.OwnerName,
			Password: req.Password,
			RoleID:   ownerRole.ID,
		}

		if err := tx.Create(&userOwner).Error; err != nil {
			return errors.New("failed to create user")
		}

		organization := models.OrganizationModel{
			Name:    req.Name,
			OwnerID: userOwner.ID,
		}

		if err := tx.Create(&organization).Error; err != nil {
			return errors.New("failed to create organization")
		}

		if err := tx.Model(&userOwner).Update("organization_id", organization.ID).Error; err != nil {
			return errors.New("failed to update owner's organization")
		}

		return nil
	})
}

// GetOrganizationPagination implements services.HubService.
func (t *hubServiceImpl) GetOrganizationPagination(filter filtersdto.FiltersDto) (*responsedto.OrganizationPaginateResponse, error) {
	var organizations []models.OrganizationModel
	var total int64
	offset := (*filter.Page - 1) * *filter.Limit

	if err := t.db.Model(&models.OrganizationModel{}).
		Count(&total).Error; err != nil {
		return nil, errors.New("failed to count organizations")
	}

	if err := t.db.Preload("Owner").Limit(*filter.Limit).
		Offset(offset).
		Find(&organizations).Error; err != nil {
		return nil, errors.New("failed to fetch organizations")
	}

	organizationResponses := make([]responsedto.HubOrganizationRecord, 0, len(organizations))

	for _, org := range organizations {
		organizationResponses = append(
			organizationResponses,
			*t.mapToOrganizationResponse(&org),
		)
	}

	return &responsedto.OrganizationPaginateResponse{
		Data: organizationResponses,
		Metadata: responsedto.PaginateMetaData{
			Total: int(total),
			Page:  *filter.Page,
			Limit: *filter.Limit,
		},
	}, nil
}

func (t *hubServiceImpl) mapToOrganizationResponse(org *models.OrganizationModel) *responsedto.HubOrganizationRecord {
	response := &responsedto.HubOrganizationRecord{
		ID:        org.ID,
		Name:      org.Name,
		CreatedAt: org.CreatedAt,
		UpdatedAt: org.UpdatedAt,
	}

	if org.Owner != nil {
		response.OwnerName = org.Owner.Name
	}

	return response
}

func NewHubServiceImpl(db *gorm.DB) services.HubService {
	return &hubServiceImpl{db: db}
}
