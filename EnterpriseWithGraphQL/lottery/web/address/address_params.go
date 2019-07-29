package address

import "gopkg.in/go-playground/validator.v9"

type CreateAddressParams struct {
	AdminId int64  `json:"admin_id" validate:"required"`
	Detail  string `json:"detail" validate:"min=1"`
}

func (c CreateAddressParams) Valid() error {
	return validator.New().Struct(c)
}

type GetAddressParams struct {
	AdminId int64  `json:"admin_id" validate:"required"`
	OrderBy string `json:"order_by"`
	Limit   int    `json:"limit" validate:"min=1"`
	Offset  int    `json:"offset" validate:"min=0"`
}

func (G GetAddressParams) Valid() error {
	return validator.New().Struct(G)
}

type UpdateAddressParams struct {
	Id     int64  `json:"id" validate:"required"`
	Detail string `json:"detail" validate:"required"`
}

func (U UpdateAddressParams) Valid() error {
	return validator.New().Struct(U)
}
