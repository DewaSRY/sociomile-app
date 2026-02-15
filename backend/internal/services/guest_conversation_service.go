package services

import (
	"DewaSRY/sociomile-app/pkg/dtos/filtersdto"
	"DewaSRY/sociomile-app/pkg/dtos/requestdto"
	"DewaSRY/sociomile-app/pkg/dtos/responsedto"
	jwtUtil "DewaSRY/sociomile-app/pkg/lib/jwt"
)



type GuestConversationService interface{
	CreateConversation(user *jwtUtil.Claims, req requestdto.CreateConversationRequest) ( error)
	GetConversation(user *jwtUtil.Claims, filter filtersdto.FiltersDto)(*responsedto.ConversationListPaginateResponse, error)
}