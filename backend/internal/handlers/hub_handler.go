package handlers

import (
	"DewaSRY/sociomile-app/internal/services"
	_ "DewaSRY/sociomile-app/pkg/dtos/filtersdto"
	"DewaSRY/sociomile-app/pkg/dtos/requestdto"
	"DewaSRY/sociomile-app/pkg/dtos/responsedto"
	"DewaSRY/sociomile-app/pkg/lib/logger"
	"DewaSRY/sociomile-app/pkg/utils"
	"encoding/json"
	"net/http"
)


type HubHandler struct {
	service services.HubService
}

func NewHubHandler(
	service services.HubService,
) *HubHandler {
	return &HubHandler{
		service: service,
	}
}

// CreateOrganization godoc
// @Summary      Create a new organization
// @Description  Create a new organization with owner (Super Admin only)
// @Tags         Hub
// @Accept       json
// @Produce      json
// @Param        request body requestdto.RegisterOrganizationRequest true "Create Organization Request"
// @Success      201  {object}  responsedto.CommonResponse
// @Failure      400  {object}  responsedto.ErrorResponse
// @Failure      500  {object}  responsedto.ErrorResponse
// @Security     BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer {your token}" to authorize
// @Router       /hub/organizations [post]
func (h *HubHandler) CreateOrganization(w http.ResponseWriter, r *http.Request) {
	var req requestdto.RegisterOrganizationRequest

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

	 err := h.service.CreateOrganization(req)
	if err != nil {
		errorData := responsedto.ErrorResponse{
			Message: "failed to create organization",
			Error:   err.Error(),
			Code:    http.StatusInternalServerError,
		}
		logger.ErrorLog("Failed to create organization", errorData)
		utils.WriteJSONResponse(w, http.StatusInternalServerError, errorData)
		return
	}
	result:= responsedto.CommonResponse{
		Message: "Success Create Organization",
		Code: http.StatusCreated,
	}

	logger.InfoLog("Organization created successfully",result)
	utils.WriteJSONResponse(w, http.StatusCreated, result)
}


// GetOrganizationPagination godoc
// @Summary      Get organizations with pagination
// @Description  Retrieve list of organizations with pagination support
// @Tags         Hub
// @Accept       json
// @Produce      json
// @Param        request  query  filtersdto.FiltersDto  false  "Pagination query"
// @Security     BearerAuth
// @Success      200      {object}  responsedto.OrganizationPaginateResponse
// @Failure      500      {object}  responsedto.ErrorResponse
// @Router       /hub/organizations [get]
func (h *HubHandler) GetOrganizationPagination(w http.ResponseWriter, r *http.Request) {

	filter := utils.ParsePagination(r)
	result ,err := h.service.GetOrganizationPagination(filter)
	if err != nil {
		errorData := responsedto.ErrorResponse{
			Message: "failed to get organization",
			Error:   err.Error(),
			Code:    http.StatusInternalServerError,
		}
		logger.ErrorLog("Failed to create organization", errorData)
		utils.WriteJSONResponse(w, http.StatusInternalServerError, errorData)
		return
	}

	logger.InfoLog("Organization created successfully",result)
	utils.WriteJSONResponse(w, http.StatusCreated, result)
}
