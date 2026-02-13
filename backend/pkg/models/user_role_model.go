package models

import (
	"time"

	"gorm.io/gorm"
)

type UserRoleModel struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Name      string         `gorm:"uniqueIndex;not null" json:"name"`
}

func (UserRoleModel) TableName() string {
	return "user_roles"
}

// Constants for role names
const (
	RoleSuperAdmin       = "super_admin"
	RoleOrganizationOwner = "organization_owner"
	RoleOrganizationSales = "organization_sales"
	RoleGuest            = "guest"
)
