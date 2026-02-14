package impl

import (
	"DewaSRY/sociomile-app/internal/services"
	"DewaSRY/sociomile-app/pkg/models"
	"errors"
	"slices"

	"gorm.io/gorm"
)

type authorizeServiceImpl struct {
	db *gorm.DB
}

// IsUserAuthorize implements services.AuthorizeService.
func (t *authorizeServiceImpl) IsUserAuthorize(roleId uint, allowedRoles []string) error {
	var userRole models.UserRoleModel

	if err := t.db.First(&userRole, roleId).Error; err != nil {
		return errors.New("not authorized")
	}

	if !slices.Contains(allowedRoles, userRole.Name) {

		return errors.New("not authorized because not in allowed")
	}

	return nil
}

func NewAuthorizeService(db *gorm.DB) services.AuthorizeService {
	return &authorizeServiceImpl{db: db}
}
