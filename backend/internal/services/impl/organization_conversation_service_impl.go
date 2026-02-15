package impl

import (
	"DewaSRY/sociomile-app/internal/services"
	"DewaSRY/sociomile-app/pkg/dtos/filtersdto"
	"DewaSRY/sociomile-app/pkg/dtos/requestdto"
	"DewaSRY/sociomile-app/pkg/dtos/responsedto"
	"DewaSRY/sociomile-app/pkg/lib/jwt"
	"DewaSRY/sociomile-app/pkg/models"
	"errors"

	"gorm.io/gorm"
)

type organizationConversationServiceImpl struct {
	db *gorm.DB
}

// AssignConversation implements services.ConversationService.
func (t *organizationConversationServiceImpl) AssignConversation(conversationID uint, req requestdto.AssignConversationRequest) (*responsedto.ConversationResponse, error) {
	var conversation models.ConversationModel
	if err := t.db.First(&conversation, conversationID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("conversation not found")
		}
		return nil, errors.New("failed to fetch conversation")
	}

	var staff models.UserModel
	if err := t.db.First(&staff, req.OrganizationStaffID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("staff user not found")
		}
		return nil, errors.New("failed to fetch staff user")
	}

	if staff.OrganizationID == nil || *staff.OrganizationID != conversation.OrganizationID {
		return nil, errors.New("staff does not belong to this organization")
	}

	conversation.OrganizationStaffID = &req.OrganizationStaffID
	conversation.Status = models.ConversationStatusInProgress

	if err := t.db.Save(&conversation).Error; err != nil {
		return nil, errors.New("failed to assign conversation")
	}

	if err := t.db.Preload("Organization").Preload("Guest").Preload("OrganizationStaff").First(&conversation, conversation.ID).Error; err != nil {
		return nil, errors.New("failed to load conversation details")
	}

	return t.mapToConversationResponse(&conversation), nil
}

// GetConversationByID implements services.ConversationService.
func (t *organizationConversationServiceImpl) GetConversationByID(id uint) (*responsedto.ConversationResponse, error) {
	var conversation models.ConversationModel
	if err := t.db.Preload("Organization").
		Preload("Guest").
		Preload("OrganizationStaff").
		Preload("ConversationMessages").
		First(&conversation, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("conversation not found")
		}
		return nil, errors.New("failed to fetch conversation")
	}

	return t.mapToConversationResponse(&conversation), nil
}

// GetConversationsByOrganization implements services.ConversationService.
func (t *organizationConversationServiceImpl) GetConversationsList(user *jwt.Claims, filter filtersdto.FiltersDto) (*responsedto.ConversationListResponse, error) {
	var conversations []models.ConversationModel
	var total int64
	offset := (*filter.Page - 1) * *filter.Limit

	if err := t.db.Model(&models.OrganizationModel{}).
		Count(&total).Error; err != nil {
		return nil, errors.New("failed to count organizations")
	}

	if err := t.db.Where("organization_id = ?", user.OrganizationId).
		Offset(offset).Limit(*filter.Limit).
		Preload("Guest").
		Preload("OrganizationStaff").
		Order("created_at DESC").
		Find(&conversations).Error; err != nil {
		return nil, errors.New("failed to fetch conversations")
	}

	var conversationResponses []responsedto.ConversationResponse
	for _, conv := range conversations {
		conversationResponses = append(conversationResponses, *t.mapToConversationResponse(&conv))
	}

	return &responsedto.ConversationListResponse{
		Conversations: conversationResponses,
		Metadata: responsedto.PaginateMetaData{
			Total: int(total),
			Page:  1,
			Limit: len(conversationResponses),
		},
	}, nil
}

// UpdateConversationStatus implements services.ConversationService.
func (t *organizationConversationServiceImpl) UpdateConversationStatus(conversationID uint, req requestdto.UpdateConversationRequest) error {
	var conversation models.ConversationModel
	if err := t.db.First(&conversation, conversationID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("conversation not found")
		}
		return errors.New("failed to fetch conversation")
	}

	conversation.Status = req.Status

	if err := t.db.Save(&conversation).Error; err != nil {
		return errors.New("failed to update conversation")
	}

	if err := t.db.Preload("Organization").
		Preload("Guest").
		Preload("OrganizationStaff").First(&conversation, conversation.ID).Error; err != nil {
		return errors.New("failed to load conversation details")
	}

	return nil
}

func (t *organizationConversationServiceImpl) mapToConversationResponse(conv *models.ConversationModel) *responsedto.ConversationResponse {
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
	   if len(conv.ConversationMessages) > 0 {
        messages := make([]responsedto.ConversationMessageResponse, 0, len(conv.ConversationMessages))
        for i := range conv.ConversationMessages {
            msg := t.mapToMessageResponse(&conv.ConversationMessages[i])
            messages = append(messages, *msg)
        }
        response.Messages = messages
    } else {
        response.Messages = []responsedto.ConversationMessageResponse{}
    }

	return response
}

func (t *organizationConversationServiceImpl) mapToMessageResponse(msg *models.ConversationMessageModel) *responsedto.ConversationMessageResponse {
	response := &responsedto.ConversationMessageResponse{
		ID:             msg.ID,
		OrganizationID: msg.OrganizationID,
		ConversationID: msg.ConversationID,
		CreatedByID:    msg.CreatedByID,
		Message:        msg.Message,
		CreatedAt:      msg.CreatedAt,
		UpdatedAt:      msg.UpdatedAt,
	}

	if msg.CreatedBy != nil {
		response.CreatedBy = &responsedto.UserData{
			ID:    msg.CreatedBy.ID,
			Email: msg.CreatedBy.Email,
			Name:  msg.CreatedBy.Name,
		}
	}

	return response
}

func NewConversationService(db *gorm.DB) services.OrganizationConversationService {
	return &organizationConversationServiceImpl{db: db}
}
