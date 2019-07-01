package admin

import "gopkg.in/go-playground/validator.v9"

type registerParam struct {
	AccountName string `json:"account_name" validate:"required"`
	Password    string `json:"password" validate:"required"`
}

func (r registerParam) Valid() error {
	return validator.New().Struct(r)
}
