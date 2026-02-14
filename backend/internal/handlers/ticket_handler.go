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

	serviceImpl "DewaSRY/sociomile-app/internal/services/impl"

	"github.com/go-chi/chi/v5"
)

type TicketHandler struct {
	service services.TicketService
}

func NewTicketHandler() *TicketHandler {
	return &TicketHandler{
		service: serviceImpl.InstanceTicketService(),
	}
}

// CreateTicket godoc
// @Summary      Create a new ticket
// @Description  Create a new ticket from a conversation
// @Tags         tickets
// @Accept       json
// @Produce      json
// @Param        request body requestdto.CreateTicketRequest true "Create Ticket Request"
// @Success      201  {object}  responsedto.TicketResponse
// @Failure      400  {object}  responsedto.ErrorResponse
// @Failure      500  {object}  responsedto.ErrorResponse
// @Security     BearerAuth
// @Router       /tickets [post]
func (h *TicketHandler) CreateTicket(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(uint)

	var req requestdto.CreateTicketRequest
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

	result, err := h.service.CreateTicket(userID, req)
	if err != nil {
		errorData := responsedto.ErrorResponse{
			Message: "failed to create ticket",
			Error:   err.Error(),
			Code:    http.StatusInternalServerError,
		}
		logger.ErrorLog("Failed to create ticket", errorData)
		utils.WriteJSONResponse(w, http.StatusInternalServerError, errorData)
		return
	}

	logger.InfoLog("Ticket created successfully", map[string]any{
		"ticket_id":     result.ID,
		"ticket_number": result.TicketNumber,
	})
	utils.WriteJSONResponse(w, http.StatusCreated, result)
}

// GetTicket godoc
// @Summary      Get ticket by ID
// @Description  Retrieve a ticket by its ID
// @Tags         tickets
// @Accept       json
// @Produce      json
// @Param        id path int true "Ticket ID"
// @Success      200  {object}  responsedto.TicketResponse
// @Failure      400  {object}  responsedto.ErrorResponse
// @Failure      404  {object}  responsedto.ErrorResponse
// @Security     BearerAuth
// @Router       /tickets/{id} [get]
func (h *TicketHandler) GetTicket(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		errorData := responsedto.ErrorResponse{
			Message: "invalid ticket id",
			Error:   err.Error(),
			Code:    http.StatusBadRequest,
		}
		logger.ErrorLog("Invalid ticket ID", errorData)
		utils.WriteJSONResponse(w, http.StatusBadRequest, errorData)
		return
	}

	result, err := h.service.GetTicketByID(uint(id))
	if err != nil {
		errorData := responsedto.ErrorResponse{
			Message: "ticket not found",
			Error:   err.Error(),
			Code:    http.StatusNotFound,
		}
		logger.ErrorLog("Ticket not found", errorData)
		utils.WriteJSONResponse(w, http.StatusNotFound, errorData)
		return
	}

	utils.WriteJSONResponse(w, http.StatusOK, result)
}

// GetTicketByNumber godoc
// @Summary      Get ticket by number
// @Description  Retrieve a ticket by its ticket number
// @Tags         tickets
// @Accept       json
// @Produce      json
// @Param        number path string true "Ticket Number"
// @Success      200  {object}  responsedto.TicketResponse
// @Failure      404  {object}  responsedto.ErrorResponse
// @Security     BearerAuth
// @Router       /tickets/number/{number} [get]
func (h *TicketHandler) GetTicketByNumber(w http.ResponseWriter, r *http.Request) {
	ticketNumber := chi.URLParam(r, "number")

	result, err := h.service.GetTicketByNumber(ticketNumber)
	if err != nil {
		errorData := responsedto.ErrorResponse{
			Message: "ticket not found",
			Error:   err.Error(),
			Code:    http.StatusNotFound,
		}
		logger.ErrorLog("Ticket not found", errorData)
		utils.WriteJSONResponse(w, http.StatusNotFound, errorData)
		return
	}

	utils.WriteJSONResponse(w, http.StatusOK, result)
}

// GetOrganizationTickets godoc
// @Summary      Get organization tickets
// @Description  Retrieve all tickets for an organization (Owner/Staff only)
// @Tags         tickets
// @Accept       json
// @Produce      json
// @Param        organization_id path int true "Organization ID"
// @Success      200  {object}  responsedto.TicketListResponse
// @Failure      400  {object}  responsedto.ErrorResponse
// @Failure      500  {object}  responsedto.ErrorResponse
// @Security     BearerAuth
// @Router       /organizations/{organization_id}/tickets [get]
func (h *TicketHandler) GetOrganizationTickets(w http.ResponseWriter, r *http.Request) {
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

	result, err := h.service.GetTicketsByOrganization(uint(organizationID))
	if err != nil {
		errorData := responsedto.ErrorResponse{
			Message: "failed to fetch tickets",
			Error:   err.Error(),
			Code:    http.StatusInternalServerError,
		}
		logger.ErrorLog("Failed to fetch tickets", errorData)
		utils.WriteJSONResponse(w, http.StatusInternalServerError, errorData)
		return
	}

	utils.WriteJSONResponse(w, http.StatusOK, result)
}

