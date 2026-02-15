package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"DewaSRY/sociomile-app/internal/services"
	_ "DewaSRY/sociomile-app/pkg/dtos/filtersdto"
	"DewaSRY/sociomile-app/pkg/dtos/requestdto"
	"DewaSRY/sociomile-app/pkg/dtos/responsedto"
	jwtLib "DewaSRY/sociomile-app/pkg/lib/jwt"
	"DewaSRY/sociomile-app/pkg/lib/logger"
	"DewaSRY/sociomile-app/pkg/utils"

	"github.com/go-chi/chi/v5"
)

type OrganizationTicketHandler struct {
	jwtService jwtLib.JwtService
	service    services.OrganizationTicketService
}

func NewOrganizationTicketHandler(
	service services.OrganizationTicketService,
) *OrganizationTicketHandler {
	return &OrganizationTicketHandler{
		service: service,
	}
}

// CreateTicket godoc
// @Summary      Create a new ticket
// @Description  Create a new ticket from a conversation
// @Tags         organization-tickets
// @Accept       json
// @Produce      json
// @Param        request body requestdto.CreateTicketRequest true "Create Ticket Request"
// @Success      201  {object}  responsedto.TicketResponse
// @Failure      400  {object}  responsedto.ErrorResponse
// @Failure      500  {object}  responsedto.ErrorResponse
// @Security     BearerAuth
// @Router       /organizations/ticket [post]
func (t *OrganizationTicketHandler) CreateTicket(w http.ResponseWriter, r *http.Request) {

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
	user, _ := t.jwtService.GetUserFromContext(r.Context())

	err := t.service.CreateTicket(user, req)
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

	result := responsedto.CommonResponse{
		Message: "Ticket created successfully",
	}
	logger.InfoLog("Ticket created successfully", result)
	utils.WriteJSONResponse(w, http.StatusCreated, result)
}

// GetTicketsList godoc
// @Summary      Get tickets list with pagination
// @Description  Retrieve list of tickets for organization with pagination support
// @Tags         organization-tickets
// @Accept       json
// @Produce      json
// @Param        request  query  filtersdto.FiltersDto  false  "Pagination query"
// @Success      200      {object}  responsedto.TicketListResponse
// @Failure      500      {object}  responsedto.ErrorResponse
// @Security     BearerAuth
// @Router       /organizations/ticket [get]
func (t *OrganizationTicketHandler) GetTicketsList(w http.ResponseWriter, r *http.Request) {
	filter := utils.ParsePagination(r)
	user, _ := t.jwtService.GetUserFromContext(r.Context())

	result, err := t.service.GetTicketsList(user, filter)
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

	logger.InfoLog("Tickets fetched successfully", map[string]any{
		"count": len(result.Tickets),
	})
	utils.WriteJSONResponse(w, http.StatusOK, result)
}

// UpdateTicket godoc
// @Summary      Update ticket
// @Description  Update a ticket's details
// @Tags         organization-tickets
// @Accept       json
// @Produce      json
// @Param        id path int true "Ticket ID"
// @Param        request body requestdto.UpdateTicketRequest true "Update Ticket Request"
// @Success      200  {object}  responsedto.TicketResponse
// @Failure      400  {object}  responsedto.ErrorResponse
// @Failure      404  {object}  responsedto.ErrorResponse
// @Security     BearerAuth
// @Router       /organizations/ticket/{id} [put]
func (t *OrganizationTicketHandler) UpdateTicket(w http.ResponseWriter, r *http.Request) {
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

	user, _ := t.jwtService.GetUserFromContext(r.Context())

	err = t.service.UpdateTicket(user, uint(id), req)
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

	result := responsedto.CommonResponse{
		Message: "Ticket updated successfully",
	}
	logger.InfoLog("Ticket updated successfully", result)
	utils.WriteJSONResponse(w, http.StatusOK, result)
}
