package handlers

import (
	"DewaSRY/sociomile-app/internal/services"
	_ "DewaSRY/sociomile-app/pkg/dtos/filtersdto"
	"DewaSRY/sociomile-app/pkg/dtos/requestdto"
	"DewaSRY/sociomile-app/pkg/dtos/responsedto"
	jwtLib "DewaSRY/sociomile-app/pkg/lib/jwt"
	"DewaSRY/sociomile-app/pkg/lib/logger"
	"DewaSRY/sociomile-app/pkg/utils"
	"encoding/json"
	"net/http"
)

type GuestConversationHandler struct {
	jwtService           jwtLib.JwtService
	guestConversationSvc services.GuestConversationService
}

func NewGuestConversationHandler(
	jwtService jwtLib.JwtService,
	guestConversationSvc services.GuestConversationService,
) *GuestConversationHandler {
	return &GuestConversationHandler{
		jwtService:           jwtService,
		guestConversationSvc: guestConversationSvc,
	}
}

// CreateConversation godoc
// @Summary      Create a new conversation
// @Description  Create a new conversation with an organization (Guest user)
// @Tags         guest-conversation
// @Accept       json
// @Produce      json
// @Param        request body requestdto.CreateConversationRequest true "Create Conversation Request"
// @Success      201  {object}  responsedto.CommonResponse
// @Failure      400  {object}  responsedto.ErrorResponse
// @Failure      500  {object}  responsedto.ErrorResponse
// @Security     BearerAuth
// @Router       /guest/conversations [post]
func (t *GuestConversationHandler) CreateConversation(w http.ResponseWriter, r *http.Request) {
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
	user, _ := t.jwtService.GetUserFromContext(r.Context())
	err := t.guestConversationSvc.CreateConversation(user, req)
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

	result := responsedto.CommonResponse{
		Message: "Success Create Conversation",
		Code:    http.StatusCreated,
	}
	utils.WriteJSONResponse(w, http.StatusCreated, result)
}

// GetConversation godoc
// @Summary      Get guest conversation 
// @Description  Retrieve a guest conversation 
// @Tags         guest-conversation
// @Accept       json
// @Produce      json
// @Param        request  query  filtersdto.FiltersDto  false  "Pagination query"
// @Security     BearerAuth
// @Success      200  {object}  responsedto.ConversationListPaginateResponse
// @Failure      400  {object}  responsedto.ErrorResponse
// @Failure      404  {object}  responsedto.ErrorResponse
// @Router       /guest/conversations [get]
func (t *GuestConversationHandler) GetConversation(w http.ResponseWriter, r *http.Request) {
	filter := utils.ParsePagination(r)
	user, _ := t.jwtService.GetUserFromContext(r.Context())
	result, err := t.guestConversationSvc.GetConversation(user,filter)
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
