package services

import (
	"DewaSRY/sociomile-app/pkg/dtos/filtersdto"
	"DewaSRY/sociomile-app/pkg/dtos/requestdto"
	"DewaSRY/sociomile-app/pkg/dtos/responsedto"
	jwtLib "DewaSRY/sociomile-app/pkg/lib/jwt"
)


type OrganizationService interface{
	CreateStaff(requestdto.RegisterRequest, *jwtLib.Claims) error
	GetStaffList(filtersdto.FiltersDto, *jwtLib.Claims) (*responsedto.OrganizationStaffPagination, error)
}