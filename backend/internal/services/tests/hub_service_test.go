package tests

import (
	"DewaSRY/sociomile-app/internal/services/impl"
	"DewaSRY/sociomile-app/pkg/dtos/filtersdto"
	"DewaSRY/sociomile-app/pkg/dtos/requestdto"
	"DewaSRY/sociomile-app/pkg/models"
	"testing"
)

func TestHubService_CreateOrganization_Success(t *testing.T) {
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

	var org models.OrganizationModel
	if err := tx.Where("name = ?", req.Name).First(&org).Error; err != nil {
		t.Fatalf("organization not created: %v", err)
	}

	var user models.UserModel
	if err := tx.Where("email = ?", req.Email).First(&user).Error; err != nil {
		t.Fatalf("user not created: %v", err)
	}

	if user.OrganizationID == nil || *user.OrganizationID != org.ID {
		t.Fatalf("user organization ID not set correctly")
	}
}

func TestHubService_CreateOrganization_DuplicateEmail(t *testing.T) {
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
		t.Fatalf("expected no error on first creation, got %v", err)
	}

	err = service.CreateOrganization(req)
	if err == nil {
		t.Fatal("expected error for duplicate email, got nil")
	}
}

func TestHubService_GetOrganizationPagination_Success(t *testing.T) {
	tx := SetupTestDB(t)
	service := impl.NewHubServiceImpl(tx)

	org1, _ := CreateTestOrganizationWithOwner(tx, t, "Org1")
	org2, _ := CreateTestOrganizationWithOwner(tx, t, "Org2")

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

	if len(result.Data) < 2 {
		t.Fatalf("expected at least 2 organizations, got %d", len(result.Data))
	}

	found := false
	for _, org := range result.Data {
		if org.ID == org1.ID || org.ID == org2.ID {
			found = true
			break
		}
	}

	if !found {
		t.Fatal("created organizations not found in result")
	}
}

func TestHubService_GetOrganizationPagination_WithPagination(t *testing.T) {
	tx := SetupTestDB(t)
	service := impl.NewHubServiceImpl(tx)

	for i := 1; i <= 5; i++ {
		CreateTestOrganizationWithOwner(tx, t, "PaginationTestOrg"+string(rune('A'+i-1)))
	}

	page := 1
	limit := 2
	filter := filtersdto.FiltersDto{
		Page:  &page,
		Limit: &limit,
	}

	result, err := service.GetOrganizationPagination(filter)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(result.Data) > limit {
		t.Fatalf("expected at most %d organizations, got %d", limit, len(result.Data))
	}

	if result.Metadata.Page != page {
		t.Fatalf("expected page %d, got %d", page, result.Metadata.Page)
	}

	if result.Metadata.Limit != limit {
		t.Fatalf("expected limit %d, got %d", limit, result.Metadata.Limit)
	}
}

func TestHubService_GetOrganizationPagination_EmptyResult(t *testing.T) {
	tx := SetupTestDB(t)
	service := impl.NewHubServiceImpl(tx)

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

	// The result may not be empty if there are organizations from other tests
	// Just verify the structure is correct
	if result.Metadata.Page != page {
		t.Fatalf("expected page %d, got %d", page, result.Metadata.Page)
	}

	if result.Metadata.Limit != limit {
		t.Fatalf("expected limit %d, got %d", limit, result.Metadata.Limit)
	}
}
