package tests

import (
	"DewaSRY/sociomile-app/internal/services/impl"
	"DewaSRY/sociomile-app/pkg/dtos/filtersdto"
	"DewaSRY/sociomile-app/pkg/dtos/requestdto"
	jwtLib "DewaSRY/sociomile-app/pkg/lib/jwt"
	"DewaSRY/sociomile-app/pkg/models"
	"testing"
)

// Test SendConversationMessage method
func TestGuestMessageService_SendConversationMessage_Success(t *testing.T) {
	tx := SetupTestDB(t)
	messageService := impl.NewGuestMessageService(tx)

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

	// Verify message was created
	var message models.ConversationMessageModel
	if err := tx.Where("conversation_id = ? AND message = ?", conversation.ID, req.Message).First(&message).Error; err != nil {
		t.Fatalf("failed to find created message: %v", err)
	}

	if message.CreatedByID != guest.ID {
		t.Fatalf("expected created_by_id %d, got %d", guest.ID, message.CreatedByID)
	}

	if message.OrganizationID != org.ID {
		t.Fatalf("expected organization_id %d, got %d", org.ID, message.OrganizationID)
	}

	if message.Message != req.Message {
		t.Fatalf("expected message %s, got %s", req.Message, message.Message)
	}
}

func TestGuestMessageService_SendConversationMessage_ConversationNotFound(t *testing.T) {
	tx := SetupTestDB(t)
	messageService := impl.NewGuestMessageService(tx)

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

	req := requestdto.CreateConversationMessageRequest{
		ConversationID: 999, // Non-existent conversation
		Message:        "This should fail",
	}

	err = messageService.SendConversationMessage(claims, req)
	if err == nil {
		t.Fatal("expected error for non-existent conversation, got nil")
	}

	if err.Error() != "conversation not found" {
		t.Fatalf("expected 'conversation not found' error, got %v", err)
	}
}

func TestGuestMessageService_SendConversationMessage_MultipleMessages(t *testing.T) {
	tx := SetupTestDB(t)
	messageService := impl.NewGuestMessageService(tx)

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

	// Send multiple messages
	messages := []string{"Message 1", "Message 2", "Message 3"}
	for _, msg := range messages {
		req := requestdto.CreateConversationMessageRequest{
			ConversationID: conversation.ID,
			Message:        msg,
		}

		err = messageService.SendConversationMessage(claims, req)
		if err != nil {
			t.Fatalf("expected no error for message '%s', got %v", msg, err)
		}
	}

	// Verify all messages were created
	var count int64
	if err := tx.Model(&models.ConversationMessageModel{}).
		Where("conversation_id = ?", conversation.ID).
		Count(&count).Error; err != nil {
		t.Fatalf("failed to count messages: %v", err)
	}

	if count != int64(len(messages)) {
		t.Fatalf("expected %d messages, got %d", len(messages), count)
	}
}

// Test GetConversationMessageList method
func TestGuestMessageService_GetConversationMessageList_Success(t *testing.T) {
	tx := SetupTestDB(t)
	messageService := impl.NewGuestMessageService(tx)

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
	conversation := models.ConversationModel{
		OrganizationID: org.ID,
		GuestID:        guest.ID,
		Status:         models.ConversationStatusPending,
	}
	if err := tx.Create(&conversation).Error; err != nil {
		t.Fatalf("failed to create conversation: %v", err)
	}

	// Create test messages
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

	if len(result.Data) != len(testMessages) {
		t.Fatalf("expected %d messages, got %d", len(testMessages), len(result.Data))
	}

	if result.Metadata.Total != len(testMessages) {
		t.Fatalf("expected total %d, got %d", len(testMessages), result.Metadata.Total)
	}

	// Verify messages are in order (ASC by created_at)
	for i, msg := range result.Data {
		if msg.Message != testMessages[i] {
			t.Fatalf("expected message %s at index %d, got %s", testMessages[i], i, msg.Message)
		}
	}
}

func TestGuestMessageService_GetConversationMessageList_EmptyResult(t *testing.T) {
	tx := SetupTestDB(t)
	messageService := impl.NewGuestMessageService(tx)

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

	// Create test conversation without messages
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

	if len(result.Data) != 0 {
		t.Fatalf("expected 0 messages, got %d", len(result.Data))
	}

	if result.Metadata.Total != 0 {
		t.Fatalf("expected total 0, got %d", result.Metadata.Total)
	}
}

