package handlers

import (
	"encoding/json"
	"net/http"

	"DewaSRY/sociomile-app/internal/services"
	_ "DewaSRY/sociomile-app/pkg/dtos/filtersdto"
	"DewaSRY/sociomile-app/pkg/dtos/requestdto"
	"DewaSRY/sociomile-app/pkg/dtos/responsedto"
	jwtLib "DewaSRY/sociomile-app/pkg/lib/jwt"
	"DewaSRY/sociomile-app/pkg/lib/logger"
	"DewaSRY/sociomile-app/pkg/utils"
)

type OrganizationStaffHandler struct {
	jwtService          jwtLib.JwtService
	serviceOrganization services.OrganizationService
}

func NewOrganizationStaffHandler(
	jwtService jwtLib.JwtService,
	serviceOrganization services.OrganizationService,
) *OrganizationStaffHandler {
	return &OrganizationStaffHandler{
		jwtService:             jwtService,
		serviceOrganization: serviceOrganization,
	}
}

// CreateOrganizationStaff godoc
// @Summary      Create a new organization staff
// @Description  Create a new organization staff (Super Admin only)
// @Tags         organizations-staff
// @Accept       json
// @Produce      json
// @Param        request body requestdto.RegisterRequest true "Create Organization staff"
// @Success      201  {object}  responsedto.OrganizationResponse
// @Failure      400  {object}  responsedto.ErrorResponse
// @Failure      500  {object}  responsedto.ErrorResponse
// @Security     BearerAuth
// @Router       /organizations/staff [post]
func (t *OrganizationStaffHandler) CreateOrganizationStaff(w http.ResponseWriter, r *http.Request) {
	var req requestdto.RegisterRequest

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

	err := t.serviceOrganization.CreateStaff(req, user)
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

	result := responsedto.CommonResponse{
		Message: "Create staff success",
	}
	logger.InfoLog("Organization created successfully", result)
	utils.WriteJSONResponse(w, http.StatusCreated, result)
}

// GetStaffListPagination godoc
// @Summary      Get organizations staff with pagination
// @Description  Retrieve list of organizations staff with pagination support
// @Tags         organizations-staff
// @Accept       json
// @Produce      json
// @Param        request  query  filtersdto.FiltersDto  false  "Pagination query"
// @Security     BearerAuth
// @Success      200      {object}  responsedto.OrganizationPaginateResponse
// @Failure      500      {object}  responsedto.ErrorResponse
// @Router       /organizations/staff [get]
func (t *OrganizationStaffHandler) GetStaffListPagination(w http.ResponseWriter, r *http.Request) {

	filter := utils.ParsePagination(r)
	user, _ := t.jwtService.GetUserFromContext(r.Context())
	result, err := t.serviceOrganization.GetStaffList(filter, user)
	if err != nil {
		errorData := responsedto.ErrorResponse{
			Message: "failed to get organization staff",
			Error:   err.Error(),
			Code:    http.StatusInternalServerError,
		}
		logger.ErrorLog("failed to get organization staff", errorData)
		utils.WriteJSONResponse(w, http.StatusInternalServerError, errorData)
		return
	}

	logger.InfoLog("success to get organization staff", result)
	utils.WriteJSONResponse(w, http.StatusCreated, result)
}
