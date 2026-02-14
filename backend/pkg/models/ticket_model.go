package models

import (
	"time"

	"gorm.io/gorm"
)

type TicketModel struct {
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
	TicketNumber   string             `gorm:"uniqueIndex;not null" json:"ticket_number"`
	Name           string             `gorm:"not null" json:"name"`
	Status         string             `gorm:"not null;default:'pending'" json:"status"`
}

func (TicketModel) TableName() string {
	return "tickets"
}

// Constants for ticket status
const (
	TicketStatusPending    = "pending"
	TicketStatusInProgress = "in_progress"
	TicketStatusDone       = "done"
)
