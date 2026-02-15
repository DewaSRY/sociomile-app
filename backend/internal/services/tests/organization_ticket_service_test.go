package tests

import (
	"DewaSRY/sociomile-app/internal/services/impl"
	"DewaSRY/sociomile-app/pkg/dtos/filtersdto"
	"DewaSRY/sociomile-app/pkg/dtos/requestdto"
	jwtLib "DewaSRY/sociomile-app/pkg/lib/jwt"
	"DewaSRY/sociomile-app/pkg/models"
	"testing"
)

func TestOrganizationTicketService_CreateTicket(t *testing.T) {
	tx := SetupTestDB(t)
	service := impl.NewTicketService(tx)

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

	req := requestdto.CreateTicketRequest{
		ConversationID: conv.ID,
		Name:           "Test Ticket",
	}

	err := service.CreateTicket(claims, req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestOrganizationTicketService_GetTicketsList(t *testing.T) {
	tx := SetupTestDB(t)
	service := impl.NewTicketService(tx)

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

	ticket := models.TicketModel{
		OrganizationID: org.ID,
		ConversationID: conv.ID,
		CreatedByID:    owner.ID,
		TicketNumber:   "TEST-001",
		Name:           "Test Ticket",
		Status:         models.TicketStatusPending,
	}
	tx.Create(&ticket)

	claims := &jwtLib.Claims{
		UserID:         owner.ID,
		OrganizationId: &org.ID,
	}

	page := 1
	limit := 10
	filter := filtersdto.FiltersDto{Page: &page, Limit: &limit}

	result, err := service.GetTicketsList(claims, filter)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result == nil {
		t.Fatal("expected result, got nil")
	}
}

func TestOrganizationTicketService_UpdateTicket(t *testing.T) {
	tx := SetupTestDB(t)
	service := impl.NewTicketService(tx)

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

	ticket := models.TicketModel{
		OrganizationID: org.ID,
		ConversationID: conv.ID,
		CreatedByID:    owner.ID,
		TicketNumber:   "TEST-001",
		Name:           "Test Ticket",
		Status:         models.TicketStatusPending,
	}
	tx.Create(&ticket)

	claims := &jwtLib.Claims{
		UserID:         owner.ID,
		OrganizationId: &org.ID,
	}

	req := requestdto.UpdateTicketRequest{
		Name:   "Updated Ticket",
		Status: models.TicketStatusInProgress,
	}

	err := service.UpdateTicket(claims, ticket.ID, req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}
