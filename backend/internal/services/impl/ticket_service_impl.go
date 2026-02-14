package impl

import (
	"DewaSRY/sociomile-app/internal/database"
	"DewaSRY/sociomile-app/internal/services"
	"DewaSRY/sociomile-app/pkg/dtos/requestdto"
	"DewaSRY/sociomile-app/pkg/dtos/responsedto"
	"DewaSRY/sociomile-app/pkg/models"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type ticketServiceImpl struct {
}

// CreateTicket implements services.TicketService.
func (t *ticketServiceImpl) CreateTicket(userID uint, req requestdto.CreateTicketRequest) (*responsedto.TicketResponse, error) {
	var conversation models.ConversationModel
	if err := database.DB.First(&conversation, req.ConversationID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("conversation not found")
		}
		return nil, errors.New("failed to fetch conversation")
	}

	ticketNumber := t.generateTicketNumber(conversation.OrganizationID)

	ticket := models.TicketModel{
		OrganizationID: conversation.OrganizationID,
		ConversationID: req.ConversationID,
		CreatedByID:    userID,
		TicketNumber:   ticketNumber,
		Name:           req.Name,
		Status:         models.TicketStatusPending,
	}

	if err := database.DB.Create(&ticket).Error; err != nil {
		return nil, errors.New("failed to create ticket")
	}

	if err := database.DB.Preload("Organization").Preload("Conversation").Preload("CreatedBy").First(&ticket, ticket.ID).Error; err != nil {
		return nil, errors.New("failed to load ticket details")
	}

	return t.mapToTicketResponse(&ticket), nil
}

// DeleteTicket implements services.TicketService.
func (t *ticketServiceImpl) DeleteTicket(id uint) error {
		var ticket models.TicketModel
	if err := database.DB.First(&ticket, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("ticket not found")
		}
		return errors.New("failed to fetch ticket")
	}

	if err := database.DB.Delete(&ticket).Error; err != nil {
		return errors.New("failed to delete ticket")
	}

	return nil
}

// GetTicketByID implements services.TicketService.
func (t *ticketServiceImpl) GetTicketByID(id uint) (*responsedto.TicketResponse, error) {
	var ticket models.TicketModel
	if err := database.DB.Preload("Organization").Preload("Conversation").Preload("CreatedBy").First(&ticket, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("ticket not found")
		}
		return nil, errors.New("failed to fetch ticket")
	}

	return t.mapToTicketResponse(&ticket), nil
}

// GetTicketByNumber implements services.TicketService.
func (t *ticketServiceImpl) GetTicketByNumber(ticketNumber string) (*responsedto.TicketResponse, error) {
		var ticket models.TicketModel
	if err := database.DB.Where("ticket_number = ?", ticketNumber).
		Preload("Organization").Preload("Conversation").Preload("CreatedBy").
		First(&ticket).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("ticket not found")
		}
		return nil, errors.New("failed to fetch ticket")
	}

	return t.mapToTicketResponse(&ticket), nil
}

// GetTicketsByConversation implements services.TicketService.
func (t *ticketServiceImpl) GetTicketsByConversation(conversationID uint) (*responsedto.TicketListResponse, error) {
		var tickets []models.TicketModel
	if err := database.DB.Where("conversation_id = ?", conversationID).
		Preload("Organization").Preload("CreatedBy").
		Order("created_at DESC").
		Find(&tickets).Error; err != nil {
		return nil, errors.New("failed to fetch tickets")
	}

	return t.buildTicketListResponse(tickets), nil
}

// GetTicketsByOrganization implements services.TicketService.
func (t *ticketServiceImpl) GetTicketsByOrganization(organizationID uint) (*responsedto.TicketListResponse, error) {
		var tickets []models.TicketModel
	if err := database.DB.Where("organization_id = ?", organizationID).
		Preload("Conversation").Preload("CreatedBy").
		Order("created_at DESC").
		Find(&tickets).Error; err != nil {
		return nil, errors.New("failed to fetch tickets")
	}

	return t.buildTicketListResponse(tickets), nil
}

// UpdateTicket implements services.TicketService.
func (t *ticketServiceImpl) UpdateTicket(ticketID uint, req requestdto.UpdateTicketRequest) (*responsedto.TicketResponse, error) {
	var ticket models.TicketModel
	if err := database.DB.First(&ticket, ticketID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("ticket not found")
		}
		return nil, errors.New("failed to fetch ticket")
	}

	if req.Name != "" {
		ticket.Name = req.Name
	}
	if req.Status != "" {
		ticket.Status = req.Status
	}

	if err := database.DB.Save(&ticket).Error; err != nil {
		return nil, errors.New("failed to update ticket")
	}

	if err := database.DB.Preload("Organization").Preload("Conversation").Preload("CreatedBy").First(&ticket, ticket.ID).Error; err != nil {
		return nil, errors.New("failed to load ticket details")
	}

	return t.mapToTicketResponse(&ticket), nil
}


func (t *ticketServiceImpl) generateTicketNumber(organizationID uint) string {
	var count int64
	database.DB.Model(&models.TicketModel{}).Where("organization_id = ?", organizationID).Count(&count)

	timestamp := time.Now().Format("20060102")
	return fmt.Sprintf("TKT-%d-%s-%04d", organizationID, timestamp, count+1)
}

func (t *ticketServiceImpl) mapToTicketResponse(ticket *models.TicketModel) *responsedto.TicketResponse {
	response := &responsedto.TicketResponse{
		ID:             ticket.ID,
		OrganizationID: ticket.OrganizationID,
		ConversationID: ticket.ConversationID,
		CreatedByID:    ticket.CreatedByID,
		TicketNumber:   ticket.TicketNumber,
		Name:           ticket.Name,
		Status:         ticket.Status,
		CreatedAt:      ticket.CreatedAt,
		UpdatedAt:      ticket.UpdatedAt,
	}

	if ticket.Organization != nil {
		response.Organization = &responsedto.OrganizationResponse{
			ID:   ticket.Organization.ID,
			Name: ticket.Organization.Name,
		}
	}

	if ticket.Conversation != nil {
		response.Conversation = &responsedto.ConversationResponse{
			ID:             ticket.Conversation.ID,
			OrganizationID: ticket.Conversation.OrganizationID,
			GuestID:        ticket.Conversation.GuestID,
			Status:         ticket.Conversation.Status,
		}
	}

	if ticket.CreatedBy != nil {
		response.CreatedBy = &responsedto.UserData{
			ID:    ticket.CreatedBy.ID,
			Email: ticket.CreatedBy.Email,
			Name:  ticket.CreatedBy.Name,
		}
	}

	return response
}

func (t *ticketServiceImpl) buildTicketListResponse(tickets []models.TicketModel) *responsedto.TicketListResponse {
	var ticketResponses []responsedto.TicketResponse
	for _, ticket := range tickets {
		ticketResponses = append(ticketResponses, *t.mapToTicketResponse(&ticket))
	}

	return &responsedto.TicketListResponse{
		Tickets: ticketResponses,
		Metadata: responsedto.PaginateMetaData{
			Total: len(ticketResponses),
			Page:  1,
			Limit: len(ticketResponses),
		},
	}
}


func NewTicketService() services.TicketService {
	return &ticketServiceImpl{}
}
