package concept

import "gopkg.in/go-playground/validator.v9"

type CreateConcept struct {
	Data struct {
		Key    string `json:"key"`
		Detail string `json:"detail"`
	} `json:"data" validate:"required"`
}

func (c CreateConcept) Valid() error {
	return validator.New().Struct(c)
}
