package filtersdto

type FiltersDto struct {
	Page  *int `json:"page" validate:"min=1"`
	Limit *int `json:"limit" validate:"min=1"`
}
