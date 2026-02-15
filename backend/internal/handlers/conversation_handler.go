package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"DewaSRY/sociomile-app/internal/services"
	"DewaSRY/sociomile-app/pkg/dtos/requestdto"
	"DewaSRY/sociomile-app/pkg/dtos/responsedto"
	"DewaSRY/sociomile-app/pkg/lib/logger"
	"DewaSRY/sociomile-app/pkg/utils"

	"github.com/go-chi/chi/v5"
)

type ConversationHandler struct {
	service        services.OrganizationConversationService
	messageService services.ConversationMessageService
}

func NewConversationHandler(
	service        services.OrganizationConversationService,
	messageService services.ConversationMessageService,
) *ConversationHandler {
	return &ConversationHandler{
		service:        service,
		messageService: messageService,
	}
}

// CreateConversation godoc
// @Summary      Create a new conversation
// @Description  Create a new conversation with an organization (Guest user)
// @Tags         conversations
// @Accept       json
// @Produce      json
// @Param        request body requestdto.CreateConversationRequest true "Create Conversation Request"
// @Success      201  {object}  responsedto.ConversationResponse
// @Failure      400  {object}  responsedto.ErrorResponse
// @Failure      500  {object}  responsedto.ErrorResponse
// @Security     BearerAuth
// @Router       /conversations [post]
func (h *ConversationHandler) CreateConversation(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(uint)

	var req requestdto.CreateConversationRequest
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

	result, err := h.service.CreateConversation(userID, req)
	if err != nil {
		errorData := responsedto.ErrorResponse{
			Message: "failed to create conversation",
			Error:   err.Error(),
			Code:    http.StatusInternalServerError,
		}
		logger.ErrorLog("Failed to create conversation", errorData)
		utils.WriteJSONResponse(w, http.StatusInternalServerError, errorData)
		return
	}

	logger.InfoLog("Conversation created successfully", map[string]any{
		"conversation_id": result.ID,
	})
	utils.WriteJSONResponse(w, http.StatusCreated, result)
}

// GetConversation godoc
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
func (h *ConversationHandler) GetConversation(w http.ResponseWriter, r *http.Request) {
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

	utils.WriteJSONResponse(w, http.StatusOK, result)
}

// GetMyConversations godoc
// @Summary      Get my conversations
// @Description  Retrieve all conversations for the authenticated user (guest or staff)
// @Tags         conversations
// @Accept       json
// @Produce      json
// @Success      200  {object}  responsedto.ConversationListResponse
// @Failure      500  {object}  responsedto.ErrorResponse
// @Security     BearerAuth
// @Router       /conversations/my [get]
func (h *ConversationHandler) GetMyConversations(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(uint)
	role := r.Context().Value("role").(string)

	var result *responsedto.ConversationListResponse
	var err error

	// If guest, get conversations as guest
	// If staff, get assigned conversations
	if role == "guest" {
		result, err = h.service.GetConversationsByGuest(userID)
	} else {
		result, err = h.service.GetConversationsByStaff(userID)
	}

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

	utils.WriteJSONResponse(w, http.StatusOK, result)
}

// GetOrganizationConversations godoc
// @Summary      Get organization conversations
// @Description  Retrieve all conversations for an organization (Owner/Staff only)
// @Tags         conversations
// @Accept       json
// @Produce      json
// @Param        organization_id path int true "Organization ID"
// @Success      200  {object}  responsedto.ConversationListResponse
// @Failure      400  {object}  responsedto.ErrorResponse
// @Failure      500  {object}  responsedto.ErrorResponse
// @Security     BearerAuth
// @Router       /organizations/{organization_id}/conversations [get]
func (h *ConversationHandler) GetOrganizationConversations(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "organization_id")
	organizationID, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		errorData := responsedto.ErrorResponse{
			Message: "invalid organization id",
			Error:   err.Error(),
			Code:    http.StatusBadRequest,
		}
		logger.ErrorLog("Invalid organization ID", errorData)
		utils.WriteJSONResponse(w, http.StatusBadRequest, errorData)
		return
	}

	result, err := h.service.GetConversationsByOrganization(uint(organizationID))
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

	utils.WriteJSONResponse(w, http.StatusOK, result)
}

