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

type guestConversationServiceImpl struct {
	db *gorm.DB
}

// CreateConversation implements services.GuestConversationService.
func (t *guestConversationServiceImpl) CreateConversation(user *jwt.Claims, req requestdto.CreateConversationRequest) error {
	var organization models.OrganizationModel
	if err := t.db.First(&organization, req.OrganizationID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("organization not found")
		}
		return errors.New("failed to fetch organization")
	}

	conversation := models.ConversationModel{
		OrganizationID: req.OrganizationID,
		GuestID:        user.UserID,
		Status:         models.ConversationStatusPending,
	}

	if err := t.db.Create(&conversation).Error; err != nil {
		return errors.New("failed to create conversation")
	}

	return nil
}

// GetConversation implements services.GuestConversationService.
func (t *guestConversationServiceImpl) GetConversation(user *jwt.Claims, filter filtersdto.FiltersDto) (*responsedto.ConversationListPaginateResponse, error) {
	var conversations []models.ConversationModel
	var total int64
	offset := (*filter.Page - 1) * *filter.Limit

	if err := t.db.Model(&models.ConversationModel{}).
		Where("guest_id = ?", user.UserID).
		Count(&total).Error; err != nil {
		return nil, errors.New("failed to count organizations")
	}

	if err := t.db.Where("guest_id = ?", user.UserID).
		Offset(offset).
		Preload("Organization").Preload("Guest").
		Order("created_at DESC").
		Find(&conversations).Error; err != nil {
		return nil, errors.New("failed to fetch conversations")
	}

	guestConversationList := make([]responsedto.ConversationResponse, 0, len(conversations))

	for _, org := range conversations {
		guestConversationList = append(
			guestConversationList,
			*t.mapToConversationResponse(&org),
		)
	}

	return &responsedto.ConversationListPaginateResponse{
		Data: guestConversationList,
		Metadata: responsedto.PaginateMetaData{
			Total: int(total),
			Page:  *filter.Page,
			Limit: *filter.Limit,
		},
	}, nil
}

func (t *guestConversationServiceImpl) mapToConversationResponse(conv *models.ConversationModel) *responsedto.ConversationResponse {
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

func NewGuestConversationService(db *gorm.DB) services.GuestConversationService {
	return &guestConversationServiceImpl{db: db}
}
