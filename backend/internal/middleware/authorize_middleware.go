package middleware

import (
	"DewaSRY/sociomile-app/internal/services"
	"DewaSRY/sociomile-app/pkg/dtos/responsedto"
	jwtutil "DewaSRY/sociomile-app/pkg/lib/jwt"
	"DewaSRY/sociomile-app/pkg/lib/logger"
	"DewaSRY/sociomile-app/pkg/utils"
	"net/http"
)

func Authorize(jwtSvc jwtutil.JwtService, authorizeSvc services.AuthorizeService, allowedRoles []string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user, err := jwtSvc.GetUserFromContext(r.Context())

			if err {
				errorResponse := responsedto.ErrorResponse{
					Message: "Failed to parse data",
					Error:   "Failed to parse data",
					Code:    http.StatusInternalServerError,
				}
				logger.ErrorLog("Failed to parse data", errorResponse)
				utils.WriteJSONResponse(w, http.StatusInternalServerError, errorResponse)
				return
			}
			if err := authorizeSvc.IsUserAuthorize(user.RoleID, allowedRoles); err != nil {
				errorResponse := responsedto.ErrorResponse{
					Message: "Not authorize",
					Error:   err.Error(),
					Code:    http.StatusInternalServerError,
				}
				logger.ErrorLog("Not authorize", errorResponse)
				utils.WriteJSONResponse(w, http.StatusForbidden, errorResponse)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
