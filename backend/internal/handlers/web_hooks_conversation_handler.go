package handlers

import (
	"DewaSRY/sociomile-app/internal/services"
	"DewaSRY/sociomile-app/internal/services/impl"
	"DewaSRY/sociomile-app/pkg/dtos/requestdto"
	"DewaSRY/sociomile-app/pkg/dtos/responsedto"
	"DewaSRY/sociomile-app/pkg/lib/logger"
	"DewaSRY/sociomile-app/pkg/utils"
	"encoding/json"
	"errors"
	"net/http"
)

type WebHookHandler struct {
	service services.WebHookConversationService
}

func NewWebHookHandler(
	service services.WebHookConversationService,
) *WebHookHandler {
	return &WebHookHandler{
		service: service,
	}
}

// CreateOrganization godoc
// @Summary      Create message conversation
// @Description  Create message conversation
// @Tags         webhooks-conversation
// @Accept       json
// @Produce      json
// @Param        request body requestdto.WebHooksRequest true "Create message conversation"
// @Success      201  {object}  responsedto.CommonResponse
// @Failure      400  {object}  responsedto.ErrorResponse
// @Failure      500  {object}  responsedto.ErrorResponse
// @Router       /webhooks/conversations [post]
func (h *WebHookHandler) CreateConversation(w http.ResponseWriter, r *http.Request) {
	var req requestdto.WebHooksRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		errorData := responsedto.ErrorResponse{
			Message: "invalid request",
			Error:   err.Error(),
			Code:    http.StatusBadRequest,
		}
		logger.ErrorLog("Failed to decode request", errorData)
		utils.WriteJSONResponse(w, http.StatusBadRequest, errorData)
		return
	}

	if err := utils.ValidateStruct(req); err != nil {
		errorData := responsedto.ErrorResponse{
			Message: "invalid request",
			Error:   err.Error(),
			Code:    http.StatusBadRequest,
		}
		logger.ErrorLog("Failed to validate request", errorData)
		utils.WriteJSONResponse(w, http.StatusBadRequest, errorData)
		return
	}

	err := h.service.ProcessConversation(req)
	if err != nil {

		if errors.Is(err, impl.ErrOrganizationNotFound) {
			errorData := responsedto.ErrorResponse{
				Message: "Organization Is not found",
				Error:   err.Error(),
				Code:    http.StatusNotFound,
			}
			logger.ErrorLog("Organization Is not found", errorData)
			utils.WriteJSONResponse(w, http.StatusNotFound, errorData)
			return
		}
		errorData := responsedto.ErrorResponse{
			Message: "failed to create conversation",
			Error:   err.Error(),
			Code:    http.StatusInternalServerError,
		}
		logger.ErrorLog("Failed to create conversation", errorData)
		utils.WriteJSONResponse(w, http.StatusInternalServerError, errorData)
		return
	}

	result := responsedto.CommonResponse{
		Message: "Success Create message conversation ",
		Code:    http.StatusCreated,
	}

	logger.InfoLog("Organization created successfully", result)
	utils.WriteJSONResponse(w, http.StatusCreated, result)
}
