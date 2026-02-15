package tests

import (
	"DewaSRY/sociomile-app/internal/services/impl"
	"DewaSRY/sociomile-app/pkg/dtos/filtersdto"
	"DewaSRY/sociomile-app/pkg/dtos/requestdto"
	jwtLib "DewaSRY/sociomile-app/pkg/lib/jwt"
	"DewaSRY/sociomile-app/pkg/models"
	"testing"
)

func TestOrganizationService_CreateStaff_Success(t *testing.T) {
	tx := SetupTestDB(t)
	orgService := impl.NewOrganizationService(tx)

	// Setup test data
	org, owner := CreateTestOrganizationWithOwner(tx, t, "Test Org")

	// Get or create sales role
	salesRole, err := GetOrCreateRole(tx, models.RoleOrganizationSales)
	if err != nil {
		t.Fatalf("failed to get sales role: %v", err)
	}

	_ = salesRole // Reference to avoid unused variable warning

	claims := &jwtLib.Claims{
		UserID:         owner.ID,
		Email:          owner.Email,
		RoleID:         owner.RoleID,
		OrganizationId: &org.ID,
	}

	req := requestdto.RegisterRequest{
		Email:    "staff@example.com",
		Name:     "Staff User",
		Password: "password123",
	}

	err = orgService.CreateStaff(req, claims)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Verify staff was created
	var staff models.UserModel
	if err := tx.Where("email = ?", req.Email).First(&staff).Error; err != nil {
		t.Fatalf("failed to find created staff: %v", err)
	}

	if staff.Name != req.Name {
		t.Fatalf("expected name %s, got %s", req.Name, staff.Name)
	}

	if staff.OrganizationID == nil || *staff.OrganizationID != org.ID {
		t.Fatalf("expected organization ID %d, got %v", org.ID, staff.OrganizationID)
	}
}

func TestOrganizationService_CreateStaff_UserAlreadyExists(t *testing.T) {
	tx := SetupTestDB(t)
	orgService := impl.NewOrganizationService(tx)

	// Setup test data
	org, owner := CreateTestOrganizationWithOwner(tx, t, "Test Org")

	// Get or create sales role
	salesRole, err := GetOrCreateRole(tx, models.RoleOrganizationSales)
	if err != nil {
		t.Fatalf("failed to get sales role: %v", err)
	}

	// Create existing user
	existingUser := models.UserModel{
		Email:          "existing@example.com",
		Name:           "Existing User",
		Password:       "password123",
		RoleID:         salesRole.ID,
		OrganizationID: &org.ID,
	}
	if err := tx.Create(&existingUser).Error; err != nil {
		t.Fatalf("failed to create existing user: %v", err)
	}

	claims := &jwtLib.Claims{
		UserID:         owner.ID,
		Email:          owner.Email,
		RoleID:         owner.RoleID,
		OrganizationId: &org.ID,
	}

	req := requestdto.RegisterRequest{
		Email:    "existing@example.com",
		Name:     "Another User",
		Password: "password123",
	}

	err = orgService.CreateStaff(req, claims)
	if err == nil {
		t.Fatal("expected error for duplicate email, got nil")
	}

	if err.Error() != "user with this email already exists" {
		t.Fatalf("expected 'user with this email already exists' error, got %v", err)
	}
}

func TestOrganizationService_CreateStaff_RoleNotFound(t *testing.T) {
	tx := SetupTestDB(t)

	// Setup test data
	org, owner := CreateTestOrganizationWithOwner(tx, t, "Test Org")

	claims := &jwtLib.Claims{
		UserID:         owner.ID,
		Email:          owner.Email,
		RoleID:         owner.RoleID,
		OrganizationId: &org.ID,
	}

	req := requestdto.RegisterRequest{
		Email:    "staff@example.com",
		Name:     "Staff User",
		Password: "password123",
	}

	// Create a separate transaction that doesn't have the role
	tx2 := tx.Begin()
	defer tx2.Rollback()
	
	// Delete the role in this sub-transaction
	var roleToDelete models.UserRoleModel
	if err := tx2.Where("name = ?", models.RoleOrganizationSales).First(&roleToDelete).Error; err == nil {
		// Update all users using this role to a different role first
		tx2.Exec("UPDATE users SET role_id = ? WHERE role_id = ?", owner.RoleID, roleToDelete.ID)
		// Now we can delete the role
		tx2.Exec("DELETE FROM user_roles WHERE id = ?", roleToDelete.ID)
	}
	
	orgService2 := impl.NewOrganizationService(tx2)
	
	err := orgService2.CreateStaff(req, claims)
	if err == nil {
		t.Fatal("expected error for missing role, got nil")
	}

	if err.Error() != "organization sales role not found" {
		t.Fatalf("expected 'organization sales role not found' error, got %v", err)
	}
}

