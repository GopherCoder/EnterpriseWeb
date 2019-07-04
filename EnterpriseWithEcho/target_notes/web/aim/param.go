package aim

import "gopkg.in/go-playground/validator.v9"

type createParam struct {
	Data struct {
		Title string `json:"title"`
	} `json:"data" validate:"required"`
}

func (c createParam) Valid() error {
	return validator.New().Struct(c)
}

type patchParam struct {
	Data struct {
		Level       string `json:"level"`
		Status      string `json:"status"`
		Description string `json:"description"`
	} `json:"data" validate:"required"`
}

func (p patchParam) Valid() error {
	return validator.New().Struct(p)
}
