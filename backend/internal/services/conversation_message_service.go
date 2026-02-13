package services

import (
	"DewaSRY/sociomile-app/pkg/dtos/requestdto"
	"DewaSRY/sociomile-app/pkg/dtos/responsedto"
)


type ConversationMessageService interface{
	CreateMessage(userID uint, req requestdto.CreateConversationMessageRequest) (*responsedto.ConversationMessageResponse, error) 
	 GetMessagesByConversation(conversationID uint) (*responsedto.ConversationMessageListResponse, error) 
}

