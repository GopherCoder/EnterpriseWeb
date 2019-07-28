package model

import "gopkg.in/go-playground/validator.v9"

type CreateAddressParams struct {
	AdminId int64  `json:"admin_id" valid:"required"`
	Detail  string `json:"detail" valid:"required, min=1"`
}

func (c CreateAddressParams) Valid() error {
	return validator.New().Struct(c)
}
