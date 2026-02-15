package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"DewaSRY/sociomile-app/internal/services"
	_ "DewaSRY/sociomile-app/pkg/dtos/filtersdto"
	"DewaSRY/sociomile-app/pkg/dtos/requestdto"
	"DewaSRY/sociomile-app/pkg/dtos/responsedto"
	"DewaSRY/sociomile-app/pkg/lib/logger"
	"DewaSRY/sociomile-app/pkg/utils"

	"github.com/go-chi/chi/v5"
)

type OrganizationHandler struct {
	service             services.OrganizationCrudService
}

func NewOrganizationHandler(
	service services.OrganizationCrudService,
) *OrganizationHandler {
	return &OrganizationHandler{
		service:             service,
	}
}
// CreateOrganization godoc
// @Summary      Create a new organization
// @Description  Create a new organization with owner (Super Admin only)
// @Tags         organizations
// @Accept       json
// @Produce      json
// @Param        request body requestdto.CreateOrganizationRequest true "Create Organization Request"
// @Success      201  {object}  responsedto.OrganizationResponse
// @Failure      400  {object}  responsedto.ErrorResponse
// @Failure      500  {object}  responsedto.ErrorResponse
// @Security     BearerAuth
// @Router       /organizations [post]
func (h *OrganizationHandler) CreateOrganization(w http.ResponseWriter, r *http.Request) {
	var req requestdto.CreateOrganizationRequest

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

	result, err := h.service.CreateOrganization(req)
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

	logger.InfoLog("Organization created successfully", map[string]any{
		"organization_id": result.ID,
	})
	utils.WriteJSONResponse(w, http.StatusCreated, result)
}

// GetOrganization godoc
// @Summary      Get organization by ID
// @Description  Retrieve an organization by its ID
// @Tags         organizations
// @Accept       json
// @Produce      json
// @Param        id path int true "Organization ID"
// @Success      200  {object}  responsedto.OrganizationResponse
// @Failure      400  {object}  responsedto.ErrorResponse
// @Failure      404  {object}  responsedto.ErrorResponse
// @Security     BearerAuth
// @Router       /organizations/{id} [get]
func (h *OrganizationHandler) GetOrganization(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseUint(idStr, 10, 32)
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

	result, err := h.service.GetOrganizationByID(uint(id))
	if err != nil {
		errorData := responsedto.ErrorResponse{
			Message: "organization not found",
			Error:   err.Error(),
			Code:    http.StatusNotFound,
		}
		logger.ErrorLog("Organization not found", errorData)
		utils.WriteJSONResponse(w, http.StatusNotFound, errorData)
		return
	}

	utils.WriteJSONResponse(w, http.StatusOK, result)
}

// GetAllOrganizations godoc
// @Summary      Get all organizations
// @Description  Retrieve all organizations
// @Tags         organizations
// @Accept       json
// @Produce      json
// @Success      200  {object}  responsedto.OrganizationListResponse
// @Failure      500  {object}  responsedto.ErrorResponse
// @Security     BearerAuth
// @Router       /organizations [get]
func (h *OrganizationHandler) GetAllOrganizations(w http.ResponseWriter, r *http.Request) {
	result, err := h.service.GetAllOrganizations()
	if err != nil {
		errorData := responsedto.ErrorResponse{
			Message: "failed to fetch organizations",
			Error:   err.Error(),
			Code:    http.StatusInternalServerError,
		}
		logger.ErrorLog("Failed to fetch organizations", errorData)
		utils.WriteJSONResponse(w, http.StatusInternalServerError, errorData)
		return
	}

	utils.WriteJSONResponse(w, http.StatusOK, result)
}

// UpdateOrganization godoc
// @Summary      Update organization
// @Description  Update an organization's details
// @Tags         organizations
// @Accept       json
// @Produce      json
// @Param        id path int true "Organization ID"
// @Param        request body requestdto.UpdateOrganizationRequest true "Update Organization Request"
// @Success      200  {object}  responsedto.OrganizationResponse
// @Failure      400  {object}  responsedto.ErrorResponse
// @Failure      404  {object}  responsedto.ErrorResponse
// @Security     BearerAuth
// @Router       /organizations/{id} [put]
func (h *OrganizationHandler) UpdateOrganization(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseUint(idStr, 10, 32)
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

	var req requestdto.UpdateOrganizationRequest
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

	result, err := h.service.UpdateOrganization(uint(id), req)
	if err != nil {
		errorData := responsedto.ErrorResponse{
			Message: "failed to update organization",
			Error:   err.Error(),
			Code:    http.StatusInternalServerError,
		}
		logger.ErrorLog("Failed to update organization", errorData)
		utils.WriteJSONResponse(w, http.StatusInternalServerError, errorData)
		return
	}

	logger.InfoLog("Organization updated successfully", map[string]any{
		"organization_id": result.ID,
	})
	utils.WriteJSONResponse(w, http.StatusOK, result)
}

// DeleteOrganization godoc
// @Summary      Delete organization
// @Description  Delete an organization
// @Tags         organizations
// @Accept       json
// @Produce      json
// @Param        id path int true "Organization ID"
// @Success      200  {object}  responsedto.SuccessResponse
// @Failure      400  {object}  responsedto.ErrorResponse
// @Failure      404  {object}  responsedto.ErrorResponse
// @Security     BearerAuth
// @Router       /organizations/{id} [delete]
func (h *OrganizationHandler) DeleteOrganization(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseUint(idStr, 10, 32)
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

	if err := h.service.DeleteOrganization(uint(id)); err != nil {
		errorData := responsedto.ErrorResponse{
			Message: "failed to delete organization",
			Error:   err.Error(),
			Code:    http.StatusInternalServerError,
		}
		logger.ErrorLog("Failed to delete organization", errorData)
		utils.WriteJSONResponse(w, http.StatusInternalServerError, errorData)
		return
	}

	logger.InfoLog("Organization deleted successfully", map[string]any{
		"organization_id": id,
	})
	utils.WriteJSONResponse(w, http.StatusOK, responsedto.SuccessResponse{
		Message: "organization deleted successfully",
	})
}

// GetOrganizationStats godoc
// @Summary      Get organization statistics
// @Description  Get statistics for an organization (conversations, tickets)
// @Tags         organizations
// @Accept       json
// @Produce      json
// @Param        id path int true "Organization ID"
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  responsedto.ErrorResponse
// @Failure      404  {object}  responsedto.ErrorResponse
// @Security     BearerAuth
// @Router       /organizations/{id}/stats [get]
func (h *OrganizationHandler) GetOrganizationStats(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseUint(idStr, 10, 32)
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

	result, err := h.service.GetOrganizationStats(uint(id))
	if err != nil {
		errorData := responsedto.ErrorResponse{
			Message: "failed to fetch organization stats",
			Error:   err.Error(),
			Code:    http.StatusInternalServerError,
		}
		logger.ErrorLog("Failed to fetch organization stats", errorData)
		utils.WriteJSONResponse(w, http.StatusInternalServerError, errorData)
		return
	}

	utils.WriteJSONResponse(w, http.StatusOK, result)
}
