package models

import (
	"time"

	"gorm.io/gorm"
)

type ConversationMessageModel struct {
	ID             uint               `gorm:"primarykey" json:"id"`
	CreatedAt      time.Time          `json:"created_at"`
	UpdatedAt      time.Time          `json:"updated_at"`
	DeletedAt      gorm.DeletedAt     `gorm:"index" json:"-"`
	OrganizationID uint               `gorm:"not null;index" json:"organization_id"`
	Organization   *OrganizationModel `gorm:"foreignKey:OrganizationID" json:"organization,omitempty"`
	ConversationID uint               `gorm:"not null;index" json:"conversation_id"`
	Conversation   *ConversationModel `gorm:"foreignKey:ConversationID" json:"conversation,omitempty"`
	CreatedByID    uint               `gorm:"not null;index" json:"created_by_id"`
	CreatedBy      *UserModel         `gorm:"foreignKey:CreatedByID" json:"created_by,omitempty"`
	Message        string             `gorm:"type:text;not null" json:"message"`
}

func (ConversationMessageModel) TableName() string {
	return "conversation_messages"
}
