package tests

import (
	"DewaSRY/sociomile-app/internal/services/impl"
	"DewaSRY/sociomile-app/pkg/models"
	"testing"
)

func TestIsUserAuthorize(t *testing.T) {
	tx := SetupTestDB(t)
	authorizeService := impl.NewAuthorizeService(tx)

	var role models.UserRoleModel
	if err := tx.Where("name = ?", models.RoleSuperAdmin).FirstOrCreate(&role, models.UserRoleModel{
		Name: models.RoleSuperAdmin,
	}).Error; err != nil {
		t.Fatalf("failed to get or create test role: %v", err)
	}

	err := authorizeService.IsUserAuthorize(role.ID, []string{models.RoleSuperAdmin, models.RoleOrganizationOwner})
	if err != nil {
		t.Fatalf("expected no error for authorized user, got %v", err)
	}
}
