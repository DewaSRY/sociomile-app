package tests

import (
	"DewaSRY/sociomile-app/internal/services/impl"
	"DewaSRY/sociomile-app/pkg/dtos/filtersdto"
	"DewaSRY/sociomile-app/pkg/dtos/requestdto"
	jwtLib "DewaSRY/sociomile-app/pkg/lib/jwt"
	"DewaSRY/sociomile-app/pkg/models"
	"testing"
)

func TestOrganizationService_CreateStaff(t *testing.T) {
	tx := SetupTestDB(t)
	orgService := impl.NewOrganizationService(tx)

	org, owner := CreateTestOrganizationWithOwner(tx, t, "Test Org")

	salesRole, err := GetOrCreateRole(tx, models.RoleOrganizationSales)
	if err != nil {
		t.Fatalf("failed to get sales role: %v", err)
	}
	_ = salesRole

	claims := &jwtLib.Claims{
		UserID:         owner.ID,
		Email:          owner.Email,
		RoleID:         owner.RoleID,
		OrganizationId: &org.ID,
	}

	req := requestdto.RegisterRequest{
		Email:    "staff@example.com",
		Name:     "Staff User",
		Password: "password123",
	}

	err = orgService.CreateStaff(req, claims)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestOrganizationService_GetStaffList(t *testing.T) {
	tx := SetupTestDB(t)
	orgService := impl.NewOrganizationService(tx)

	org, owner := CreateTestOrganizationWithOwner(tx, t, "Test Org")

	salesRole, err := GetOrCreateRole(tx, models.RoleOrganizationSales)
	if err != nil {
		t.Fatalf("failed to get sales role: %v", err)
	}

	staff1 := models.UserModel{
		Email:          "staff1@example.com",
		Name:           "Staff One",
		Password:       "password123",
		RoleID:         salesRole.ID,
		OrganizationID: &org.ID,
	}
	if err := tx.Create(&staff1).Error; err != nil {
		t.Fatalf("failed to create staff1: %v", err)
	}

	staff2 := models.UserModel{
		Email:          "staff2@example.com",
		Name:           "Staff Two",
		Password:       "password123",
		RoleID:         salesRole.ID,
		OrganizationID: &org.ID,
	}
	if err := tx.Create(&staff2).Error; err != nil {
		t.Fatalf("failed to create staff2: %v", err)
	}

	claims := &jwtLib.Claims{
		UserID:         owner.ID,
		Email:          owner.Email,
		RoleID:         owner.RoleID,
		OrganizationId: &org.ID,
	}

	page := 1
	limit := 10
	filter := filtersdto.FiltersDto{
		Page:  &page,
		Limit: &limit,
	}

	result, err := orgService.GetStaffList(filter, claims)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result == nil {
		t.Fatal("expected result, got nil")
	}
}
