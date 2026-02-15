package tests

import (
	"DewaSRY/sociomile-app/internal/services/impl"
	"DewaSRY/sociomile-app/pkg/dtos/filtersdto"
	"DewaSRY/sociomile-app/pkg/dtos/requestdto"
	jwtLib "DewaSRY/sociomile-app/pkg/lib/jwt"
	"DewaSRY/sociomile-app/pkg/models"
	"testing"
)

func TestOrganizationConversationService_GetConversationsList_Success(t *testing.T) {
	tx := SetupTestDB(t)
	service := impl.NewConversationService(tx)

	org, owner := CreateTestOrganizationWithOwner(tx, t, "Test Org")

	guestRole, _ := GetOrCreateRole(tx, models.RoleGuest)
	guest := models.UserModel{
		Email:    "guest@test.com",
		Name:     "Guest",
		Password: "password",
		RoleID:   guestRole.ID,
	}
	tx.Create(&guest)

	conv := models.ConversationModel{
		OrganizationID: org.ID,
		GuestID:        guest.ID,
		Status:         models.ConversationStatusPending,
	}
	tx.Create(&conv)

	claims := &jwtLib.Claims{
		UserID:         owner.ID,
		OrganizationId: &org.ID,
	}

	page := 1
	limit := 10
	filter := filtersdto.FiltersDto{Page: &page, Limit: &limit}

	result, err := service.GetConversationsList(claims, filter)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result == nil {
		t.Fatal("expected result, got nil")
	}
}

func TestOrganizationConversationService_GetConversationByID_Success(t *testing.T) {
	tx := SetupTestDB(t)
	service := impl.NewConversationService(tx)

	org, _ := CreateTestOrganizationWithOwner(tx, t, "Test Org")

	guestRole, _ := GetOrCreateRole(tx, models.RoleGuest)
	guest := models.UserModel{
		Email:    "guest@test.com",
		Name:     "Guest",
		Password: "password",
		RoleID:   guestRole.ID,
	}
	tx.Create(&guest)

	conv := models.ConversationModel{
		OrganizationID: org.ID,
		GuestID:        guest.ID,
		Status:         models.ConversationStatusPending,
	}
	tx.Create(&conv)

	result, err := service.GetConversationByID(conv.ID)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result == nil {
		t.Fatal("expected result, got nil")
	}

	if result.ID != conv.ID {
		t.Fatalf("expected ID %d, got %d", conv.ID, result.ID)
	}
}

func TestOrganizationConversationService_AssignConversation_Success(t *testing.T) {
	tx := SetupTestDB(t)
	service := impl.NewConversationService(tx)

	org, owner := CreateTestOrganizationWithOwner(tx, t, "Test Org")

	guestRole, _ := GetOrCreateRole(tx, models.RoleGuest)
	guest := models.UserModel{
		Email:    "guest@test.com",
		Name:     "Guest",
		Password: "password",
		RoleID:   guestRole.ID,
	}
	tx.Create(&guest)

	conv := models.ConversationModel{
		OrganizationID: org.ID,
		GuestID:        guest.ID,
		Status:         models.ConversationStatusPending,
	}
	tx.Create(&conv)

	req := requestdto.AssignConversationRequest{
		OrganizationStaffID: owner.ID,
	}

	result, err := service.AssignConversation(conv.ID, req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result == nil {
		t.Fatal("expected result, got nil")
	}

	if result.Status != models.ConversationStatusInProgress {
		t.Fatalf("expected status %s, got %s", models.ConversationStatusInProgress, result.Status)
	}
}

func TestOrganizationConversationService_UpdateConversationStatus_Success(t *testing.T) {
	tx := SetupTestDB(t)
	service := impl.NewConversationService(tx)

	org, _ := CreateTestOrganizationWithOwner(tx, t, "Test Org")

	guestRole, _ := GetOrCreateRole(tx, models.RoleGuest)
	guest := models.UserModel{
		Email:    "guest@test.com",
		Name:     "Guest",
		Password: "password",
		RoleID:   guestRole.ID,
	}
	tx.Create(&guest)

	conv := models.ConversationModel{
		OrganizationID: org.ID,
		GuestID:        guest.ID,
		Status:         models.ConversationStatusPending,
	}
	tx.Create(&conv)

	req := requestdto.UpdateConversationRequest{
		Status: models.ConversationStatusDone,
	}

	err := service.UpdateConversationStatus(conv.ID, req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	var updated models.ConversationModel
	tx.First(&updated, conv.ID)

	if updated.Status != models.ConversationStatusDone {
		t.Fatalf("expected status %s, got %s", models.ConversationStatusDone, updated.Status)
	}
}
