package impl

import (
	"DewaSRY/sociomile-app/internal/database"
	"DewaSRY/sociomile-app/internal/services"
	"DewaSRY/sociomile-app/pkg/dtos/requestdto"
	"DewaSRY/sociomile-app/pkg/dtos/responsedto"
	"DewaSRY/sociomile-app/pkg/models"
	"errors"

	"gorm.io/gorm"
)

type conversationServiceImpl struct {
}

// AssignConversation implements services.ConversationService.
func (t *conversationServiceImpl) AssignConversation(conversationID uint, req requestdto.AssignConversationRequest) (*responsedto.ConversationResponse, error) {
	var conversation models.ConversationModel
	if err := database.DB.First(&conversation, conversationID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("conversation not found")
		}
		return nil, errors.New("failed to fetch conversation")
	}

	// Verify staff exists and belongs to the organization
	var staff models.UserModel
	if err := database.DB.First(&staff, req.OrganizationStaffID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("staff user not found")
		}
		return nil, errors.New("failed to fetch staff user")
	}

	if staff.OrganizationID == nil || *staff.OrganizationID != conversation.OrganizationID {
		return nil, errors.New("staff does not belong to this organization")
	}

	// Update conversation
	conversation.OrganizationStaffID = &req.OrganizationStaffID
	conversation.Status = models.ConversationStatusInProgress

	if err := database.DB.Save(&conversation).Error; err != nil {
		return nil, errors.New("failed to assign conversation")
	}

	// Reload with associations
	if err := database.DB.Preload("Organization").Preload("Guest").Preload("OrganizationStaff").First(&conversation, conversation.ID).Error; err != nil {
		return nil, errors.New("failed to load conversation details")
	}

	return t.mapToConversationResponse(&conversation), nil
}

// CreateConversation implements services.ConversationService.
func (t *conversationServiceImpl) CreateConversation(guestID uint, req requestdto.CreateConversationRequest) (*responsedto.ConversationResponse, error) {
	var organization models.OrganizationModel
	if err := database.DB.First(&organization, req.OrganizationID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("organization not found")
		}
		return nil, errors.New("failed to fetch organization")
	}

	conversation := models.ConversationModel{
		OrganizationID: req.OrganizationID,
		GuestID:        guestID,
		Status:         models.ConversationStatusPending,
	}

	if err := database.DB.Create(&conversation).Error; err != nil {
		return nil, errors.New("failed to create conversation")
	}

	if err := database.DB.Preload("Organization").Preload("Guest").First(&conversation, conversation.ID).Error; err != nil {
		return nil, errors.New("failed to load conversation details")
	}

	return t.mapToConversationResponse(&conversation), nil
}

// GetConversationByID implements services.ConversationService.
func (t *conversationServiceImpl) GetConversationByID(id uint) (*responsedto.ConversationResponse, error) {
	var conversation models.ConversationModel
	if err := database.DB.Preload("Organization").Preload("Guest").Preload("OrganizationStaff").First(&conversation, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("conversation not found")
		}
		return nil, errors.New("failed to fetch conversation")
	}

	return t.mapToConversationResponse(&conversation), nil
}

// GetConversationsByGuest implements services.ConversationService.
func (t *conversationServiceImpl) GetConversationsByGuest(guestID uint) (*responsedto.ConversationListResponse, error) {
	var conversations []models.ConversationModel
	
	if err := database.DB.Where("guest_id = ?", guestID).
		Preload("Organization").
		Preload("OrganizationStaff").
		Order("created_at DESC").
		Find(&conversations).Error; err != nil {
		return nil, errors.New("failed to fetch conversations")
	}

	return t.buildConversationListResponse(conversations), nil
}

// GetConversationsByOrganization implements services.ConversationService.
func (t *conversationServiceImpl) GetConversationsByOrganization(organizationID uint) (*responsedto.ConversationListResponse, error) {
	var conversations []models.ConversationModel
	if err := database.DB.Where("organization_id = ?", organizationID).
		Preload("Guest").Preload("OrganizationStaff").
		Order("created_at DESC").
		Find(&conversations).Error; err != nil {
		return nil, errors.New("failed to fetch conversations")
	}

	return t.buildConversationListResponse(conversations), nil
}

// GetConversationsByStaff implements services.ConversationService.
func (t *conversationServiceImpl) GetConversationsByStaff(staffID uint) (*responsedto.ConversationListResponse, error) {
	var conversations []models.ConversationModel
	if err := database.DB.Where("organization_staff_id = ?", staffID).
		Preload("Organization").Preload("Guest").
		Order("created_at DESC").
		Find(&conversations).Error; err != nil {
		return nil, errors.New("failed to fetch conversations")
	}

	return t.buildConversationListResponse(conversations), nil
}

// UpdateConversationStatus implements services.ConversationService.
func (t *conversationServiceImpl) UpdateConversationStatus(conversationID uint, req requestdto.UpdateConversationRequest) (*responsedto.ConversationResponse, error) {
	var conversation models.ConversationModel
	if err := database.DB.First(&conversation, conversationID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("conversation not found")
		}
		return nil, errors.New("failed to fetch conversation")
	}

	conversation.Status = req.Status

	if err := database.DB.Save(&conversation).Error; err != nil {
		return nil, errors.New("failed to update conversation")
	}

	// Reload with associations
	if err := database.DB.Preload("Organization").Preload("Guest").Preload("OrganizationStaff").First(&conversation, conversation.ID).Error; err != nil {
		return nil, errors.New("failed to load conversation details")
	}

	return t.mapToConversationResponse(&conversation), nil
}

func (t *conversationServiceImpl) mapToConversationResponse(conv *models.ConversationModel) *responsedto.ConversationResponse {
	response := &responsedto.ConversationResponse{
		ID:             conv.ID,
		OrganizationID: conv.OrganizationID,
		GuestID:        conv.GuestID,
		Status:         conv.Status,
		CreatedAt:      conv.CreatedAt,
		UpdatedAt:      conv.UpdatedAt,
	}

	if conv.OrganizationStaffID != nil {
		response.OrganizationStaffID = conv.OrganizationStaffID
	}

	if conv.Organization != nil {
		response.Organization = &responsedto.OrganizationResponse{
			ID:   conv.Organization.ID,
			Name: conv.Organization.Name,
		}
	}

	if conv.Guest != nil {
		response.Guest = &responsedto.UserData{
			ID:    conv.Guest.ID,
			Email: conv.Guest.Email,
			Name:  conv.Guest.Name,
		}
	}

	if conv.OrganizationStaff != nil {
		response.OrganizationStaff = &responsedto.UserData{
			ID:    conv.OrganizationStaff.ID,
			Email: conv.OrganizationStaff.Email,
			Name:  conv.OrganizationStaff.Name,
		}
	}

	return response
}

func (t *conversationServiceImpl) buildConversationListResponse(conversations []models.ConversationModel) *responsedto.ConversationListResponse {
	var conversationResponses []responsedto.ConversationResponse
	for _, conv := range conversations {
		conversationResponses = append(conversationResponses, *t.mapToConversationResponse(&conv))
	}

	return &responsedto.ConversationListResponse{
		Conversations: conversationResponses,
		Metadata: responsedto.PaginateMetaData{
			Total: len(conversationResponses),
			Page:  1,
			Limit: len(conversationResponses),
		},
	}
}

func NewConversationService() services.ConversationService {
	return &conversationServiceImpl{}
}