// AssignConversation godoc
// @Summary      Assign conversation to staff
// @Description  Assign a conversation to a staff member (Organization Owner only)
// @Tags         conversations
// @Accept       json
// @Produce      json
// @Param        id path int true "Conversation ID"
// @Param        request body requestdto.AssignConversationRequest true "Assign Conversation Request"
// @Success      200  {object}  responsedto.ConversationResponse
// @Failure      400  {object}  responsedto.ErrorResponse
// @Failure      500  {object}  responsedto.ErrorResponse
// @Security     BearerAuth
// @Router       /conversations/{id}/assign [post]
func (h *ConversationHandler) AssignConversation(w http.ResponseWriter, r *http.Request) {
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
		"conversation_id": result.ID,
		"staff_id":        req.OrganizationStaffID,
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
// @Success      200  {object}  responsedto.ConversationResponse
// @Failure      400  {object}  responsedto.ErrorResponse
// @Failure      500  {object}  responsedto.ErrorResponse
// @Security     BearerAuth
// @Router       /conversations/{id}/status [put]
func (h *ConversationHandler) UpdateConversationStatus(w http.ResponseWriter, r *http.Request) {
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

	result, err := h.service.UpdateConversationStatus(uint(id), req)
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

	logger.InfoLog("Conversation status updated successfully", map[string]any{
		"conversation_id": result.ID,
		"status":          result.Status,
	})
	utils.WriteJSONResponse(w, http.StatusOK, result)
}

// CreateMessage godoc
// @Summary      Create a message in a conversation
// @Description  Create a new message in a conversation
// @Tags         conversations
// @Accept       json
// @Produce      json
// @Param        request body requestdto.CreateConversationMessageRequest true "Create Message Request"
// @Success      201  {object}  responsedto.ConversationMessageResponse
// @Failure      400  {object}  responsedto.ErrorResponse
// @Failure      500  {object}  responsedto.ErrorResponse
// @Security     BearerAuth
// @Router       /conversations/messages [post]
func (h *ConversationHandler) CreateMessage(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(uint)

	var req requestdto.CreateConversationMessageRequest
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

	result, err := h.messageService.CreateMessage(userID, req)
	if err != nil {
		errorData := responsedto.ErrorResponse{
			Message: "failed to create message",
			Error:   err.Error(),
			Code:    http.StatusInternalServerError,
		}
		logger.ErrorLog("Failed to create message", errorData)
		utils.WriteJSONResponse(w, http.StatusInternalServerError, errorData)
		return
	}

	logger.InfoLog("Message created successfully", map[string]any{
		"message_id":      result.ID,
		"conversation_id": result.ConversationID,
	})
	utils.WriteJSONResponse(w, http.StatusCreated, result)
}

// GetConversationMessages godoc
// @Summary      Get conversation messages
// @Description  Retrieve all messages for a conversation
// @Tags         conversations
// @Accept       json
// @Produce      json
// @Param        id path int true "Conversation ID"
// @Success      200  {object}  responsedto.ConversationMessageListResponse
// @Failure      400  {object}  responsedto.ErrorResponse
// @Failure      500  {object}  responsedto.ErrorResponse
// @Security     BearerAuth
// @Router       /conversations/{id}/messages [get]
func (h *ConversationHandler) GetConversationMessages(w http.ResponseWriter, r *http.Request) {
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

	result, err := h.messageService.GetMessagesByConversation(uint(id))
	if err != nil {
		errorData := responsedto.ErrorResponse{
			Message: "failed to fetch messages",
			Error:   err.Error(),
			Code:    http.StatusInternalServerError,
		}
		logger.ErrorLog("Failed to fetch messages", errorData)
		utils.WriteJSONResponse(w, http.StatusInternalServerError, errorData)
		return
	}

	utils.WriteJSONResponse(w, http.StatusOK, result)
}