func TestOrganizationService_GetStaffList_Success(t *testing.T) {
	tx := SetupTestDB(t)
	orgService := impl.NewOrganizationService(tx)

	// Setup test data
	org, owner := CreateTestOrganizationWithOwner(tx, t, "Test Org")

	// Get or create sales role
	salesRole, err := GetOrCreateRole(tx, models.RoleOrganizationSales)
	if err != nil {
		t.Fatalf("failed to get sales role: %v", err)
	}

	// Create staff users
	staffUsers := []models.UserModel{
		{
			Email:          "staff1@example.com",
			Name:           "Staff One",
			Password:       "password123",
			RoleID:         salesRole.ID,
			OrganizationID: &org.ID,
		},
		{
			Email:          "staff2@example.com",
			Name:           "Staff Two",
			Password:       "password123",
			RoleID:         salesRole.ID,
			OrganizationID: &org.ID,
		},
	}

	for _, user := range staffUsers {
		if err := tx.Create(&user).Error; err != nil {
			t.Fatalf("failed to create staff user: %v", err)
		}
	}

	claims := &jwtLib.Claims{
		UserID:         owner.ID,
		Email:          owner.Email,
		RoleID:         owner.RoleID,
		OrganizationId: &org.ID,
	}

	page := 1
	limit := 10
	filter := filtersdto.FiltersDto{
		Page:  &page,
		Limit: &limit,
	}

	result, err := orgService.GetStaffList(filter, claims)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result == nil {
		t.Fatal("expected result, got nil")
	}

	// Should have owner + 2 staff = 3 total
	if len(result.Data) != 3 {
		t.Fatalf("expected 3 staff members, got %d", len(result.Data))
	}

	if result.Metadata.Total != 3 {
		t.Fatalf("expected total 3, got %d", result.Metadata.Total)
	}
}

func TestOrganizationService_GetStaffList_EmptyResult(t *testing.T) {
	tx := SetupTestDB(t)
	orgService := impl.NewOrganizationService(tx)

	// Setup test data with no additional staff
	org, owner := CreateTestOrganizationWithOwner(tx, t, "Test Org")

	claims := &jwtLib.Claims{
		UserID:         owner.ID,
		Email:          owner.Email,
		RoleID:         owner.RoleID,
		OrganizationId: &org.ID,
	}

	page := 1
	limit := 10
	filter := filtersdto.FiltersDto{
		Page:  &page,
		Limit: &limit,
	}

	result, err := orgService.GetStaffList(filter, claims)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result == nil {
		t.Fatal("expected result, got nil")
	}

	// Should have only the owner
	if len(result.Data) != 1 {
		t.Fatalf("expected 1 staff member (owner), got %d", len(result.Data))
	}

	if result.Metadata.Total != 1 {
		t.Fatalf("expected total 1, got %d", result.Metadata.Total)
	}
}

func TestOrganizationService_GetStaffList_WithPagination(t *testing.T) {
	tx := SetupTestDB(t)
	orgService := impl.NewOrganizationService(tx)

	// Setup test data
	org, owner := CreateTestOrganizationWithOwner(tx, t, "Test Org")

	// Get or create sales role
	salesRole, err := GetOrCreateRole(tx, models.RoleOrganizationSales)
	if err != nil {
		t.Fatalf("failed to get sales role: %v", err)
	}

	// Create 5 staff users
	for i := 1; i <= 5; i++ {
		user := models.UserModel{
			Email:          "staff_" + string(rune(i+'0')) + "@example.com",
			Name:           "Staff " + string(rune(i+'0')),
			Password:       "password123",
			RoleID:         salesRole.ID,
			OrganizationID: &org.ID,
		}
		if err := tx.Create(&user).Error; err != nil {
			t.Fatalf("failed to create staff user: %v", err)
		}
	}

	claims := &jwtLib.Claims{
		UserID:         owner.ID,
		Email:          owner.Email,
		RoleID:         owner.RoleID,
		OrganizationId: &org.ID,
	}

	// Test first page with limit 2
	page := 1
	limit := 2
	filter := filtersdto.FiltersDto{
		Page:  &page,
		Limit: &limit,
	}

	result, err := orgService.GetStaffList(filter, claims)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(result.Data) != 2 {
		t.Fatalf("expected 2 staff members on first page, got %d", len(result.Data))
	}

	// Total should be owner + 5 staff = 6
	if result.Metadata.Total != 6 {
		t.Fatalf("expected total 6, got %d", result.Metadata.Total)
	}

	// Test second page
	page = 2
	filter.Page = &page
	result, err = orgService.GetStaffList(filter, claims)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(result.Data) != 2 {
		t.Fatalf("expected 2 staff members on second page, got %d", len(result.Data))
	}
}

func TestOrganizationService_GetStaffList_RolePreloaded(t *testing.T) {
	tx := SetupTestDB(t)
	orgService := impl.NewOrganizationService(tx)

	// Setup test data
	org, owner := CreateTestOrganizationWithOwner(tx, t, "Test Org")

	// Get or create sales role
	salesRole, err := GetOrCreateRole(tx, models.RoleOrganizationSales)
	if err != nil {
		t.Fatalf("failed to get sales role: %v", err)
	}

	// Create staff user
	user := models.UserModel{
		Email:          "staff@example.com",
		Name:           "Staff User",
		Password:       "password123",
		RoleID:         salesRole.ID,
		OrganizationID: &org.ID,
	}
	if err := tx.Create(&user).Error; err != nil {
		t.Fatalf("failed to create staff user: %v", err)
	}

	claims := &jwtLib.Claims{
		UserID:         owner.ID,
		Email:          owner.Email,
		RoleID:         owner.RoleID,
		OrganizationId: &org.ID,
	}

	page := 1
	limit := 10
	filter := filtersdto.FiltersDto{
		Page:  &page,
		Limit: &limit,
	}

	result, err := orgService.GetStaffList(filter, claims)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(result.Data) < 1 {
		t.Fatalf("expected at least 1 staff member, got %d", len(result.Data))
	}

	// Find the sales staff in results
	var salesStaff *models.UserModel
	for _, staff := range result.Data {
		if staff.Email == "staff@example.com" {
			// Check role name is preloaded
			if staff.RoleName != models.RoleOrganizationSales {
				t.Fatalf("expected role name %s, got %s", models.RoleOrganizationSales, staff.RoleName)
			}
			salesStaff = &models.UserModel{}
			break
		}
	}

	if salesStaff == nil {
		t.Fatal("created staff not found in results")
	}
}
