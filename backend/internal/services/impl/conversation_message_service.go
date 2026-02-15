package impl

import (
	"DewaSRY/sociomile-app/internal/services"
	"DewaSRY/sociomile-app/pkg/dtos/requestdto"
	"DewaSRY/sociomile-app/pkg/dtos/responsedto"
	"DewaSRY/sociomile-app/pkg/models"
	"errors"

	"gorm.io/gorm"
)

type conversationMessageServiceImpl struct {
	db *gorm.DB
}

// CreateMessage implements services.ConversationMessageService.
func (t *conversationMessageServiceImpl) CreateMessage(userID uint, req requestdto.CreateConversationMessageRequest) (*responsedto.ConversationMessageResponse, error) {
	var conversation models.ConversationModel
	if err := t.db.First(&conversation, req.ConversationID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("conversation not found")
		}
		return nil, errors.New("failed to fetch conversation")
	}

	message := models.ConversationMessageModel{
		OrganizationID: conversation.OrganizationID,
		ConversationID: req.ConversationID,
		CreatedByID:    userID,
		Message:        req.Message,
	}

	if err := t.db.Create(&message).Error; err != nil {
		return nil, errors.New("failed to create message")
	}

	if err := t.db.Preload("CreatedBy").First(&message, message.ID).Error; err != nil {
		return nil, errors.New("failed to load message details")
	}

	return t.mapToMessageResponse(&message), nil
}

// GetMessagesByConversation implements services.ConversationMessageService.
func (t *conversationMessageServiceImpl) GetMessagesByConversation(conversationID uint) (*responsedto.ConversationMessageListResponse, error) {
		var messages []models.ConversationMessageModel
	if err := t.db.Where("conversation_id = ?", conversationID).
		Preload("CreatedBy").
		Order("created_at ASC").
		Find(&messages).Error; err != nil {
		return nil, errors.New("failed to fetch messages")
	}

	var messageResponses []responsedto.ConversationMessageResponse
	for _, msg := range messages {
		messageResponses = append(messageResponses, *t.mapToMessageResponse(&msg))
	}

	return &responsedto.ConversationMessageListResponse{
		Messages: messageResponses,
		Metadata: responsedto.PaginateMetaData{
			Total: len(messages),
			Page: 1,
			Limit: len(messages),
		} ,
	}, nil
}

func (conversationMessageServiceImpl *conversationMessageServiceImpl) mapToMessageResponse(msg *models.ConversationMessageModel) *responsedto.ConversationMessageResponse {
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


func NewConversationMessageService(db *gorm.DB) services.ConversationMessageService {
	return &conversationMessageServiceImpl{db:db}
}
