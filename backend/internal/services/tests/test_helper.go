package tests

import (
	"DewaSRY/sociomile-app/internal/config"
	"DewaSRY/sociomile-app/internal/database"
	"DewaSRY/sociomile-app/pkg/lib/logger"
	"DewaSRY/sociomile-app/pkg/models"
	"testing"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func SetupTestDB(t *testing.T) *gorm.DB {
	godotenv.Load("../../../.env")
	config.Load()
	logger.Init()
	db := database.Connect()

	tx := db.Begin()
	if tx.Error != nil {
		t.Fatalf("failed to start transaction: %v", tx.Error)
	}

	t.Cleanup(func() {
		tx.Rollback()
	})

	return tx
}

func GetOrCreateRole(tx *gorm.DB, roleName string) (*models.UserRoleModel, error) {
	var role models.UserRoleModel
	err := tx.Where("name = ?", roleName).FirstOrCreate(&role, models.UserRoleModel{
		Name: roleName,
	}).Error
	return &role, err
}

func CreateTestOrganizationWithOwner(tx *gorm.DB, t *testing.T, orgName string) (*models.OrganizationModel, *models.UserModel) {
	ownerRole, err := GetOrCreateRole(tx, models.RoleOrganizationOwner)
	if err != nil {
		t.Fatalf("failed to get or create owner role: %v", err)
	}

	owner := models.UserModel{
		Email:    "owner_" + orgName + "@example.com",
		Name:     "Owner of " + orgName,
		Password: "password123",
		RoleID:   ownerRole.ID,
	}
	if err := tx.Create(&owner).Error; err != nil {
		t.Fatalf("failed to create owner user: %v", err)
	}

	org := models.OrganizationModel{
		Name:    orgName,
		OwnerID: owner.ID,
	}
	if err := tx.Create(&org).Error; err != nil {
		t.Fatalf("failed to create organization: %v", err)
	}

	owner.OrganizationID = &org.ID
	if err := tx.Save(&owner).Error; err != nil {
		t.Fatalf("failed to update owner's organization: %v", err)
	}

	return &org, &owner
}
