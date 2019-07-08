package company

import "gopkg.in/go-playground/validator.v9"

type getCompaniesParam struct {
	Top         int    `json:"top"`
	CountryName string `json:"country_name"`
}

type sumCompaniesParam struct {
	CountryName string `json:"country_name"`
}

type oneCompany struct {
	Name string `json:"name" validate:"required"`
}

func (o oneCompany) Valid() error {
	return validator.New().Struct(o)
}
