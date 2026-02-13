package services

import (
	"DewaSRY/sociomile-app/pkg/dtos/requestdto"
	"DewaSRY/sociomile-app/pkg/dtos/responsedto"
)

type ConversationService interface{
	CreateConversation(guestID uint, req requestdto.CreateConversationRequest) (*responsedto.ConversationResponse, error)
	GetConversationByID(id uint) (*responsedto.ConversationResponse, error)
	GetConversationsByOrganization(organizationID uint) (*responsedto.ConversationListResponse, error)
	GetConversationsByGuest(guestID uint) (*responsedto.ConversationListResponse, error) 
	GetConversationsByStaff(staffID uint) (*responsedto.ConversationListResponse, error) 
	AssignConversation(conversationID uint, req requestdto.AssignConversationRequest) (*responsedto.ConversationResponse, error)
	UpdateConversationStatus(conversationID uint, req requestdto.UpdateConversationRequest) (*responsedto.ConversationResponse, error)
}

