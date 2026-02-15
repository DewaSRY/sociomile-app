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
	"strconv"

	"github.com/go-chi/chi/v5"
)

type GuestMessageHandler struct {
	jwtService        jwtLib.JwtService
	guestMessageSvc   services.GuestMessageService
}

func NewGuestMessageHandler(
	jwtService jwtLib.JwtService,
	guestMessageSvc services.GuestMessageService,
) *GuestMessageHandler {
	return &GuestMessageHandler{
		jwtService:        jwtService,
		guestMessageSvc:   guestMessageSvc,
	}
}

// SendConversationMessage godoc
// @Summary      Send a message in a conversation
// @Description  Guest user sends a message in their conversation
// @Tags         guest-messages
// @Accept       json
// @Produce      json
// @Param        request body requestdto.CreateConversationMessageRequest true "Send Message Request"
// @Success      201  {object}  responsedto.CommonResponse
// @Failure      400  {object}  responsedto.ErrorResponse
// @Failure      500  {object}  responsedto.ErrorResponse
// @Security     BearerAuth
// @Router       /guest/conversations/messages [post]
func (t *GuestMessageHandler) SendConversationMessage(w http.ResponseWriter, r *http.Request) {
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

	user, _ := t.jwtService.GetUserFromContext(r.Context())
	err := t.guestMessageSvc.SendConversationMessage(user, req)
	if err != nil {
		errorData := responsedto.ErrorResponse{
			Message: "failed to send message",
			Error:   err.Error(),
			Code:    http.StatusInternalServerError,
		}
		logger.ErrorLog("Failed to send message", errorData)
		utils.WriteJSONResponse(w, http.StatusInternalServerError, errorData)
		return
	}

	result := responsedto.CommonResponse{
		Message: "Message sent successfully",
		Code:    http.StatusCreated,
	}
	logger.InfoLog("Message sent successfully", result)
	utils.WriteJSONResponse(w, http.StatusCreated, result)
}

// GetConversationMessageList godoc
// @Summary      Get conversation messages
// @Description  Retrieve messages for a specific conversation
// @Tags         guest-messages
// @Accept       json
// @Produce      json
// @Param        id path int true "Conversation ID"
// @Param        request  query  filtersdto.FiltersDto  false  "Pagination query"
// @Success      200  {object}  responsedto.ConversationMessagePaginateResponse
// @Failure      400  {object}  responsedto.ErrorResponse
// @Failure      404  {object}  responsedto.ErrorResponse
// @Security     BearerAuth
// @Router       /guest/conversations/{id}/messages [get]
func (t *GuestMessageHandler) GetConversationMessageList(w http.ResponseWriter, r *http.Request) {
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

	filter := utils.ParsePagination(r)
	user, _ := t.jwtService.GetUserFromContext(r.Context())

	result, err := t.guestMessageSvc.GetConversationMessageList(user, filter, uint(id))
	if err != nil {
		errorData := responsedto.ErrorResponse{
			Message: "failed to fetch messages",
			Error:   err.Error(),
			Code:    http.StatusNotFound,
		}
		logger.ErrorLog("Failed to fetch messages", errorData)
		utils.WriteJSONResponse(w, http.StatusNotFound, errorData)
		return
	}

	logger.InfoLog("Messages fetched successfully",result)
	utils.WriteJSONResponse(w, http.StatusOK, result)
}
