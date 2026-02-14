package middleware

import (
	"DewaSRY/sociomile-app/internal/services"
	"DewaSRY/sociomile-app/pkg/dtos/responsedto"
	jwtutil "DewaSRY/sociomile-app/pkg/lib/jwt"
	"DewaSRY/sociomile-app/pkg/lib/logger"
	"DewaSRY/sociomile-app/pkg/utils"
	"fmt"
	"net/http"
)

func Authorize(jwtSvc jwtutil.JwtService, authorizeSvc services.AuthorizeService, allowedRoles []string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user, err := jwtSvc.GetUserFromContext(r.Context())
			
			fmt.Println(" the calim get ", user, err)
			fmt.Println("do thete is error ", err)

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

			fmt.Println("user role is ", user.RoleID)
			fmt.Print("test test ")
			if err := authorizeSvc.IsUserAuthorize(1, allowedRoles); err != nil {
				errorResponse := responsedto.ErrorResponse{
					Message: "Not authorize",
					Error:   "Not authorize",
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
