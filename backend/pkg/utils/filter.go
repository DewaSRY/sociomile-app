package utils

import (
	"DewaSRY/sociomile-app/pkg/dtos/filtersdto"
	"net/http"
	"strconv"
)

func ParsePagination(r *http.Request) filtersdto.FiltersDto {
	query := r.URL.Query()

	page, _ := strconv.Atoi(query.Get("page"))
	limit, _ := strconv.Atoi(query.Get("limit"))

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	return filtersdto.FiltersDto{
		Page:  &page,
		Limit: &limit,
	}
}