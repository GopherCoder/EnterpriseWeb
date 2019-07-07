package country

import "gopkg.in/go-playground/validator.v9"

type CreateCountry struct {
	Data struct {
		Name string `json:"name" validate:"required"`
	} `json:"data" validate:"required"`
}

func (c CreateCountry) Valid() error {
	return validator.New().Struct(c)
}

type U struct {
	User     string `json:"user"`
	Password string `json:"password"`
}
