package tests

import (
	"DewaSRY/sociomile-app/internal/services/impl"
	"testing"
)



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

