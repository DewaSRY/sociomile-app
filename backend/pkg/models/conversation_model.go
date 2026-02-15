package models

import (
	"time"

	"gorm.io/gorm"
)

type ConversationModel struct {
	ID                   uint               `gorm:"primarykey" json:"id"`
	CreatedAt            time.Time          `json:"created_at"`
	UpdatedAt            time.Time          `json:"updated_at"`
	DeletedAt            gorm.DeletedAt     `gorm:"index" json:"-"`
	OrganizationID       uint               `gorm:"not null;index" json:"organization_id"`
	Organization         *OrganizationModel `gorm:"foreignKey:OrganizationID" json:"organization,omitempty"`
	GuestID              uint               `gorm:"not null;index" json:"guest_id"`
	Guest                *UserModel         `gorm:"foreignKey:GuestID" json:"guest,omitempty"`
	OrganizationStaffID  *uint              `gorm:"index" json:"organization_staff_id,omitempty"`
	OrganizationStaff    *UserModel         `gorm:"foreignKey:OrganizationStaffID" json:"organization_staff,omitempty"`
	ConversationMessages  []ConversationMessageModel `gorm:"foreignKey:ConversationID"`
	Status               string             `gorm:"not null;default:'pending'" json:"status"`
}

func (ConversationModel) TableName() string {
	return "conversations"
}

// Constants for conversation status
const (
	ConversationStatusPending    = "pending"
	ConversationStatusInProgress = "in_progress"
	ConversationStatusDone       = "done"
)
