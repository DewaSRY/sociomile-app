package seeders

import (
	"fmt"
	"log"

	"DewaSRY/sociomile-app/internal/database"
	"DewaSRY/sociomile-app/pkg/models"
)

func ClearAllInitTables() error {
	db := database.DB
	log.Println("Starting to clear all tables...")

	if err := db.Exec("DELETE FROM tickets").Error; err != nil {
		return fmt.Errorf("failed to clear tickets: %v", err)
	}
	log.Println("Cleared tickets table")

	if err := db.Exec("DELETE FROM conversation_messages").Error; err != nil {
		return fmt.Errorf("failed to clear conversation_messages: %v", err)
	}
	log.Println("Cleared conversation_messages table")

	if err := db.Exec("DELETE FROM conversations").Error; err != nil {
		return fmt.Errorf("failed to clear conversations: %v", err)
	}
	log.Println("Cleared conversations table")

	if err := db.Exec("DELETE FROM organizations").Error; err != nil {
		return fmt.Errorf("failed to clear organizations: %v", err)
	}
	log.Println("Cleared organizations table")

	if err := db.Exec("DELETE FROM users").Error; err != nil {
		return fmt.Errorf("failed to clear users: %v", err)
	}
	log.Println("Cleared users table")

	tables := []string{"tickets", "conversation_messages", "conversations", "organizations", "users"}
	for _, table := range tables {
		if err := db.Exec(fmt.Sprintf("ALTER TABLE %s AUTO_INCREMENT = 1", table)).Error; err != nil {
			log.Printf("Warning: Could not reset auto-increment for %s: %v", table, err)
		}
	}

	log.Println("All tables cleared successfully!")
	return nil
}

