package tests

import (
	"DewaSRY/sociomile-app/internal/services/impl"
	"DewaSRY/sociomile-app/pkg/dtos/filtersdto"
	"DewaSRY/sociomile-app/pkg/dtos/requestdto"
	"DewaSRY/sociomile-app/pkg/models"
	"testing"
)

func TestHubService_CreateOrganization(t *testing.T) {
	tx := SetupTestDB(t)
	service := impl.NewHubServiceImpl(tx)

	_, err := GetOrCreateRole(tx, models.RoleOrganizationOwner)
	if err != nil {
		t.Fatalf("failed to create role: %v", err)
	}

	req := requestdto.RegisterOrganizationRequest{
		Name:      "Test Organization",
		Email:     "owner@test.com",
		OwnerName: "Test Owner",
		Password:  "password123",
	}

	err = service.CreateOrganization(req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestHubService_GetOrganizationPagination(t *testing.T) {
	tx := SetupTestDB(t)
	service := impl.NewHubServiceImpl(tx)

	CreateTestOrganizationWithOwner(tx, t, "Org1")
	CreateTestOrganizationWithOwner(tx, t, "Org2")

	page := 1
	limit := 10
	filter := filtersdto.FiltersDto{
		Page:  &page,
		Limit: &limit,
	}

	result, err := service.GetOrganizationPagination(filter)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result == nil {
		t.Fatal("expected result, got nil")
	}
}

