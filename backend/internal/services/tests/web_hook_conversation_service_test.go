package tests

import (
	"DewaSRY/sociomile-app/internal/services/impl"
	"DewaSRY/sociomile-app/pkg/dtos/requestdto"
	"testing"
)

func TestWebHookConversationService_ProcessConversation_NoInternalError(t *testing.T) {
	tx := SetupTestDB(t)
	service := impl.NewWebHookConversationService(tx)

	org, _ := CreateTestOrganizationWithOwner(tx, t, "Test Organization")

	req := requestdto.WebHooksRequest{
		OrganizationID: org.ID,
		Email:          "webhook@example.com",
		Message:        "Test webhook message",
	}

	err := service.ProcessConversation(req)

	if err != nil {
		t.Errorf("Expected no internal error, got: %v", err)
	}
}