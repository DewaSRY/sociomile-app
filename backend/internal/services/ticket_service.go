package services

import (
	"DewaSRY/sociomile-app/pkg/dtos/requestdto"
	"DewaSRY/sociomile-app/pkg/dtos/responsedto"
)

type TicketService interface {
	CreateTicket(userID uint, req requestdto.CreateTicketRequest) (*responsedto.TicketResponse, error)
	GetTicketByID(id uint) (*responsedto.TicketResponse, error)
	GetTicketByNumber(ticketNumber string) (*responsedto.TicketResponse, error)
	GetTicketsByOrganization(organizationID uint) (*responsedto.TicketListResponse, error)
	GetTicketsByConversation(conversationID uint) (*responsedto.TicketListResponse, error)
	UpdateTicket(ticketID uint, req requestdto.UpdateTicketRequest) (*responsedto.TicketResponse, error)
	DeleteTicket(id uint) error
}
