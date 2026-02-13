package handlers

import (
	"encoding/json"
	"net/http"

	"DewaSRY/sociomile-app/internal/services"
	"DewaSRY/sociomile-app/pkg/dtos/requestdto"
	"DewaSRY/sociomile-app/pkg/dtos/responsedto"
	"DewaSRY/sociomile-app/pkg/lib/logger"
	"DewaSRY/sociomile-app/pkg/lib/utils"
)

type AuthHandler struct {
	service *services.AuthService
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{
		service: services.NewAuthService(),
	}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req requestdto.RegisterRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		errorData := map[string]any{
			"message": "invalid request",
			"error":   err.Error(),
		}

		logger.ErrorLog("Failed to decode request", errorData)
		utils.WriteErrorResponse(w, http.StatusBadRequest, errorData)
		return
	}

	if err := utils.ValidateStruct(req); err != nil {
		errorData := map[string]any{
			"message": "invalid request",
			"error":   err.Error(),
		}

		logger.ErrorLog("Failed to validate request", errorData)
		utils.WriteErrorResponse(w, http.StatusBadRequest, errorData)
	}

	result, err := h.service.Register(req)

	if err != nil {
		errorData := map[string]any{
			"message": "invalid request",
			"error":   err.Error(),
		}
		logger.ErrorLog("Failed to register user", errorData)
		utils.WriteErrorResponse(w, http.StatusConflict, errorData)
		return
	}

	logger.InfoLog("User registered successfully", map[string]any{
		"email": req.Email,
	})
	utils.WriteJSONResponse(w, http.StatusCreated, result)
}

// Login handles user authentication
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req requestdto.LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		errorData := map[string]any{
			"message": "invalid request",
			"error":   err.Error(),
		}
		logger.ErrorLog("Failed to decode request", errorData)
		utils.WriteErrorResponse(w, http.StatusBadRequest, errorData)
		return
	}

	if err := utils.ValidateStruct(req); err != nil {
		errorData := map[string]any{
			"message": "invalid request",
			"error":   err.Error(),
		}

		logger.ErrorLog("Failed to validate request", errorData)
		utils.WriteErrorResponse(w, http.StatusBadRequest, errorData)
	}
	result, err := h.service.Login(req)
	if err != nil {
		errorData := map[string]any{
			"message": "invalid request",
			"error":   err.Error(),
		}
		logger.ErrorLog("Failed to login user", errorData)
		utils.WriteErrorResponse(w, http.StatusConflict, errorData)
		return
	}

	logger.InfoLog("User login successfully", map[string]any{
		"email": req.Email,
	})
	utils.WriteJSONResponse(w, http.StatusCreated, result)
}

// GetProfile retrieves the authenticated user's profile
func (h *AuthHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	// Get user ID from context (set by JWT middleware)
	userID, ok := r.Context().Value("userID").(uint)

	if !ok {
		errorData := map[string]any{
			"message": "User not authenticated",
			"error":   "Unauthorized",
		}
		logger.ErrorLog("Failed to authenticate", errorData)
		utils.WriteErrorResponse(w, http.StatusUnauthorized, errorData)
		return
	}

	user, err := h.service.GetUserByID(userID)
	if err != nil {
		errorData := map[string]any{
			"message": "User not authenticated",
			"error":   err.Error(),
		}

		logger.ErrorLog("User not authenticated", errorData)
		utils.WriteErrorResponse(w, http.StatusUnauthorized, errorData)
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

// RefreshToken generates a new JWT token
func (h *AuthHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	var req requestdto.RefreshTokenRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		errorData := map[string]any{
			"message": "invalid request",
			"error":   err.Error(),
		}

		logger.ErrorLog("Failed to decode request", errorData)
		utils.WriteErrorResponse(w, http.StatusBadRequest, errorData)
		return
	}

	if req.Token == "" {
		errorData := map[string]any{
			"message": "Token is required",
			"error":   "Bad Request",
		}
		logger.ErrorLog("Failed to decode request", errorData)
		utils.WriteErrorResponse(w, http.StatusBadRequest, errorData)
		return
	}

	newToken, err := h.service.RefreshToken(req.Token)
	if err != nil {
		errorData := map[string]any{
			"message": "Invalid or expired token",
			"error":   err.Error(),
		}

		logger.ErrorLog("Invalid or expired token", errorData)
		utils.WriteErrorResponse(w, http.StatusBadRequest, errorData)
		return
	}

	result:= map[string]string{
		"token": newToken,
	}

	logger.InfoLog("User registered successfully", result)
	utils.WriteJSONResponse(w, http.StatusCreated, result)
}
