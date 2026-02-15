package tests

import (
	"DewaSRY/sociomile-app/internal/services/impl"
	"DewaSRY/sociomile-app/pkg/dtos/requestdto"
	jwtLib "DewaSRY/sociomile-app/pkg/lib/jwt"
	"DewaSRY/sociomile-app/pkg/models"
	"context"
	"testing"
)

// MockJwtService implements jwtLib.JwtService for testing
type MockJwtService struct {
	GenerateTokenFunc func(userID uint, email string, roleID uint, organizationID *uint) (string, error)
	RefreshTokenFunc  func(tokenString string) (string, error)
}

func (m *MockJwtService) GenerateToken(userID uint, email string, roleID uint, organizationID *uint) (string, error) {
	if m.GenerateTokenFunc != nil {
		return m.GenerateTokenFunc(userID, email, roleID, organizationID)
	}
	return "mock-token", nil
}

func (m *MockJwtService) ValidateToken(tokenString string) (*jwtLib.Claims, error) {
	return nil, nil
}

func (m *MockJwtService) RefreshToken(tokenString string) (string, error) {
	if m.RefreshTokenFunc != nil {
		return m.RefreshTokenFunc(tokenString)
	}
	return "new-mock-token", nil
}

func (m *MockJwtService) GetUserFromContext(ctx context.Context) (*jwtLib.Claims, bool) {
	return nil, false
}

// Test Register method
func TestAuthService_Register(t *testing.T) {
	tx := SetupTestDB(t)
	mockJwt := &MockJwtService{}
	authService := impl.NewAuthService(tx, mockJwt)

	guestRole, err := GetOrCreateRole(tx, models.RoleGuest)
	if err != nil {
		t.Fatalf("failed to get guest role: %v", err)
	}
	_ = guestRole

	req := requestdto.RegisterRequest{
		Email:    "newuser@example.com",
		Name:     "New User",
		Password: "password123",
	}

	result, err := authService.Register(req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result == nil {
		t.Fatal("expected result, got nil")
	}

	if result.Token != "mock-token" {
		t.Fatalf("expected token 'mock-token', got %s", result.Token)
	}
}

// Test Login method
func TestAuthService_Login(t *testing.T) {
	tx := SetupTestDB(t)
	mockJwt := &MockJwtService{}
	authService := impl.NewAuthService(tx, mockJwt)

	guestRole, err := GetOrCreateRole(tx, models.RoleGuest)
	if err != nil {
		t.Fatalf("failed to get guest role: %v", err)
	}

	user := models.UserModel{
		Email:    "login@example.com",
		Name:     "Login User",
		Password: "password123",
		RoleID:   guestRole.ID,
	}
	if err := tx.Create(&user).Error; err != nil {
		t.Fatalf("failed to create test user: %v", err)
	}

	req := requestdto.LoginRequest{
		Email:    "login@example.com",
		Password: "password123",
	}

	result, err := authService.Login(req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result == nil {
		t.Fatal("expected result, got nil")
	}
}

// Test GetUserByID method
func TestAuthService_GetUserByID(t *testing.T) {
	tx := SetupTestDB(t)
	mockJwt := &MockJwtService{}
	authService := impl.NewAuthService(tx, mockJwt)

	guestRole, err := GetOrCreateRole(tx, models.RoleGuest)
	if err != nil {
		t.Fatalf("failed to get guest role: %v", err)
	}

	user := models.UserModel{
		Email:    "getbyid@example.com",
		Name:     "Get By ID User",
		Password: "password123",
		RoleID:   guestRole.ID,
	}
	if err := tx.Create(&user).Error; err != nil {
		t.Fatalf("failed to create test user: %v", err)
	}

	result, err := authService.GetUserByID(user.ID)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result == nil {
		t.Fatal("expected result, got nil")
	}
}

// Test RefreshToken method
func TestAuthService_RefreshToken(t *testing.T) {
	tx := SetupTestDB(t)
	mockJwt := &MockJwtService{}
	authService := impl.NewAuthService(tx, mockJwt)

	result, err := authService.RefreshToken("valid-token")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result == "" {
		t.Fatal("expected token, got empty string")
	}
}
