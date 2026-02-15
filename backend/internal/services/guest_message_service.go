package services

import (
	"DewaSRY/sociomile-app/pkg/dtos/filtersdto"
	"DewaSRY/sociomile-app/pkg/dtos/requestdto"
	"DewaSRY/sociomile-app/pkg/dtos/responsedto"
	"DewaSRY/sociomile-app/pkg/lib/jwt"
)

type GuestMessageService interface {
	SendConversationMessage(user *jwt.Claims, req requestdto.CreateConversationMessageRequest) error
	GetConversationMessageList(user *jwt.Claims, filter filtersdto.FiltersDto, conversationId uint) (*responsedto.ConversationMessagePaginateResponse, error)
}
