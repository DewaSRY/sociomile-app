package tests

import (
	"DewaSRY/sociomile-app/internal/services/impl"
	"DewaSRY/sociomile-app/pkg/dtos/requestdto"
	jwtLib "DewaSRY/sociomile-app/pkg/lib/jwt"
	"testing"
)


func TestCreateStaff_Success(t *testing.T) {
	tx := SetupTestDB(t)

	orgService := impl.NewOrganizationService(tx)

	claims := &jwtLib.Claims{
		UserID:         1,
		Email: "example@test.com",
		RoleID: 1,
	}

	req := requestdto.RegisterRequest{
		Email:    "owner@example.com",
		Name:     "Staff User",
		Password: "password123",
	}

	err := orgService.CreateStaff(req, claims)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}