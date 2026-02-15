package impl

import (
	"DewaSRY/sociomile-app/internal/services"
	"DewaSRY/sociomile-app/pkg/dtos/requestdto"
	"DewaSRY/sociomile-app/pkg/models"
	"errors"

	"gorm.io/gorm"
)

var (
	ErrOrganizationNotFound = errors.New("organization not found")
)

type webHookConversationServiceImpl struct {
	db *gorm.DB
}

// ProcessConversation implements services.WebHookConversationService.
func (t *webHookConversationServiceImpl) ProcessConversation(req requestdto.WebHooksRequest) error {
	return t.db.Transaction(func(tx *gorm.DB) error {
		var organization models.OrganizationModel

		if err := tx.Model(&models.OrganizationModel{}).
			First(&organization, req.OrganizationID).Error; err != nil {
			return ErrOrganizationNotFound
		}

		user, err := t.findOrCreateUser(tx, req.Email)

		if err != nil {
			return errors.New("failed to create or find user")
		}

		conversation, err := t.findOrCreateConversation(tx, user.ID, organization.ID)

		if err != nil {
			return errors.New("failed to create or find conversation")
		}

		newMessages := models.ConversationMessageModel{
			OrganizationID: conversation.OrganizationID,
			CreatedByID:    user.ID,
			Message:        req.Message,
			ConversationID: conversation.ID,
		}

		if err := tx.Create(&newMessages).Error; err != nil {
			return errors.New("failed to create message")
		}
		return nil
	})
}

func (t *webHookConversationServiceImpl) findOrCreateUser(tx *gorm.DB, email string) (*models.UserModel, error) {
	var userModel models.UserModel

	err := tx.Where("email = ?", email).First(&userModel).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			userModel = models.UserModel{
				Email:    email,
				Password: "",
				Name:     "",
			}
			if err := tx.Create(&userModel).Error; err != nil {
				return nil, errors.New("failed to create user")
			}
		} else {
			return nil, err
		}
	}
	return &userModel, nil
}

func (t *webHookConversationServiceImpl) findOrCreateConversation(tx *gorm.DB, userId uint, organizationId uint) (*models.ConversationModel, error) {
	var conversation models.ConversationModel

	err := tx.
		Where("guest_id = ?", userId).
		Where("organization_id = ?", organizationId).
		Where("status != ?", models.ConversationStatusDone).
		First(&conversation).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			conversation = models.ConversationModel{
				OrganizationID: organizationId,
				GuestID:        userId,
				Status:         models.ConversationStatusPending,
			}

			if err := tx.Create(&conversation).Error; err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	return &conversation, nil
}

func NewWebHookConversationService(db *gorm.DB) services.WebHookConversationService {
	return &webHookConversationServiceImpl{db: db}
}
