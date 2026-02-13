package handlers

import (
	"encoding/json"
	"net/http"

	"DewaSRY/sociomile-app/internal/services"
	serviceImpl "DewaSRY/sociomile-app/internal/services/impl"
	"DewaSRY/sociomile-app/pkg/dtos/requestdto"
	"DewaSRY/sociomile-app/pkg/dtos/responsedto"
	"DewaSRY/sociomile-app/pkg/lib/logger"
	"DewaSRY/sociomile-app/pkg/utils"
)

type AuthHandler struct {
	service services.AuthService
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{
		service: serviceImpl.InstanceAuthServiceImpl(),
	}
}

// Register godoc
// @Summary      Register a new user
// @Description  Register a new user with email, password and name
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body   requestdto.RegisterRequest true "Register Request"
// @Success      201  {object}  responsedto.AuthResponse
// @Failure      400  {object}  responsedto.ErrorResponse
// @Failure      409  {object}  responsedto.ErrorResponse
// @Router       /auth/register [post]
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
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
	}

	result, err := h.service.Register(req)

	if err != nil {
		errorData := responsedto.ErrorResponse{
			Message: "invalid request",
			Error:   err.Error(),
			Code:    http.StatusConflict,
		}

		logger.ErrorLog("Failed to register user", errorData)
		utils.WriteJSONResponse(w, http.StatusConflict, errorData)
		return
	}

	logger.InfoLog("User registered successfully", map[string]any{
		"email": req.Email,
	})
	utils.WriteJSONResponse(w, http.StatusCreated, result)
}

// Login godoc
// @Summary      Login user
// @Description  Authenticate user with email and password
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body requestdto.LoginRequest true "Login Request"
// @Success      201  {object}  responsedto.AuthResponse
// @Failure      400  {object}  responsedto.ErrorResponse
// @Failure      409  {object}  responsedto.ErrorResponse
// @Router       /auth/login [post]
// Login handles user authentication
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req requestdto.LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		errorData := responsedto.ErrorResponse{
			Message: "invalid request",
			Error:   err.Error(),
		}
		logger.ErrorLog("Failed to decode request", errorData)
		utils.WriteJSONResponse(w, http.StatusBadRequest, errorData)
		return
	}

	if err := utils.ValidateStruct(req); err != nil {
		errorData := responsedto.ErrorResponse{
			Message: "invalid request",
			Error:   err.Error(),
		}

		logger.ErrorLog("Failed to validate request", errorData)
		utils.WriteJSONResponse(w, http.StatusBadRequest, errorData)
	}
	result, err := h.service.Login(req)
	if err != nil {
		errorData := responsedto.ErrorResponse{
			Message: "invalid request",
			Error:   err.Error(),
		}
		logger.ErrorLog("Failed to login user", errorData)
		utils.WriteJSONResponse(w, http.StatusConflict, errorData)
		return
	}

	logger.InfoLog("User login successfully", map[string]any{
		"Message": req.Email,
	})
	utils.WriteJSONResponse(w, http.StatusCreated, result)
}

// GetProfile godoc
// @Summary      Get user profile
// @Description  Get authenticated user's profile information
// @Tags         auth
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      201  {object}  responsedto.UserData
// @Failure      401  {object}  responsedto.ErrorResponse
// @Router       /auth/profile [get]
// GetProfile retrieves the authenticated user's profile
func (h *AuthHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	// Get user ID from context (set by JWT middleware)
	userID, ok := r.Context().Value("userID").(uint)

	if !ok {
		errorData := responsedto.ErrorResponse{
			Message: "User not authenticated",
			Error:   "Unauthorized",
		}
		logger.ErrorLog("Failed to authenticate", errorData)
		utils.WriteJSONResponse(w, http.StatusUnauthorized, errorData)
		return
	}

	user, err := h.service.GetUserByID(userID)
	if err != nil {
		errorData := responsedto.ErrorResponse{
			Message: "User not authenticated",
			Error:   err.Error(),
		}

		logger.ErrorLog("User not authenticated", errorData)
		utils.WriteJSONResponse(w, http.StatusUnauthorized, errorData)
		return
	}

	result := responsedto.UserData{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
	}

	logger.InfoLog("User login successfully", result)
	utils.WriteJSONResponse(w, http.StatusCreated, result)
}

// RefreshToken godoc
// @Summary      Refresh JWT token
// @Description  Generate a new JWT token using existing token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body requestdto.RefreshTokenRequest true "Refresh Token Request"
// @Success      201  {object}  map[string]string
// @Failure      400  {object}  responsedto.ErrorResponse
// @Router       /auth/refresh [post]
// RefreshToken generates a new JWT token
func (h *AuthHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	var req requestdto.RefreshTokenRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		errorData := responsedto.ErrorResponse{
			Message: "invalid request",
			Error:   err.Error(),
		}

		logger.ErrorLog("Failed to decode request", errorData)
		utils.WriteJSONResponse(w, http.StatusBadRequest, errorData)
		return
	}

	if req.Token == "" {
		errorData := responsedto.ErrorResponse{
			Message: "Token is required",
			Error:   "Bad Request",
		}
		logger.ErrorLog("Failed to decode request", errorData)
		utils.WriteJSONResponse(w, http.StatusBadRequest, errorData)
		return
	}

	newToken, err := h.service.RefreshToken(req.Token)
	if err != nil {
		errorData := responsedto.ErrorResponse{
			Message: "Invalid or expired token",
			Error:   err.Error(),
		}

		logger.ErrorLog("Invalid or expired token", errorData)
		utils.WriteJSONResponse(w, http.StatusBadRequest, errorData)
		return
	}

	result := map[string]string{
		"token": newToken,
	}

	logger.InfoLog("User registered successfully", result)
	utils.WriteJSONResponse(w, http.StatusCreated, result)
}
