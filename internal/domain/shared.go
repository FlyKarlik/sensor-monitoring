package domain

type PaginationInput struct {
	Page  *int `json:"page,omitempty"`
	Limit *int `json:"limit,omitempty"`
}

type SortInput struct {
	IsReverse *bool   `json:"is_reverse,omitempty"`
	Field     *string `json:"field,omitempty"`
}
