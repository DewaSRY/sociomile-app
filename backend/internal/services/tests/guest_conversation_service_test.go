package tests

import (
	"DewaSRY/sociomile-app/internal/services/impl"
	"DewaSRY/sociomile-app/pkg/dtos/filtersdto"
	"DewaSRY/sociomile-app/pkg/dtos/requestdto"
	jwtLib "DewaSRY/sociomile-app/pkg/lib/jwt"
	"DewaSRY/sociomile-app/pkg/models"
	"testing"
)

func TestGuestConversation_CreateConversation(t *testing.T) {
	tx := SetupTestDB(t)
	guestConversationService := impl.NewGuestConversationService(tx)

	org, _ := CreateTestOrganizationWithOwner(tx, t, "Test Organization")

	guestRole, err := GetOrCreateRole(tx, models.RoleGuest)
	if err != nil {
		t.Fatalf("failed to get guest role: %v", err)
	}

	guest := models.UserModel{
		Email:    "guest@example.com",
		Name:     "Guest User",
		Password: "password123",
		RoleID:   guestRole.ID,
	}
	if err := tx.Create(&guest).Error; err != nil {
		t.Fatalf("failed to create guest user: %v", err)
	}

	claims := &jwtLib.Claims{
		UserID: guest.ID,
		Email:  guest.Email,
		RoleID: guestRole.ID,
	}

	req := requestdto.CreateConversationRequest{
		OrganizationID: org.ID,
	}

	err = guestConversationService.CreateConversation(claims, req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestGuestConversation_GetConversation(t *testing.T) {
	tx := SetupTestDB(t)
	guestConversationService := impl.NewGuestConversationService(tx)

	org, _ := CreateTestOrganizationWithOwner(tx, t, "Test Organization")

	guestRole, err := GetOrCreateRole(tx, models.RoleGuest)
	if err != nil {
		t.Fatalf("failed to get guest role: %v", err)
	}

	guest := models.UserModel{
		Email:    "guest@example.com",
		Name:     "Guest User",
		Password: "password123",
		RoleID:   guestRole.ID,
	}
	if err := tx.Create(&guest).Error; err != nil {
		t.Fatalf("failed to create guest user: %v", err)
	}

	conversations := []models.ConversationModel{
		{
			OrganizationID: org.ID,
			GuestID:        guest.ID,
			Status:         models.ConversationStatusPending,
		},
		{
			OrganizationID: org.ID,
			GuestID:        guest.ID,
			Status:         models.ConversationStatusInProgress,
		},
	}
	for _, conv := range conversations {
		if err := tx.Create(&conv).Error; err != nil {
			t.Fatalf("failed to create test conversation: %v", err)
		}
	}

	claims := &jwtLib.Claims{
		UserID: guest.ID,
		Email:  guest.Email,
		RoleID: guestRole.ID,
	}

	page := 1
	limit := 10
	filter := filtersdto.FiltersDto{
		Page:  &page,
		Limit: &limit,
	}

	result, err := guestConversationService.GetConversation(claims, filter)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result == nil {
		t.Fatal("expected result, got nil")
	}
}
