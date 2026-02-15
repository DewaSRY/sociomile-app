package tests

import (
	"DewaSRY/sociomile-app/internal/services/impl"
	"DewaSRY/sociomile-app/pkg/dtos/requestdto"
	"DewaSRY/sociomile-app/pkg/models"
	"testing"
)

func TestOrganizationCrudService_CreateOrganization(t *testing.T) {
	tx := SetupTestDB(t)
	service := impl.NewOrganizationCrudService(tx)

	ownerRole, _ := GetOrCreateRole(tx, models.RoleOrganizationOwner)
	owner := models.UserModel{
		Email:    "owner@test.com",
		Name:     "Test Owner",
		Password: "password123",
		RoleID:   ownerRole.ID,
	}
	tx.Create(&owner)

	req := requestdto.CreateOrganizationRequest{
		Name:    "Test Organization",
		OwnerID: owner.ID,
	}

	result, err := service.CreateOrganization(req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result == nil {
		t.Fatal("expected result, got nil")
	}
}

func TestOrganizationCrudService_GetOrganizationByID(t *testing.T) {
	tx := SetupTestDB(t)
	service := impl.NewOrganizationCrudService(tx)

	org, _ := CreateTestOrganizationWithOwner(tx, t, "Test Org")

	result, err := service.GetOrganizationByID(org.ID)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result == nil {
		t.Fatal("expected result, got nil")
	}
}

func TestOrganizationCrudService_GetAllOrganizations(t *testing.T) {
	tx := SetupTestDB(t)
	service := impl.NewOrganizationCrudService(tx)

	CreateTestOrganizationWithOwner(tx, t, "Org1")
	CreateTestOrganizationWithOwner(tx, t, "Org2")

	result, err := service.GetAllOrganizations()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result == nil {
		t.Fatal("expected result, got nil")
	}
}

func TestOrganizationCrudService_UpdateOrganization(t *testing.T) {
	tx := SetupTestDB(t)
	service := impl.NewOrganizationCrudService(tx)

	org, _ := CreateTestOrganizationWithOwner(tx, t, "Original Name")

	req := requestdto.UpdateOrganizationRequest{
		Name: "Updated Name",
	}

	result, err := service.UpdateOrganization(org.ID, req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result == nil {
		t.Fatal("expected result, got nil")
	}
}

func TestOrganizationCrudService_DeleteOrganization(t *testing.T) {
	tx := SetupTestDB(t)
	service := impl.NewOrganizationCrudService(tx)

	org, _ := CreateTestOrganizationWithOwner(tx, t, "Test Org")

	err := service.DeleteOrganization(org.ID)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestOrganizationCrudService_GetOrganizationStats(t *testing.T) {
	tx := SetupTestDB(t)
	service := impl.NewOrganizationCrudService(tx)

	org, _ := CreateTestOrganizationWithOwner(tx, t, "Test Org")

	guestRole, _ := GetOrCreateRole(tx, models.RoleGuest)
	guest := models.UserModel{
		Email:    "guest@test.com",
		Name:     "Guest User",
		Password: "password123",
		RoleID:   guestRole.ID,
	}
	tx.Create(&guest)

	conv1 := models.ConversationModel{
		OrganizationID: org.ID,
		GuestID:        guest.ID,
		Status:         models.ConversationStatusPending,
	}
	tx.Create(&conv1)

	result, err := service.GetOrganizationStats(org.ID)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result == nil {
		t.Fatal("expected result, got nil")
	}
}

func TestOrganizationCrudService_CreateOwnerUser(t *testing.T) {
	tx := SetupTestDB(t)
	service := impl.NewOrganizationCrudService(tx)

	_, err := GetOrCreateRole(tx, models.RoleOrganizationOwner)
	if err != nil {
		t.Fatalf("failed to create role: %v", err)
	}

	user, err := service.CreateOwnerUser("newowner@test.com", "New Owner", "password123")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if user == nil {
		t.Fatal("expected user, got nil")
	}
}
