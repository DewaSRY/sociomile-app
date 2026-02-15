package handlers

import (
	"net/http"

	"DewaSRY/sociomile-app/internal/services"
	_ "DewaSRY/sociomile-app/pkg/dtos/filtersdto"
	"DewaSRY/sociomile-app/pkg/dtos/responsedto"
	"DewaSRY/sociomile-app/pkg/lib/logger"
	"DewaSRY/sociomile-app/pkg/utils"
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



