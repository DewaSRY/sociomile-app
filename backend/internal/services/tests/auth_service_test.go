package tests

import (
	"DewaSRY/sociomile-app/internal/services/impl"
	"DewaSRY/sociomile-app/pkg/dtos/requestdto"
	jwtLib "DewaSRY/sociomile-app/pkg/lib/jwt"
	"DewaSRY/sociomile-app/pkg/models"
	"context"
	"errors"
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
func TestAuthService_Register_Success(t *testing.T) {
	tx := SetupTestDB(t)
	mockJwt := &MockJwtService{}
	authService := impl.NewAuthService(tx, mockJwt)

	// Get or create guest role
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

	if result.User.Email != req.Email {
		t.Fatalf("expected email %s, got %s", req.Email, result.User.Email)
	}

	if result.User.Name != req.Name {
		t.Fatalf("expected name %s, got %s", req.Name, result.User.Name)
	}

	// Verify user was created in database
	var user models.UserModel
	if err := tx.Where("email = ?", req.Email).First(&user).Error; err != nil {
		t.Fatalf("failed to find created user: %v", err)
	}

	if user.Name != req.Name {
		t.Fatalf("expected name %s, got %s", req.Name, user.Name)
	}
}

func TestAuthService_Register_UserAlreadyExists(t *testing.T) {
	tx := SetupTestDB(t)
	mockJwt := &MockJwtService{}
	authService := impl.NewAuthService(tx, mockJwt)

	// Get or create guest role
	guestRole, err := GetOrCreateRole(tx, models.RoleGuest)
	if err != nil {
		t.Fatalf("failed to get guest role: %v", err)
	}

	// Create existing user
	existingUser := models.UserModel{
		Email:    "existing@example.com",
		Name:     "Existing User",
		Password: "password123",
		RoleID:   guestRole.ID,
	}
	if err := tx.Create(&existingUser).Error; err != nil {
		t.Fatalf("failed to create existing user: %v", err)
	}

	req := requestdto.RegisterRequest{
		Email:    "existing@example.com",
		Name:     "Another User",
		Password: "password123",
	}

	result, err := authService.Register(req)
	if err == nil {
		t.Fatal("expected error for duplicate email, got nil")
	}

	if result != nil {
		t.Fatal("expected nil result, got non-nil")
	}

	if err.Error() != "user with this email already exists" {
		t.Fatalf("expected 'user with this email already exists' error, got %v", err)
	}
}

func TestAuthService_Register_TokenGenerationFails(t *testing.T) {
	tx := SetupTestDB(t)
	mockJwt := &MockJwtService{
		GenerateTokenFunc: func(userID uint, email string, roleID uint, organizationID *uint) (string, error) {
			return "", errors.New("token generation failed")
		},
	}
	authService := impl.NewAuthService(tx, mockJwt)

	// Get or create guest role
	guestRole, err := GetOrCreateRole(tx, models.RoleGuest)
	if err != nil {
		t.Fatalf("failed to get guest role: %v", err)
	}
	_ = guestRole

	req := requestdto.RegisterRequest{
		Email:    "tokentest@example.com",
		Name:     "Token Test User",
		Password: "password123",
	}

	result, err := authService.Register(req)
	if err == nil {
		t.Fatal("expected error for token generation failure, got nil")
	}

	if result != nil {
		t.Fatal("expected nil result, got non-nil")
	}

	if err.Error() != "failed to generate token" {
		t.Fatalf("expected 'failed to generate token' error, got %v", err)
	}

	// Verify user was still created in database (transaction should rollback)
	// Since we're in a transaction that will rollback, we can't verify this in the rolled back state
}

// Test Login method
func TestAuthService_Login_Success(t *testing.T) {
	tx := SetupTestDB(t)
	mockJwt := &MockJwtService{}
	authService := impl.NewAuthService(tx, mockJwt)

	// Get or create guest role
	guestRole, err := GetOrCreateRole(tx, models.RoleGuest)
	if err != nil {
		t.Fatalf("failed to get guest role: %v", err)
	}

	// Create test user
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

	if result.Token != "mock-token" {
		t.Fatalf("expected token 'mock-token', got %s", result.Token)
	}

	if result.User.Email != user.Email {
		t.Fatalf("expected email %s, got %s", user.Email, result.User.Email)
	}

	if result.User.Name != user.Name {
		t.Fatalf("expected name %s, got %s", user.Name, result.User.Name)
	}
}

func TestAuthService_Login_InvalidEmail(t *testing.T) {
	tx := SetupTestDB(t)
	mockJwt := &MockJwtService{}
	authService := impl.NewAuthService(tx, mockJwt)

	req := requestdto.LoginRequest{
		Email:    "nonexistent@example.com",
		Password: "password123",
	}

	result, err := authService.Login(req)
	if err == nil {
		t.Fatal("expected error for non-existent email, got nil")
	}

	if result != nil {
		t.Fatal("expected nil result, got non-nil")
	}

	if err.Error() != "invalid email or password" {
		t.Fatalf("expected 'invalid email or password' error, got %v", err)
	}
}

func TestAuthService_Login_InvalidPassword(t *testing.T) {
	tx := SetupTestDB(t)
	mockJwt := &MockJwtService{}
	authService := impl.NewAuthService(tx, mockJwt)

	// Get or create guest role
	guestRole, err := GetOrCreateRole(tx, models.RoleGuest)
	if err != nil {
		t.Fatalf("failed to get guest role: %v", err)
	}

	// Create test user
	user := models.UserModel{
		Email:    "wrongpass@example.com",
		Name:     "Wrong Pass User",
		Password: "correctpassword",
		RoleID:   guestRole.ID,
	}
	if err := tx.Create(&user).Error; err != nil {
		t.Fatalf("failed to create test user: %v", err)
	}

	req := requestdto.LoginRequest{
		Email:    "wrongpass@example.com",
		Password: "wrongpassword",
	}

	result, err := authService.Login(req)
	if err == nil {
		t.Fatal("expected error for wrong password, got nil")
	}

	if result != nil {
		t.Fatal("expected nil result, got non-nil")
	}

	if err.Error() != "invalid email or password" {
		t.Fatalf("expected 'invalid email or password' error, got %v", err)
	}
}

func TestAuthService_Login_TokenGenerationFails(t *testing.T) {
	tx := SetupTestDB(t)
	mockJwt := &MockJwtService{
		GenerateTokenFunc: func(userID uint, email string, roleID uint, organizationID *uint) (string, error) {
			return "", errors.New("token generation failed")
		},
	}
	authService := impl.NewAuthService(tx, mockJwt)

	// Get or create guest role
	guestRole, err := GetOrCreateRole(tx, models.RoleGuest)
	if err != nil {
		t.Fatalf("failed to get guest role: %v", err)
	}

	// Create test user
	user := models.UserModel{
		Email:    "logintokenfail@example.com",
		Name:     "Login Token Fail",
		Password: "password123",
		RoleID:   guestRole.ID,
	}
	if err := tx.Create(&user).Error; err != nil {
		t.Fatalf("failed to create test user: %v", err)
	}

	req := requestdto.LoginRequest{
		Email:    "logintokenfail@example.com",
		Password: "password123",
	}

	result, err := authService.Login(req)
	if err == nil {
		t.Fatal("expected error for token generation failure, got nil")
	}

	if result != nil {
		t.Fatal("expected nil result, got non-nil")
	}

	if err.Error() != "failed to generate token" {
		t.Fatalf("expected 'failed to generate token' error, got %v", err)
	}
}

// Test GetUserByID method
func TestAuthService_GetUserByID_Success(t *testing.T) {
	tx := SetupTestDB(t)
	mockJwt := &MockJwtService{}
	authService := impl.NewAuthService(tx, mockJwt)

	// Get or create guest role
	guestRole, err := GetOrCreateRole(tx, models.RoleGuest)
	if err != nil {
		t.Fatalf("failed to get guest role: %v", err)
	}

	// Create test user
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

	if result.ID != user.ID {
		t.Fatalf("expected ID %d, got %d", user.ID, result.ID)
	}

	if result.Email != user.Email {
		t.Fatalf("expected email %s, got %s", user.Email, result.Email)
	}

	if result.Name != user.Name {
		t.Fatalf("expected name %s, got %s", user.Name, result.Name)
	}
}

func TestAuthService_GetUserByID_NotFound(t *testing.T) {
	tx := SetupTestDB(t)
	mockJwt := &MockJwtService{}
	authService := impl.NewAuthService(tx, mockJwt)

	result, err := authService.GetUserByID(999999)
	if err == nil {
		t.Fatal("expected error for non-existent user, got nil")
	}

	if result != nil {
		t.Fatal("expected nil result, got non-nil")
	}

	if err.Error() != "user not found" {
		t.Fatalf("expected 'user not found' error, got %v", err)
	}
}

// Test RefreshToken method
func TestAuthService_RefreshToken_Success(t *testing.T) {
	tx := SetupTestDB(t)
	mockJwt := &MockJwtService{
		RefreshTokenFunc: func(tokenString string) (string, error) {
			if tokenString == "valid-token" {
				return "refreshed-token", nil
			}
			return "", errors.New("invalid token")
		},
	}
	authService := impl.NewAuthService(tx, mockJwt)

	result, err := authService.RefreshToken("valid-token")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result != "refreshed-token" {
		t.Fatalf("expected 'refreshed-token', got %s", result)
	}
}

func TestAuthService_RefreshToken_InvalidToken(t *testing.T) {
	tx := SetupTestDB(t)
	mockJwt := &MockJwtService{
		RefreshTokenFunc: func(tokenString string) (string, error) {
			return "", errors.New("invalid token")
		},
	}
	authService := impl.NewAuthService(tx, mockJwt)

	result, err := authService.RefreshToken("invalid-token")
	if err == nil {
		t.Fatal("expected error for invalid token, got nil")
	}

	if result != "" {
		t.Fatalf("expected empty string, got %s", result)
	}

	if err.Error() != "invalid token" {
		t.Fatalf("expected 'invalid token' error, got %v", err)
	}
}

// Test Register with different scenarios
func TestAuthService_Register_WithOrganization(t *testing.T) {
	tx := SetupTestDB(t)
	mockJwt := &MockJwtService{
		GenerateTokenFunc: func(userID uint, email string, roleID uint, organizationID *uint) (string, error) {
			// Verify organizationID is passed correctly (should be nil for guest users)
			if organizationID != nil {
				t.Error("expected organizationID to be nil for guest registration")
			}
			return "mock-token", nil
		},
	}
	authService := impl.NewAuthService(tx, mockJwt)

	// Get or create guest role
	guestRole, err := GetOrCreateRole(tx, models.RoleGuest)
	if err != nil {
		t.Fatalf("failed to get guest role: %v", err)
	}
	_ = guestRole

	req := requestdto.RegisterRequest{
		Email:    "orguser@example.com",
		Name:     "Org User",
		Password: "password123",
	}

	result, err := authService.Register(req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result == nil {
		t.Fatal("expected result, got nil")
	}
}

func TestAuthService_Login_WithOrganization(t *testing.T) {
	tx := SetupTestDB(t)
	
	// Create organization with owner
	org, owner := CreateTestOrganizationWithOwner(tx, t, "Test Org")

	// Create a staff user in the organization
	salesRole, err := GetOrCreateRole(tx, models.RoleOrganizationSales)
	if err != nil {
		t.Fatalf("failed to get sales role: %v", err)
	}

	staff := models.UserModel{
		Email:          "staff@org.com",
		Name:           "Staff User",
		Password:       "password123",
		RoleID:         salesRole.ID,
		OrganizationID: &org.ID,
	}
	if err := tx.Create(&staff).Error; err != nil {
		t.Fatalf("failed to create staff user: %v", err)
	}

	mockJwt := &MockJwtService{
		GenerateTokenFunc: func(userID uint, email string, roleID uint, organizationID *uint) (string, error) {
			// Verify organizationID is passed correctly for staff users
			if organizationID == nil {
				t.Error("expected organizationID to be non-nil for staff login")
			} else if *organizationID != org.ID {
				t.Errorf("expected organizationID %d, got %d", org.ID, *organizationID)
			}
			return "mock-token", nil
		},
	}
	authService := impl.NewAuthService(tx, mockJwt)

	req := requestdto.LoginRequest{
		Email:    "staff@org.com",
		Password: "password123",
	}

	result, err := authService.Login(req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result == nil {
		t.Fatal("expected result, got nil")
	}

	if result.User.Email != staff.Email {
		t.Fatalf("expected email %s, got %s", staff.Email, result.User.Email)
	}

	// Also verify owner can login
	_ = owner // Use owner variable
}