func TestGuestMessageService_GetConversationMessageList_WithCreatedByPreload(t *testing.T) {
	tx := SetupTestDB(t)
	messageService := impl.NewGuestMessageService(tx)

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
	conversation := models.ConversationModel{
		OrganizationID: org.ID,
		GuestID:        guest.ID,
		Status:         models.ConversationStatusPending,
	}
	if err := tx.Create(&conversation).Error; err != nil {
		t.Fatalf("failed to create conversation: %v", err)
	}

	// Create test message
	message := models.ConversationMessageModel{
		OrganizationID: org.ID,
		ConversationID: conversation.ID,
		CreatedByID:    guest.ID,
		Message:        "Test message",
	}
	if err := tx.Create(&message).Error; err != nil {
		t.Fatalf("failed to create test message: %v", err)
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

	if len(result.Data) != 1 {
		t.Fatalf("expected 1 message, got %d", len(result.Data))
	}

	// Verify CreatedBy is preloaded
	if result.Data[0].CreatedBy == nil {
		t.Fatal("expected CreatedBy to be preloaded, got nil")
	}

	if result.Data[0].CreatedBy.ID != guest.ID {
		t.Fatalf("expected CreatedBy ID %d, got %d", guest.ID, result.Data[0].CreatedBy.ID)
	}

	if result.Data[0].CreatedBy.Email != guest.Email {
		t.Fatalf("expected CreatedBy email %s, got %s", guest.Email, result.Data[0].CreatedBy.Email)
	}

	if result.Data[0].CreatedBy.Name != guest.Name {
		t.Fatalf("expected CreatedBy name %s, got %s", guest.Name, result.Data[0].CreatedBy.Name)
	}
}

func TestGuestMessageService_GetConversationMessageList_OrderByCreatedAtASC(t *testing.T) {
	tx := SetupTestDB(t)
	messageService := impl.NewGuestMessageService(tx)

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
	conversation := models.ConversationModel{
		OrganizationID: org.ID,
		GuestID:        guest.ID,
		Status:         models.ConversationStatusPending,
	}
	if err := tx.Create(&conversation).Error; err != nil {
		t.Fatalf("failed to create conversation: %v", err)
	}

	// Create test messages (they should be ordered by creation time)
	testMessages := []string{"First", "Second", "Third"}
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

	if len(result.Data) != len(testMessages) {
		t.Fatalf("expected %d messages, got %d", len(testMessages), len(result.Data))
	}

	// Verify messages are ordered correctly (ASC)
	for i, msg := range result.Data {
		if msg.Message != testMessages[i] {
			t.Fatalf("expected message '%s' at index %d, got '%s'", testMessages[i], i, msg.Message)
		}
	}
}

func TestGuestMessageService_GetConversationMessageList_DifferentUsers(t *testing.T) {
	tx := SetupTestDB(t)
	messageService := impl.NewGuestMessageService(tx)

	// Create test organization
	org, owner := CreateTestOrganizationWithOwner(tx, t, "Test Organization")

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
	conversation := models.ConversationModel{
		OrganizationID: org.ID,
		GuestID:        guest.ID,
		Status:         models.ConversationStatusInProgress,
		OrganizationStaffID: &owner.ID,
	}
	if err := tx.Create(&conversation).Error; err != nil {
		t.Fatalf("failed to create conversation: %v", err)
	}

	// Create messages from both guest and staff
	guestMessage := models.ConversationMessageModel{
		OrganizationID: org.ID,
		ConversationID: conversation.ID,
		CreatedByID:    guest.ID,
		Message:        "Message from guest",
	}
	if err := tx.Create(&guestMessage).Error; err != nil {
		t.Fatalf("failed to create guest message: %v", err)
	}

	staffMessage := models.ConversationMessageModel{
		OrganizationID: org.ID,
		ConversationID: conversation.ID,
		CreatedByID:    owner.ID,
		Message:        "Message from staff",
	}
	if err := tx.Create(&staffMessage).Error; err != nil {
		t.Fatalf("failed to create staff message: %v", err)
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

	if len(result.Data) != 2 {
		t.Fatalf("expected 2 messages, got %d", len(result.Data))
	}

	// Verify both users' messages are included
	foundGuestMessage := false
	foundStaffMessage := false

	for _, msg := range result.Data {
		if msg.Message == "Message from guest" && msg.CreatedBy != nil && msg.CreatedBy.ID == guest.ID {
			foundGuestMessage = true
		}
		if msg.Message == "Message from staff" && msg.CreatedBy != nil && msg.CreatedBy.ID == owner.ID {
			foundStaffMessage = true
		}
	}

	if !foundGuestMessage {
		t.Fatal("guest message not found in results")
	}

	if !foundStaffMessage {
		t.Fatal("staff message not found in results")
	}
}

func TestGuestMessageService_SendConversationMessage_EmptyMessage(t *testing.T) {
	tx := SetupTestDB(t)
	messageService := impl.NewGuestMessageService(tx)

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
		Message:        "", // Empty message
	}

	// The service should still accept empty messages (no validation in current implementation)
	err = messageService.SendConversationMessage(claims, req)
	if err != nil {
		t.Fatalf("expected no error for empty message, got %v", err)
	}

	// Verify message was created
	var message models.ConversationMessageModel
	if err := tx.Where("conversation_id = ?", conversation.ID).First(&message).Error; err != nil {
		t.Fatalf("failed to find created message: %v", err)
	}

	if message.Message != "" {
		t.Fatalf("expected empty message, got %s", message.Message)
	}
}