func SeedInitialData() error {
	db := database.DB
	log.Println("Starting database seeding...")

	// Get role IDs
	var superAdminRole, ownerRole, salesRole, guestRole models.UserRoleModel
	if err := db.Where("name = ?", models.RoleSuperAdmin).First(&superAdminRole).Error; err != nil {
		return fmt.Errorf("failed to get super admin role: %v", err)
	}
	if err := db.Where("name = ?", models.RoleOrganizationOwner).First(&ownerRole).Error; err != nil {
		return fmt.Errorf("failed to get owner role: %v", err)
	}
	if err := db.Where("name = ?", models.RoleOrganizationSales).First(&salesRole).Error; err != nil {
		return fmt.Errorf("failed to get sales role: %v", err)
	}
	if err := db.Where("name = ?", models.RoleGuest).First(&guestRole).Error; err != nil {
		return fmt.Errorf("failed to get guest role: %v", err)
	}
	//##############################################################################################
	superAdmin := models.UserModel{
		Name:     "Super Admin",
		Email:    "admin@sociomile.com",
		Password: "password123",
		RoleID:   superAdminRole.ID,
	}
	if err := db.Create(&superAdmin).Error; err != nil {
		return fmt.Errorf("failed to create super admin: %v", err)
	}

	log.Printf("Created Super Admin: %s", superAdmin.Email)
	//##############################################################################################
	// 2. Create Organization Owner
	orgOwner := models.UserModel{
		Name:     "John Doe",
		Email:    "owner@techcorp.com",
		Password: "password123",
		RoleID:   ownerRole.ID,
}

	if err := db.Create(&orgOwner).Error; err != nil {
		return fmt.Errorf("failed to create organization owner: %v", err)
	}
	log.Printf("Created Organization Owner: %s", orgOwner.Email)

	organization := models.OrganizationModel{
		Name:    "TechCorp Solutions",
		OwnerID: orgOwner.ID,
	}
	if err := db.Create(&organization).Error; err != nil {
		return fmt.Errorf("failed to create organization: %v", err)
	}
	log.Printf("Created Organization: %s", organization.Name)

	orgOwner.OrganizationID = &organization.ID
	if err := db.Save(&orgOwner).Error; err != nil {
		return fmt.Errorf("failed to update owner organization: %v", err)
	}

	salesStaff1 := models.UserModel{
		Name:           "Alice Johnson",
		Email:          "alice@techcorp.com",
		Password:       "password123",
		RoleID:         salesRole.ID,
		OrganizationID: &organization.ID,
	}
	if err := db.Create(&salesStaff1).Error; err != nil {
		return fmt.Errorf("failed to create sales staff 1: %v", err)
	}
	log.Printf("Created Sales Staff: %s", salesStaff1.Email)

	salesStaff2 := models.UserModel{
		Name:           "Bob Smith",
		Email:          "bob@techcorp.com",
		Password:       "password123",
		RoleID:         salesRole.ID,
		OrganizationID: &organization.ID,
	}
	if err := db.Create(&salesStaff2).Error; err != nil {
		return fmt.Errorf("failed to create sales staff 2: %v", err)
	}

	log.Printf("Created Sales Staff: %s", salesStaff2.Email)
	//##############################################################################################

	guest1 := models.UserModel{
		Name:     "Customer One",
		Email:    "customer1@example.com",
		Password: "password123",
		RoleID:   guestRole.ID,
	}
	if err := db.Create(&guest1).Error; err != nil {
		return fmt.Errorf("failed to create guest 1: %v", err)
	}

	log.Printf("✓ Created Guest User: %s", guest1.Email)

	guest2 := models.UserModel{
		Name:     "Customer Two",
		Email:    "customer2@example.com",
		Password: "password123",
		RoleID:   guestRole.ID,
	}
	if err := db.Create(&guest2).Error; err != nil {
		return fmt.Errorf("failed to create guest 2: %v", err)
	}

	log.Printf("✓ Created Guest User: %s", guest2.Email)

	guest3 := models.UserModel{
		Name:     "Customer Three",
		Email:    "customer3@example.com",
		Password: "password123",
		RoleID:   guestRole.ID,
	}
	if err := db.Create(&guest3).Error; err != nil {
		return fmt.Errorf("failed to create guest 3: %v", err)
	}
	log.Printf("✓ Created Guest User: %s", guest3.Email)
	//##############################################################################################

	conv1 := models.ConversationModel{
		OrganizationID: organization.ID,
		GuestID:        guest1.ID,
		Status:         models.ConversationStatusPending,
	}
	if err := db.Create(&conv1).Error; err != nil {
		return fmt.Errorf("failed to create conversation 1: %v", err)
	}
	log.Printf("Created Conversation 1 (Pending)")

	conv2 := models.ConversationModel{
		OrganizationID:      organization.ID,
		GuestID:             guest2.ID,
		OrganizationStaffID: &salesStaff1.ID,
		Status:              models.ConversationStatusInProgress,
	}
	if err := db.Create(&conv2).Error; err != nil {
		return fmt.Errorf("failed to create conversation 2: %v", err)
	}
	log.Printf("Created Conversation 2 (In Progress)")

	conv3 := models.ConversationModel{
		OrganizationID:      organization.ID,
		GuestID:             guest3.ID,
		OrganizationStaffID: &salesStaff2.ID,
		Status:              models.ConversationStatusInProgress,
	}
	if err := db.Create(&conv3).Error; err != nil {
		return fmt.Errorf("failed to create conversation 3: %v", err)
	}
	log.Printf("Created Conversation 3 (In Progress)")

	conv4 := models.ConversationModel{
		OrganizationID:      organization.ID,
		GuestID:             guest1.ID,
		OrganizationStaffID: &salesStaff1.ID,
		Status:              models.ConversationStatusDone,
	}
	if err := db.Create(&conv4).Error; err != nil {
		return fmt.Errorf("failed to create conversation 4: %v", err)
	}
	log.Printf("Created Conversation 4 (Done)")

	messages := []models.ConversationMessageModel{
		{
			OrganizationID: organization.ID,
			ConversationID: conv1.ID,
			CreatedByID:    guest1.ID,
			Message:        "Hello, I need help with your product pricing.",
		},
	
	}

	for i, msg := range messages {
		if err := db.Create(&msg).Error; err != nil {
			return fmt.Errorf("failed to create message %d: %v", i+1, err)
		}
	}
	log.Printf("Created %d Conversation Messages", len(messages))

	ticket1 := models.TicketModel{
		OrganizationID: organization.ID,
		ConversationID: conv2.ID,
		CreatedByID:    salesStaff1.ID,
		TicketNumber:   fmt.Sprintf("TKT-%d-20260214-0001", organization.ID),
		Name:           "Account Login Issue",
		Status:         models.TicketStatusInProgress,
	}
	if err := db.Create(&ticket1).Error; err != nil {
		return fmt.Errorf("failed to create ticket 1: %v", err)
	}
	log.Printf("✓ Created Ticket: %s", ticket1.TicketNumber)

	ticket2 := models.TicketModel{
		OrganizationID: organization.ID,
		ConversationID: conv3.ID,
		CreatedByID:    salesStaff2.ID,
		TicketNumber:   fmt.Sprintf("TKT-%d-20260214-0002", organization.ID),
		Name:           "Enterprise Plan Inquiry",
		Status:         models.TicketStatusPending,
	}
	if err := db.Create(&ticket2).Error; err != nil {
		return fmt.Errorf("failed to create ticket 2: %v", err)
	}
	log.Printf("✓ Created Ticket: %s", ticket2.TicketNumber)

	ticket3 := models.TicketModel{
		OrganizationID: organization.ID,
		ConversationID: conv4.ID,
		CreatedByID:    salesStaff1.ID,
		TicketNumber:   fmt.Sprintf("TKT-%d-20260214-0003", organization.ID),
		Name:           "Product Setup Assistance",
		Status:         models.TicketStatusDone,
	}
	if err := db.Create(&ticket3).Error; err != nil {
		return fmt.Errorf("failed to create ticket 3: %v", err)
	}
	return nil
}
