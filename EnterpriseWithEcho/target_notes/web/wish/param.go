package wish

import "gopkg.in/go-playground/validator.v9"

type postWishPara struct {
	Title string `json:"title" validate:"required"`
}

func (p postWishPara) Valid() error {
	return validator.New().Struct(p)
}
