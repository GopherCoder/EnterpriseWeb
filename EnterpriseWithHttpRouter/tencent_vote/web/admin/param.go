package admin

import "gopkg.in/go-playground/validator.v9"

type RegisterParams struct {
	Data struct {
		Phone    string `json:"phone" validate:"required"`
		Password string `json:"password" validate:"min=8"`
	} `json:"data" validate:"required"`
}

func (r RegisterParams) Valid() error {
	return validator.New().Struct(r)
}
