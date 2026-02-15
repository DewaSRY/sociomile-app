package tests

import (
	"DewaSRY/sociomile-app/internal/services/impl"
	"DewaSRY/sociomile-app/pkg/dtos/requestdto"
	"DewaSRY/sociomile-app/pkg/models"
	"testing"
)

func TestOrganizationCrudService_CreateOrganization_Success(t *testing.T) {
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

	if result.Name != req.Name {
		t.Fatalf("expected name %s, got %s", req.Name, result.Name)
	}

	if result.OwnerID != req.OwnerID {
		t.Fatalf("expected owner ID %d, got %d", req.OwnerID, result.OwnerID)
	}
}

func TestOrganizationCrudService_CreateOrganization_OwnerNotFound(t *testing.T) {
	tx := SetupTestDB(t)
	service := impl.NewOrganizationCrudService(tx)

	req := requestdto.CreateOrganizationRequest{
		Name:    "Test Organization",
		OwnerID: 99999,
	}

	_, err := service.CreateOrganization(req)
	if err == nil {
		t.Fatal("expected error for non-existent owner, got nil")
	}
}

func TestOrganizationCrudService_GetOrganizationByID_Success(t *testing.T) {
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

	if result.ID != org.ID {
		t.Fatalf("expected ID %d, got %d", org.ID, result.ID)
	}

	if result.Name != org.Name {
		t.Fatalf("expected name %s, got %s", org.Name, result.Name)
	}
}

func TestOrganizationCrudService_GetOrganizationByID_NotFound(t *testing.T) {
	tx := SetupTestDB(t)
	service := impl.NewOrganizationCrudService(tx)

	_, err := service.GetOrganizationByID(99999)
	if err == nil {
		t.Fatal("expected error for non-existent organization, got nil")
	}
}

func TestOrganizationCrudService_GetAllOrganizations_Success(t *testing.T) {
	tx := SetupTestDB(t)
	service := impl.NewOrganizationCrudService(tx)

	org1, _ := CreateTestOrganizationWithOwner(tx, t, "Org1")
	org2, _ := CreateTestOrganizationWithOwner(tx, t, "Org2")

	result, err := service.GetAllOrganizations()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result == nil {
		t.Fatal("expected result, got nil")
	}

	if len(result.Organizations) < 2 {
		t.Fatalf("expected at least 2 organizations, got %d", len(result.Organizations))
	}

	foundOrg1 := false
	foundOrg2 := false
	for _, org := range result.Organizations {
		if org.ID == org1.ID {
			foundOrg1 = true
		}
		if org.ID == org2.ID {
			foundOrg2 = true
		}
	}

	if !foundOrg1 || !foundOrg2 {
		t.Fatal("created organizations not found in result")
	}
}

func TestOrganizationCrudService_UpdateOrganization_Success(t *testing.T) {
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

	if result.Name != req.Name {
		t.Fatalf("expected name %s, got %s", req.Name, result.Name)
	}

	if result.ID != org.ID {
		t.Fatalf("expected ID %d, got %d", org.ID, result.ID)
	}
}

func TestOrganizationCrudService_UpdateOrganization_NotFound(t *testing.T) {
	tx := SetupTestDB(t)
	service := impl.NewOrganizationCrudService(tx)

	req := requestdto.UpdateOrganizationRequest{
		Name: "Updated Name",
	}

	_, err := service.UpdateOrganization(99999, req)
	if err == nil {
		t.Fatal("expected error for non-existent organization, got nil")
	}
}

func TestOrganizationCrudService_DeleteOrganization_Success(t *testing.T) {
	tx := SetupTestDB(t)
	service := impl.NewOrganizationCrudService(tx)

	org, _ := CreateTestOrganizationWithOwner(tx, t, "Test Org")

	err := service.DeleteOrganization(org.ID)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	var deletedOrg models.OrganizationModel
	err = tx.First(&deletedOrg, org.ID).Error
	if err == nil {
		t.Fatal("expected organization to be deleted")
	}
}

func TestOrganizationCrudService_DeleteOrganization_NotFound(t *testing.T) {
	tx := SetupTestDB(t)
	service := impl.NewOrganizationCrudService(tx)

	err := service.DeleteOrganization(99999)
	if err == nil {
		t.Fatal("expected error for non-existent organization, got nil")
	}
}

func TestOrganizationCrudService_GetOrganizationStats_Success(t *testing.T) {
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

	conv2 := models.ConversationModel{
		OrganizationID: org.ID,
		GuestID:        guest.ID,
		Status:         models.ConversationStatusInProgress,
	}
	tx.Create(&conv2)

	ticket1 := models.TicketModel{
		OrganizationID: org.ID,
		ConversationID: conv1.ID,
		Name:           "Test Ticket",
		TicketNumber:   "TICK-001",
		CreatedByID:    guest.ID,
		Status:         models.TicketStatusPending,
	}
	tx.Create(&ticket1)

	result, err := service.GetOrganizationStats(org.ID)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result == nil {
		t.Fatal("expected result, got nil")
	}

	if result["total_conversations"].(int64) != 2 {
		t.Fatalf("expected 2 conversations, got %d", result["total_conversations"])
	}

	if result["pending_conversations"].(int64) != 1 {
		t.Fatalf("expected 1 pending conversation, got %d", result["pending_conversations"])
	}

	if result["total_tickets"].(int64) != 1 {
		t.Fatalf("expected 1 ticket, got %d", result["total_tickets"])
	}

	if result["pending_tickets"].(int64) != 1 {
		t.Fatalf("expected 1 pending ticket, got %d", result["pending_tickets"])
	}
}

func TestOrganizationCrudService_CreateOwnerUser_Success(t *testing.T) {
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

	if user.Email != "newowner@test.com" {
		t.Fatalf("expected email newowner@test.com, got %s", user.Email)
	}

	if user.Name != "New Owner" {
		t.Fatalf("expected name New Owner, got %s", user.Name)
	}
}

func TestOrganizationCrudService_CreateOwnerUser_DuplicateEmail(t *testing.T) {
	tx := SetupTestDB(t)
	service := impl.NewOrganizationCrudService(tx)

	_, err := GetOrCreateRole(tx, models.RoleOrganizationOwner)
	if err != nil {
		t.Fatalf("failed to create role: %v", err)
	}

	_, err = service.CreateOwnerUser("owner@test.com", "Owner", "password123")
	if err != nil {
		t.Fatalf("expected no error on first creation, got %v", err)
	}

	_, err = service.CreateOwnerUser("owner@test.com", "Another Owner", "password456")
	if err == nil {
		t.Fatal("expected error for duplicate email, got nil")
	}
}
