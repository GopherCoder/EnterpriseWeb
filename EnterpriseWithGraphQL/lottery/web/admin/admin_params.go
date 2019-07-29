package admin

import "gopkg.in/go-playground/validator.v9"

type LoginParam struct {
	Phone    string `json:"phone" validate:"len=11"`
	Password string `json:"password" validate:"required,min=8"`
}

func (L LoginParam) Valid() error {
	return validator.New().Struct(L)
}

type UpdateAdminParams struct {
	AdminId int64  `json:"admin_id" validate:"required"`
	Name    string `json:"name" validate:"required"`
}

func (U UpdateAdminParams) Valid() error {
	return validator.New().Struct(U)
}
