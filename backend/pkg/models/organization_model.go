package models

import (
	"time"

	"gorm.io/gorm"
)

type OrganizationModel struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Name      string         `gorm:"not null" json:"name"`
	OwnerID   uint           `gorm:"not null" json:"owner_id"`
	
	Owner     *UserModel     `gorm:"foreignKey:OwnerID" json:"owner,omitempty"`
}

func (OrganizationModel) TableName() string {
	return "organizations"
}
