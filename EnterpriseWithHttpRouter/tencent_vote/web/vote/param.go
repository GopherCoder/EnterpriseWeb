package vote

import "gopkg.in/go-playground/validator.v9"

type CreateParam struct {
	Title       string   `json:"title" validate:"required"`
	Description string   `json:"description"`
	Choice      []string `json:"choice" validate:"min=2"`
	Date        string   `json:"date"` // Default: one year、时间戳
	IsAnonymous string   `json:"is_anonymous" validate:"required"`
}

func (c CreateParam) Valid() error {
	return validator.New().Struct(c)
}
