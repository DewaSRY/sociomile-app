package impl

import (
	"DewaSRY/sociomile-app/internal/database"
	"DewaSRY/sociomile-app/internal/services"
	"DewaSRY/sociomile-app/pkg/dtos/requestdto"
	"DewaSRY/sociomile-app/pkg/dtos/responsedto"
	"DewaSRY/sociomile-app/pkg/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type organizationServiceImpl struct {
}

// CreateOrganization implements services.OrganizationService.
func (t *organizationServiceImpl) CreateOrganization(req requestdto.CreateOrganizationRequest) (*responsedto.OrganizationResponse, error) {
	var owner models.UserModel
	if err := database.DB.First(&owner, req.OwnerID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("owner user not found")
		}
		return nil, errors.New("failed to fetch owner user")
	}

	organization := models.OrganizationModel{
		Name:    req.Name,
		OwnerID: req.OwnerID,
	}

	if err := database.DB.Create(&organization).Error; err != nil {
		return nil, errors.New("failed to create organization")
	}

	if err := database.DB.Model(&owner).Updates(map[string]interface{}{
		"organization_id": organization.ID,
	}).Error; err != nil {
		return nil, errors.New("failed to update owner's organization")
	}

	if err := database.DB.Preload("Owner").First(&organization, organization.ID).Error; err != nil {
		return nil, errors.New("failed to load organization details")
	}

	return t.mapToOrganizationResponse(&organization), nil
}

// CreateOwnerUser implements services.OrganizationService.
func (t *organizationServiceImpl) CreateOwnerUser(email string, name string, password string) (*models.UserModel, error) {
	var existingUser models.UserModel
	if err := database.DB.Where("email = ?", email).First(&existingUser).Error; err == nil {
		return nil, errors.New("user with this email already exists")
	}

	var ownerRole models.UserRoleModel
	if err := database.DB.Where("name = ?", models.RoleOrganizationOwner).First(&ownerRole).Error; err != nil {
		return nil, fmt.Errorf("organization owner role not found: %v", err)
	}

	user := models.UserModel{
		Email:    email,
		Name:     name,
		Password: password,
		RoleID:   ownerRole.ID,
	}

	if err := database.DB.Create(&user).Error; err != nil {
		return nil, errors.New("failed to create user")
	}

	return &user, nil
}

// DeleteOrganization implements services.OrganizationService.
func (t *organizationServiceImpl) DeleteOrganization(id uint) error {
	var organization models.OrganizationModel
	if err := database.DB.First(&organization, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("organization not found")
		}
		return errors.New("failed to fetch organization")
	}

	if err := database.DB.Delete(&organization).Error; err != nil {
		return errors.New("failed to delete organization")
	}

	return nil
}

// GetAllOrganizations implements services.OrganizationService.
func (t *organizationServiceImpl) GetAllOrganizations() (*responsedto.OrganizationListResponse, error) {
	var organizations []models.OrganizationModel
	if err := database.DB.Preload("Owner").Find(&organizations).Error; err != nil {
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

// GetOrganizationByID implements services.OrganizationService.
func (t *organizationServiceImpl) GetOrganizationByID(id uint) (*responsedto.OrganizationResponse, error) {
	var organization models.OrganizationModel
	if err := database.DB.Preload("Owner").First(&organization, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("organization not found")
		}
		return nil, errors.New("failed to fetch organization")
	}

	return t.mapToOrganizationResponse(&organization), nil
}

// GetOrganizationStats implements services.OrganizationService.
func (t *organizationServiceImpl) GetOrganizationStats(organizationID uint) (map[string]interface{}, error) {
	var totalConversations int64
	var totalTickets int64
	var pendingConversations int64
	var pendingTickets int64

	if err := database.DB.Model(&models.ConversationModel{}).
		Where("organization_id = ?", organizationID).
		Count(&totalConversations).Error; err != nil {
		return nil, errors.New("failed to count conversations")
	}

	if err := database.DB.Model(&models.ConversationModel{}).
		Where("organization_id = ? AND status = ?", organizationID, models.ConversationStatusPending).
		Count(&pendingConversations).Error; err != nil {
		return nil, errors.New("failed to count pending conversations")
	}

	if err := database.DB.Model(&models.TicketModel{}).
		Where("organization_id = ?", organizationID).
		Count(&totalTickets).Error; err != nil {
		return nil, errors.New("failed to count tickets")
	}

	if err := database.DB.Model(&models.TicketModel{}).
		Where("organization_id = ? AND status = ?", organizationID, models.TicketStatusPending).
		Count(&pendingTickets).Error; err != nil {
		return nil, errors.New("failed to count pending tickets")
	}

	return map[string]interface{}{
		"total_conversations":   totalConversations,
		"pending_conversations": pendingConversations,
		"total_tickets":         totalTickets,
		"pending_tickets":       pendingTickets,
	}, nil
}

// UpdateOrganization implements services.OrganizationService.
func (t *organizationServiceImpl) UpdateOrganization(id uint, req requestdto.UpdateOrganizationRequest) (*responsedto.OrganizationResponse, error) {
	var organization models.OrganizationModel
	if err := database.DB.First(&organization, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("organization not found")
		}
		return nil, errors.New("failed to fetch organization")
	}

	organization.Name = req.Name

	if err := database.DB.Save(&organization).Error; err != nil {
		return nil, errors.New("failed to update organization")
	}

	if err := database.DB.Preload("Owner").First(&organization, organization.ID).Error; err != nil {
		return nil, errors.New("failed to load organization details")
	}

	return t.mapToOrganizationResponse(&organization), nil
}

func (t *organizationServiceImpl) mapToOrganizationResponse(org *models.OrganizationModel) *responsedto.OrganizationResponse {
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

func NewOrganizationService() services.OrganizationService {
	return &organizationServiceImpl{}
}
