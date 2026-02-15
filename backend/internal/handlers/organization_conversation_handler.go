package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"DewaSRY/sociomile-app/internal/services"
	"DewaSRY/sociomile-app/pkg/dtos/requestdto"
	"DewaSRY/sociomile-app/pkg/dtos/responsedto"
	jwtLib "DewaSRY/sociomile-app/pkg/lib/jwt"
	"DewaSRY/sociomile-app/pkg/lib/logger"
	"DewaSRY/sociomile-app/pkg/utils"

	"github.com/go-chi/chi/v5"
)

type OrganizationCOnversationHandler struct {
	jwtService jwtLib.JwtService
	service    services.OrganizationConversationService
}

func NewOrganizationConversationHandler(
	jwtService jwtLib.JwtService,
	service services.OrganizationConversationService,
) *OrganizationCOnversationHandler {
	return &OrganizationCOnversationHandler{
		jwtService: jwtService,
		service:    service,
	}
}

// GetConversationsList godoc
// @Summary      Get conversations list with pagination
// @Description  Retrieve list of conversations for organization with pagination support
// @Tags         conversations
// @Accept       json
// @Produce      json
// @Param        request  query  filtersdto.FiltersDto  false  "Pagination query"
// @Success      200      {object}  responsedto.ConversationListResponse
// @Failure      500      {object}  responsedto.ErrorResponse
// @Security     BearerAuth
// @Router       /conversations [get]
func (h *OrganizationCOnversationHandler) GetConversationsList(w http.ResponseWriter, r *http.Request) {
	filter := utils.ParsePagination(r)
	user, _ := h.jwtService.GetUserFromContext(r.Context())

	result, err := h.service.GetConversationsList(user, filter)
	if err != nil {
		errorData := responsedto.ErrorResponse{
			Message: "failed to fetch conversations",
			Error:   err.Error(),
			Code:    http.StatusInternalServerError,
		}
		logger.ErrorLog("Failed to fetch conversations", errorData)
		utils.WriteJSONResponse(w, http.StatusInternalServerError, errorData)
		return
	}

	logger.InfoLog("Conversations fetched successfully", map[string]any{
		"count": len(result.Conversations),
	})
	utils.WriteJSONResponse(w, http.StatusOK, result)
}

// GetConversationByID godoc
// @Summary      Get conversation by ID
// @Description  Retrieve a conversation by its ID
// @Tags         conversations
// @Accept       json
// @Produce      json
// @Param        id path int true "Conversation ID"
// @Success      200  {object}  responsedto.ConversationResponse
// @Failure      400  {object}  responsedto.ErrorResponse
// @Failure      404  {object}  responsedto.ErrorResponse
// @Security     BearerAuth
// @Router       /conversations/{id} [get]
func (h *OrganizationCOnversationHandler) GetConversationByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		errorData := responsedto.ErrorResponse{
			Message: "invalid conversation id",
			Error:   err.Error(),
			Code:    http.StatusBadRequest,
		}
		logger.ErrorLog("Invalid conversation ID", errorData)
		utils.WriteJSONResponse(w, http.StatusBadRequest, errorData)
		return
	}

	result, err := h.service.GetConversationByID(uint(id))
	if err != nil {
		errorData := responsedto.ErrorResponse{
			Message: "conversation not found",
			Error:   err.Error(),
			Code:    http.StatusNotFound,
		}
		logger.ErrorLog("Conversation not found", errorData)
		utils.WriteJSONResponse(w, http.StatusNotFound, errorData)
		return
	}

	logger.InfoLog("Conversation fetched successfully", map[string]any{
		"conversation_id": result.ID,
	})
	utils.WriteJSONResponse(w, http.StatusOK, result)
}

// AssignConversation godoc
// @Summary      Assign conversation to staff
// @Description  Assign a conversation to an organization staff member
// @Tags         conversations
// @Accept       json
// @Produce      json
// @Param        id path int true "Conversation ID"
// @Param        request body requestdto.AssignConversationRequest true "Assign Conversation Request"
// @Success      200  {object}  responsedto.ConversationResponse
// @Failure      400  {object}  responsedto.ErrorResponse
// @Failure      404  {object}  responsedto.ErrorResponse
// @Failure      500  {object}  responsedto.ErrorResponse
// @Security     BearerAuth
// @Router       /conversations/{id}/assign [put]
func (h *OrganizationCOnversationHandler) AssignConversation(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		errorData := responsedto.ErrorResponse{
			Message: "invalid conversation id",
			Error:   err.Error(),
			Code:    http.StatusBadRequest,
		}
		logger.ErrorLog("Invalid conversation ID", errorData)
		utils.WriteJSONResponse(w, http.StatusBadRequest, errorData)
		return
	}

	var req requestdto.AssignConversationRequest
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

	result, err := h.service.AssignConversation(uint(id), req)
	if err != nil {
		errorData := responsedto.ErrorResponse{
			Message: "failed to assign conversation",
			Error:   err.Error(),
			Code:    http.StatusInternalServerError,
		}
		logger.ErrorLog("Failed to assign conversation", errorData)
		utils.WriteJSONResponse(w, http.StatusInternalServerError, errorData)
		return
	}

	logger.InfoLog("Conversation assigned successfully", map[string]any{
		"conversation_id":        result.ID,
		"organization_staff_id": result.OrganizationStaffID,
	})
	utils.WriteJSONResponse(w, http.StatusOK, result)
}

// UpdateConversationStatus godoc
// @Summary      Update conversation status
// @Description  Update the status of a conversation
// @Tags         conversations
// @Accept       json
// @Produce      json
// @Param        id path int true "Conversation ID"
// @Param        request body requestdto.UpdateConversationRequest true "Update Conversation Request"
// @Success      200  {object}  responsedto.CommonResponse
// @Failure      400  {object}  responsedto.ErrorResponse
// @Failure      404  {object}  responsedto.ErrorResponse
// @Failure      500  {object}  responsedto.ErrorResponse
// @Security     BearerAuth
// @Router       /conversations/{id}/status [put]
func (h *OrganizationCOnversationHandler) UpdateConversationStatus(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		errorData := responsedto.ErrorResponse{
			Message: "invalid conversation id",
			Error:   err.Error(),
			Code:    http.StatusBadRequest,
		}
		logger.ErrorLog("Invalid conversation ID", errorData)
		utils.WriteJSONResponse(w, http.StatusBadRequest, errorData)
		return
	}

	var req requestdto.UpdateConversationRequest
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

	err = h.service.UpdateConversationStatus(uint(id), req)
	if err != nil {
		errorData := responsedto.ErrorResponse{
			Message: "failed to update conversation status",
			Error:   err.Error(),
			Code:    http.StatusInternalServerError,
		}
		logger.ErrorLog("Failed to update conversation status", errorData)
		utils.WriteJSONResponse(w, http.StatusInternalServerError, errorData)
		return
	}

	result := responsedto.CommonResponse{
		Message: "Conversation status updated successfully",
	}
	logger.InfoLog("Conversation status updated successfully", map[string]any{
		"conversation_id": id,
		"status":          req.Status,
	})
	utils.WriteJSONResponse(w, http.StatusOK, result)
}
