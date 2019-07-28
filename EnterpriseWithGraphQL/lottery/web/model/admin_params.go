package model

import "gopkg.in/go-playground/validator.v9"

type LoginParam struct {
	Phone    string `json:"phone" valid:"required, len=11"`
	Password string `json:"password" valid:"required, min=8"`
}

func (L LoginParam) Valid() error {
	return validator.New().Struct(L)
}
