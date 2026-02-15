package tests

import (
	"DewaSRY/sociomile-app/internal/services/impl"
	"DewaSRY/sociomile-app/pkg/models"
	"testing"
)

func TestIsUserAuthorize_Success(t *testing.T) {
	tx := SetupTestDB(t)
	authorizeService := impl.NewAuthorizeService(tx)

	// Get or create a test role
	var role models.UserRoleModel
	if err := tx.Where("name = ?", models.RoleSuperAdmin).FirstOrCreate(&role, models.UserRoleModel{
		Name: models.RoleSuperAdmin,
	}).Error; err != nil {
		t.Fatalf("failed to get or create test role: %v", err)
	}

	// Test with allowed role
	err := authorizeService.IsUserAuthorize(role.ID, []string{models.RoleSuperAdmin, models.RoleOrganizationOwner})
	if err != nil {
		t.Fatalf("expected no error for authorized user, got %v", err)
	}
}

func TestIsUserAuthorize_RoleNotFound(t *testing.T) {
	tx := SetupTestDB(t)
	authorizeService := impl.NewAuthorizeService(tx)

	// Test with non-existent role ID
	err := authorizeService.IsUserAuthorize(999, []string{models.RoleSuperAdmin})
	if err == nil {
		t.Fatal("expected error for non-existent role, got nil")
	}

	if err.Error() != "not authorized" {
		t.Fatalf("expected 'not authorized' error, got %v", err)
	}
}

func TestIsUserAuthorize_NotInAllowedRoles(t *testing.T) {
	tx := SetupTestDB(t)
	authorizeService := impl.NewAuthorizeService(tx)

	// Get or create a test role
	var role models.UserRoleModel
	if err := tx.Where("name = ?", models.RoleGuest).FirstOrCreate(&role, models.UserRoleModel{
		Name: models.RoleGuest,
	}).Error; err != nil {
		t.Fatalf("failed to get or create test role: %v", err)
	}

	// Test with role not in allowed list
	err := authorizeService.IsUserAuthorize(role.ID, []string{models.RoleSuperAdmin, models.RoleOrganizationOwner})
	if err == nil {
		t.Fatal("expected error for unauthorized role, got nil")
	}

	if err.Error() != "not authorized because not in allowed" {
		t.Fatalf("expected 'not authorized because not in allowed' error, got %v", err)
	}
}

func TestIsUserAuthorize_MultipleAllowedRoles(t *testing.T) {
	tx := SetupTestDB(t)
	authorizeService := impl.NewAuthorizeService(tx)

	// Get or create test roles
	roleNames := []string{models.RoleOrganizationOwner, models.RoleOrganizationSales}

	for _, roleName := range roleNames {
		var role models.UserRoleModel
		if err := tx.Where("name = ?", roleName).FirstOrCreate(&role, models.UserRoleModel{
			Name: roleName,
		}).Error; err != nil {
			t.Fatalf("failed to get or create test role: %v", err)
		}

		// Test each role is authorized
		err := authorizeService.IsUserAuthorize(role.ID, []string{models.RoleOrganizationOwner, models.RoleOrganizationSales})
		if err != nil {
			t.Fatalf("expected no error for authorized user with role %s, got %v", role.Name, err)
		}
	}
}

func TestIsUserAuthorize_EmptyAllowedRoles(t *testing.T) {
	tx := SetupTestDB(t)
	authorizeService := impl.NewAuthorizeService(tx)

	// Get or create a test role
	var role models.UserRoleModel
	if err := tx.Where("name = ?", models.RoleSuperAdmin).FirstOrCreate(&role, models.UserRoleModel{
		Name: models.RoleSuperAdmin,
	}).Error; err != nil {
		t.Fatalf("failed to get or create test role: %v", err)
	}

	// Test with empty allowed roles list
	err := authorizeService.IsUserAuthorize(role.ID, []string{})
	if err == nil {
		t.Fatal("expected error for empty allowed roles list, got nil")
	}

	if err.Error() != "not authorized because not in allowed" {
		t.Fatalf("expected 'not authorized because not in allowed' error, got %v", err)
	}
}