// GetConversationTickets godoc
// @Summary      Get conversation tickets
// @Description  Retrieve all tickets for a conversation
// @Tags         tickets
// @Accept       json
// @Produce      json
// @Param        conversation_id path int true "Conversation ID"
// @Success      200  {object}  responsedto.TicketListResponse
// @Failure      400  {object}  responsedto.ErrorResponse
// @Failure      500  {object}  responsedto.ErrorResponse
// @Security     BearerAuth
// @Router       /conversations/{conversation_id}/tickets [get]
func (h *TicketHandler) GetConversationTickets(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "conversation_id")
	conversationID, err := strconv.ParseUint(idStr, 10, 32)
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

	result, err := h.service.GetTicketsByConversation(uint(conversationID))
	if err != nil {
		errorData := responsedto.ErrorResponse{
			Message: "failed to fetch tickets",
			Error:   err.Error(),
			Code:    http.StatusInternalServerError,
		}
		logger.ErrorLog("Failed to fetch tickets", errorData)
		utils.WriteJSONResponse(w, http.StatusInternalServerError, errorData)
		return
	}

	utils.WriteJSONResponse(w, http.StatusOK, result)
}

// UpdateTicket godoc
// @Summary      Update ticket
// @Description  Update a ticket's details
// @Tags         tickets
// @Accept       json
// @Produce      json
// @Param        id path int true "Ticket ID"
// @Param        request body requestdto.UpdateTicketRequest true "Update Ticket Request"
// @Success      200  {object}  responsedto.TicketResponse
// @Failure      400  {object}  responsedto.ErrorResponse
// @Failure      404  {object}  responsedto.ErrorResponse
// @Security     BearerAuth
// @Router       /tickets/{id} [put]
func (h *TicketHandler) UpdateTicket(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		errorData := responsedto.ErrorResponse{
			Message: "invalid ticket id",
			Error:   err.Error(),
			Code:    http.StatusBadRequest,
		}
		logger.ErrorLog("Invalid ticket ID", errorData)
		utils.WriteJSONResponse(w, http.StatusBadRequest, errorData)
		return
	}

	var req requestdto.UpdateTicketRequest
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

	result, err := h.service.UpdateTicket(uint(id), req)
	if err != nil {
		errorData := responsedto.ErrorResponse{
			Message: "failed to update ticket",
			Error:   err.Error(),
			Code:    http.StatusInternalServerError,
		}
		logger.ErrorLog("Failed to update ticket", errorData)
		utils.WriteJSONResponse(w, http.StatusInternalServerError, errorData)
		return
	}

	logger.InfoLog("Ticket updated successfully", map[string]any{
		"ticket_id": result.ID,
	})
	utils.WriteJSONResponse(w, http.StatusOK, result)
}

// DeleteTicket godoc
// @Summary      Delete ticket
// @Description  Delete a ticket
// @Tags         tickets
// @Accept       json
// @Produce      json
// @Param        id path int true "Ticket ID"
// @Success      200  {object}  responsedto.SuccessResponse
// @Failure      400  {object}  responsedto.ErrorResponse
// @Failure      404  {object}  responsedto.ErrorResponse
// @Security     BearerAuth
// @Router       /tickets/{id} [delete]
func (h *TicketHandler) DeleteTicket(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		errorData := responsedto.ErrorResponse{
			Message: "invalid ticket id",
			Error:   err.Error(),
			Code:    http.StatusBadRequest,
		}
		logger.ErrorLog("Invalid ticket ID", errorData)
		utils.WriteJSONResponse(w, http.StatusBadRequest, errorData)
		return
	}

	if err := h.service.DeleteTicket(uint(id)); err != nil {
		errorData := responsedto.ErrorResponse{
			Message: "failed to delete ticket",
			Error:   err.Error(),
			Code:    http.StatusInternalServerError,
		}
		logger.ErrorLog("Failed to delete ticket", errorData)
		utils.WriteJSONResponse(w, http.StatusInternalServerError, errorData)
		return
	}

	logger.InfoLog("Ticket deleted successfully", map[string]any{
		"ticket_id": id,
	})
	utils.WriteJSONResponse(w, http.StatusOK, responsedto.SuccessResponse{
		Message: "ticket deleted successfully",
	})
}
