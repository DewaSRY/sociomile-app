package impl

import (
	"DewaSRY/sociomile-app/internal/services"
	"DewaSRY/sociomile-app/pkg/dtos/filtersdto"
	"DewaSRY/sociomile-app/pkg/dtos/requestdto"
	"DewaSRY/sociomile-app/pkg/dtos/responsedto"
	jwtLib "DewaSRY/sociomile-app/pkg/lib/jwt"
	"DewaSRY/sociomile-app/pkg/models"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type OrganizationTicketServiceImpl struct {
	db *gorm.DB
}

// CreateTicket implements services.TicketService.
func (t *OrganizationTicketServiceImpl) CreateTicket(user *jwtLib.Claims, req requestdto.CreateTicketRequest) error {
	var conversation models.ConversationModel
	if err := t.db.First(&conversation, req.ConversationID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("conversation not found")
		}
		return errors.New("failed to fetch conversation")
	}

	ticketNumber := t.generateTicketNumber(conversation.OrganizationID)

	ticket := models.TicketModel{
		OrganizationID: conversation.OrganizationID,
		ConversationID: req.ConversationID,
		CreatedByID:    user.UserID,
		TicketNumber:   ticketNumber,
		Name:           req.Name,
		Status:         models.TicketStatusPending,
	}

	if err := t.db.Create(&ticket).Error; err != nil {
		return errors.New("failed to create ticket")
	}
	return nil
}

// GetTicketsList implements services.TicketService.
func (t *OrganizationTicketServiceImpl) GetTicketsList(user *jwtLib.Claims, filter filtersdto.FiltersDto) (*responsedto.TicketListResponse, error) {
	var tickets []models.TicketModel
	if err := t.db.Where("organization_id = ?", user.OrganizationId).
		Preload("Conversation").Preload("CreatedBy").
		Order("created_at DESC").
		Find(&tickets).Error; err != nil {
		return nil, errors.New("failed to fetch tickets")
	}

	return t.buildTicketListResponse(tickets), nil
}

// UpdateTicket implements services.TicketService.
func (t *OrganizationTicketServiceImpl) UpdateTicket(user *jwtLib.Claims, ticketID uint, req requestdto.UpdateTicketRequest) error {
	var ticket models.TicketModel
	if err := t.db.First(&ticket, ticketID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("ticket not found")
		}
		return errors.New("failed to fetch ticket")
	}

	if req.Name != "" {
		ticket.Name = req.Name
	}
	if req.Status != "" {
		ticket.Status = req.Status
	}

	if err := t.db.Save(&ticket).Error; err != nil {
		return errors.New("failed to update ticket")
	}

	if err := t.db.Preload("Organization").Preload("Conversation").Preload("CreatedBy").First(&ticket, ticket.ID).Error; err != nil {
		return errors.New("failed to load ticket details")
	}

	return nil
}

func (t *OrganizationTicketServiceImpl) generateTicketNumber(organizationID uint) string {
	var count int64
	t.db.Model(&models.TicketModel{}).Where("organization_id = ?", organizationID).Count(&count)

	timestamp := time.Now().Format("20060102")
	return fmt.Sprintf("TKT-%d-%s-%04d", organizationID, timestamp, count+1)
}

func (t *OrganizationTicketServiceImpl) mapToTicketResponse(ticket *models.TicketModel) *responsedto.TicketResponse {
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

func (t *OrganizationTicketServiceImpl) buildTicketListResponse(tickets []models.TicketModel) *responsedto.TicketListResponse {
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

func NewTicketService(db *gorm.DB) services.OrganizationTicketService {
	return &OrganizationTicketServiceImpl{db:db}
}
