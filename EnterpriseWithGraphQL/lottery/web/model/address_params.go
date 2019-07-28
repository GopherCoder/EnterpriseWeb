package model

import "gopkg.in/go-playground/validator.v9"

type CreateAddressParams struct {
	AdminId int64  `json:"admin_id" validate:"required"`
	Detail  string `json:"detail" validate:"required, min=1"`
}

func (c CreateAddressParams) Valid() error {
	return validator.New().Struct(c)
}

type GetAddressParams struct {
	AdminId int64  `json:"admin_id" Validate:"required"`
	OrderBy string `json:"order_by"`
	Limit   int    `json:"limit" Validate:""`
	Offset  int    `json:"offset" Validate:""`
}
