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

type guestMessageServiceImpl struct {
	db *gorm.DB
}

// GetConversationMessageList implements services.GuestMessageService.
func (t *guestMessageServiceImpl) GetConversationMessageList(user *jwt.Claims, filter filtersdto.FiltersDto, conversationId uint) (*responsedto.ConversationMessagePaginateResponse, error) {
	var messages []models.ConversationMessageModel
	var total int64
	offset := (*filter.Page - 1) * *filter.Limit
	if err := t.db.Model(&models.ConversationMessageModel{}).
		Where("conversation_id = ?", conversationId).
		Count(&total).Error; err != nil {
		return nil, errors.New("failed to count organizations")
	}
	if err := t.db.
		Where("conversation_id = ?", conversationId).
		Offset(offset).
		Limit(*filter.Limit).
		Preload("CreatedBy").
		Order("created_at ASC").
		Find(&messages).Error; err != nil {
		return nil, errors.New("failed to fetch messages")
	}

	var messageResponses []responsedto.ConversationMessageResponse
	for _, msg := range messages {
		messageResponses = append(messageResponses, *t.mapToMessageResponse(&msg))
	}

	return &responsedto.ConversationMessagePaginateResponse{
		Data: messageResponses,
		Metadata: responsedto.PaginateMetaData{
			Total: int(total),
			Page:  1,
			Limit: len(messages),
		},
	}, nil
}

// SendConversationMessage implements services.GuestMessageService.
func (t *guestMessageServiceImpl) SendConversationMessage(user *jwt.Claims, req requestdto.CreateConversationMessageRequest) error {
	var conversation models.ConversationModel
	if err := t.db.First(&conversation, req.ConversationID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("conversation not found")
		}
		return errors.New("failed to fetch organization")
	}

	newMessages := models.ConversationMessageModel{
		OrganizationID: conversation.OrganizationID,
		CreatedByID:    user.UserID,
		Message:        req.Message,
		ConversationID: conversation.ID,
	}
	if err := t.db.Create(&newMessages).Error; err != nil {
		return errors.New("failed to create message")
	}

	return nil
}

func (t *guestMessageServiceImpl) mapToMessageResponse(msg *models.ConversationMessageModel) *responsedto.ConversationMessageResponse {
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

func NewGuestMessageService(db *gorm.DB) services.GuestMessageService {
	return &guestMessageServiceImpl{db: db}
}
