package tests

import (
	"DewaSRY/sociomile-app/internal/services/impl"
	"DewaSRY/sociomile-app/pkg/dtos/filtersdto"
	"DewaSRY/sociomile-app/pkg/dtos/requestdto"
	jwtLib "DewaSRY/sociomile-app/pkg/lib/jwt"
	"DewaSRY/sociomile-app/pkg/models"
	"testing"
)

func TestGuestConversation_CreateConversation_Success(t *testing.T) {
	tx := SetupTestDB(t)
	guestConversationService := impl.NewGuestConversationService(tx)

	// Create test organization
	org, _ := CreateTestOrganizationWithOwner(tx, t, "Test Organization")

	// Get or create guest role
	guestRole, err := GetOrCreateRole(tx, models.RoleGuest)
	if err != nil {
		t.Fatalf("failed to get guest role: %v", err)
	}

	// Create test guest user
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

	// Verify conversation was created
	var conversation models.ConversationModel
	if err := tx.Where("guest_id = ? AND organization_id = ?", guest.ID, org.ID).First(&conversation).Error; err != nil {
		t.Fatalf("failed to find created conversation: %v", err)
	}

	if conversation.Status != models.ConversationStatusPending {
		t.Fatalf("expected status %s, got %s", models.ConversationStatusPending, conversation.Status)
	}
}

func TestGuestConversation_CreateConversation_OrganizationNotFound(t *testing.T) {
	tx := SetupTestDB(t)
	guestConversationService := impl.NewGuestConversationService(tx)

	// Get or create guest role
	guestRole, err := GetOrCreateRole(tx, models.RoleGuest)
	if err != nil {
		t.Fatalf("failed to get guest role: %v", err)
	}

	// Create test guest user
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
		OrganizationID: 999, // Non-existent organization
	}

	err = guestConversationService.CreateConversation(claims, req)
	if err == nil {
		t.Fatal("expected error for non-existent organization, got nil")
	}

	if err.Error() != "organization not found" {
		t.Fatalf("expected 'organization not found' error, got %v", err)
	}
}

func TestGuestConversation_GetConversation_Success(t *testing.T) {
	tx := SetupTestDB(t)
	guestConversationService := impl.NewGuestConversationService(tx)

	// Create test organization
	org, _ := CreateTestOrganizationWithOwner(tx, t, "Test Organization")

	// Get or create guest role
	guestRole, err := GetOrCreateRole(tx, models.RoleGuest)
	if err != nil {
		t.Fatalf("failed to get guest role: %v", err)
	}

	// Create test guest user
	guest := models.UserModel{
		Email:    "guest@example.com",
		Name:     "Guest User",
		Password: "password123",
		RoleID:   guestRole.ID,
	}
	if err := tx.Create(&guest).Error; err != nil {
		t.Fatalf("failed to create guest user: %v", err)
	}

	// Create test conversations
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

	if len(result.Data) != 2 {
		t.Fatalf("expected 2 conversations, got %d", len(result.Data))
	}

	if result.Metadata.Total != 2 {
		t.Fatalf("expected total 2, got %d", result.Metadata.Total)
	}

	if result.Metadata.Page != 1 {
		t.Fatalf("expected page 1, got %d", result.Metadata.Page)
	}

	if result.Metadata.Limit != 10 {
		t.Fatalf("expected limit 10, got %d", result.Metadata.Limit)
	}
}

func TestGuestConversation_GetConversation_EmptyResult(t *testing.T) {
	tx := SetupTestDB(t)
	guestConversationService := impl.NewGuestConversationService(tx)

	// Get or create guest role
	guestRole, err := GetOrCreateRole(tx, models.RoleGuest)
	if err != nil {
		t.Fatalf("failed to get guest role: %v", err)
	}

	// Create test guest user without conversations
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

	if len(result.Data) != 0 {
		t.Fatalf("expected 0 conversations, got %d", len(result.Data))
	}

	if result.Metadata.Total != 0 {
		t.Fatalf("expected total 0, got %d", result.Metadata.Total)
	}
}

func TestGuestConversation_GetConversation_WithPagination(t *testing.T) {
	tx := SetupTestDB(t)
	guestConversationService := impl.NewGuestConversationService(tx)

	// Create test organization
	org, _ := CreateTestOrganizationWithOwner(tx, t, "Test Organization")

	// Get or create guest role
	guestRole, err := GetOrCreateRole(tx, models.RoleGuest)
	if err != nil {
		t.Fatalf("failed to get guest role: %v", err)
	}

	// Create test guest user
	guest := models.UserModel{
		Email:    "guest@example.com",
		Name:     "Guest User",
		Password: "password123",
		RoleID:   guestRole.ID,
	}
	if err := tx.Create(&guest).Error; err != nil {
		t.Fatalf("failed to create guest user: %v", err)
	}

	// Create 5 test conversations
	for i := 0; i < 5; i++ {
		conv := models.ConversationModel{
			OrganizationID: org.ID,
			GuestID:        guest.ID,
			Status:         models.ConversationStatusPending,
		}
		if err := tx.Create(&conv).Error; err != nil {
			t.Fatalf("failed to create test conversation: %v", err)
		}
	}

	claims := &jwtLib.Claims{
		UserID: guest.ID,
		Email:  guest.Email,
		RoleID: guestRole.ID,
	}

	// Test first page
	page := 1
	limit := 2
	filter := filtersdto.FiltersDto{
		Page:  &page,
		Limit: &limit,
	}

	result, err := guestConversationService.GetConversation(claims, filter)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result.Metadata.Total != 5 {
		t.Fatalf("expected total 5, got %d", result.Metadata.Total)
	}

	// Test second page
	page = 2
	filter.Page = &page
	result, err = guestConversationService.GetConversation(claims, filter)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(result.Data) == 0 {
		t.Fatal("expected at least some conversations on second page")
	}
}

func TestGuestConversation_GetConversation_WithOrganizationPreload(t *testing.T) {
	tx := SetupTestDB(t)
	guestConversationService := impl.NewGuestConversationService(tx)

	// Create test organization
	org, _ := CreateTestOrganizationWithOwner(tx, t, "Test Organization")

	// Get or create guest role
	guestRole, err := GetOrCreateRole(tx, models.RoleGuest)
	if err != nil {
		t.Fatalf("failed to get guest role: %v", err)
	}

	// Create test guest user
	guest := models.UserModel{
		Email:    "guest@example.com",
		Name:     "Guest User",
		Password: "password123",
		RoleID:   guestRole.ID,
	}
	if err := tx.Create(&guest).Error; err != nil {
		t.Fatalf("failed to create guest user: %v", err)
	}

	// Create test conversation
	conv := models.ConversationModel{
		OrganizationID: org.ID,
		GuestID:        guest.ID,
		Status:         models.ConversationStatusPending,
	}
	if err := tx.Create(&conv).Error; err != nil {
		t.Fatalf("failed to create test conversation: %v", err)
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

	if len(result.Data) != 1 {
		t.Fatalf("expected 1 conversation, got %d", len(result.Data))
	}

	// Verify organization is preloaded
	if result.Data[0].Organization == nil {
		t.Fatal("expected organization to be preloaded, got nil")
	}

	if result.Data[0].Organization.Name != org.Name {
		t.Fatalf("expected organization name %s, got %s", org.Name, result.Data[0].Organization.Name)
	}
}
