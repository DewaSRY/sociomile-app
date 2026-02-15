package services

import (
	"DewaSRY/sociomile-app/pkg/dtos/filtersdto"
	"DewaSRY/sociomile-app/pkg/dtos/requestdto"
	"DewaSRY/sociomile-app/pkg/dtos/responsedto"
	"DewaSRY/sociomile-app/pkg/lib/jwt"
)

type OrganizationConversationService interface{
	GetConversationsList(user *jwt.Claims, filter filtersdto.FiltersDto) (*responsedto.ConversationListResponse, error)

	GetConversationByID(id uint) (*responsedto.ConversationResponse, error)
	AssignConversation(conversationID uint, req requestdto.AssignConversationRequest) (*responsedto.ConversationResponse, error)
	UpdateConversationStatus(conversationID uint, req requestdto.UpdateConversationRequest) ( error)
}

