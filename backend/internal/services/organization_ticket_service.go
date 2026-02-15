package services

import (
	"DewaSRY/sociomile-app/pkg/dtos/filtersdto"
	"DewaSRY/sociomile-app/pkg/dtos/requestdto"
	"DewaSRY/sociomile-app/pkg/dtos/responsedto"
	jwtLib "DewaSRY/sociomile-app/pkg/lib/jwt"
)

type OrganizationTicketService interface {
	CreateTicket(user *jwtLib.Claims, req requestdto.CreateTicketRequest)error
	GetTicketsList(user *jwtLib.Claims,filter filtersdto.FiltersDto) (*responsedto.TicketListResponse, error)
	UpdateTicket(user *jwtLib.Claims,ticketID uint, req requestdto.UpdateTicketRequest)error
}
