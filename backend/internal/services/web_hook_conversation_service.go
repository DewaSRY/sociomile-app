package services

import "DewaSRY/sociomile-app/pkg/dtos/requestdto"


type WebHookConversationService interface{
	ProcessConversation(requestdto.WebHooksRequest) (error)
}