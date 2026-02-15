package tests

import (
	"DewaSRY/sociomile-app/internal/services/impl"
	"DewaSRY/sociomile-app/pkg/dtos/filtersdto"
	"DewaSRY/sociomile-app/pkg/dtos/requestdto"
	jwtLib "DewaSRY/sociomile-app/pkg/lib/jwt"
	"DewaSRY/sociomile-app/pkg/models"
	"testing"
)

func TestGuestMessageService_SendConversationMessage(t *testing.T) {
	tx := SetupTestDB(t)
	messageService := impl.NewGuestMessageService(tx)

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

	conversation := models.ConversationModel{
		OrganizationID: org.ID,
		GuestID:        guest.ID,
		Status:         models.ConversationStatusPending,
	}
	if err := tx.Create(&conversation).Error; err != nil {
		t.Fatalf("failed to create conversation: %v", err)
	}

	claims := &jwtLib.Claims{
		UserID: guest.ID,
		Email:  guest.Email,
		RoleID: guestRole.ID,
	}

	req := requestdto.CreateConversationMessageRequest{
		ConversationID: conversation.ID,
		Message:        "Hello, this is a test message",
	}

	err = messageService.SendConversationMessage(claims, req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestGuestMessageService_GetConversationMessageList(t *testing.T) {
	tx := SetupTestDB(t)
	messageService := impl.NewGuestMessageService(tx)

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

	conversation := models.ConversationModel{
		OrganizationID: org.ID,
		GuestID:        guest.ID,
		Status:         models.ConversationStatusPending,
	}
	if err := tx.Create(&conversation).Error; err != nil {
		t.Fatalf("failed to create conversation: %v", err)
	}

	testMessages := []string{"Message 1", "Message 2", "Message 3"}
	for _, msg := range testMessages {
		message := models.ConversationMessageModel{
			OrganizationID: org.ID,
			ConversationID: conversation.ID,
			CreatedByID:    guest.ID,
			Message:        msg,
		}
		if err := tx.Create(&message).Error; err != nil {
			t.Fatalf("failed to create test message: %v", err)
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

	result, err := messageService.GetConversationMessageList(claims, filter, conversation.ID)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result == nil {
		t.Fatal("expected result, got nil")
	}
}
