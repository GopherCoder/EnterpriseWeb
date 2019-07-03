package wish

import (
	"EnterpriseWeb/EnterpriseWithEcho/target_notes/web/param"

	"gopkg.in/go-playground/validator.v9"
)

type postWishParam struct {
	Title string `json:"title" validate:"required"`
}

func (p postWishParam) Valid() error {
	return validator.New().Struct(p)
}

type patchWishParam struct {
	Data struct {
		Title          string `json:"title"`
		Hope           string `json:"hope"`
		TargetId       int    `json:"target_id" validate:"min=0"`
		DesireLevel    int    `json:"desire_level" validate:"min=0,max=9"`
		ChallengeLevel int    `json:"challenge_level" validate:"min=0,max=9"`
		TimeLevel      int    `json:"time_level" validate:"min=0,max=9"`
	} `json:"data" validate:"required"`
}

func (p patchWishParam) Valid() error {
	return validator.New().Struct(p)
}

type getWishParam struct {
	param.ReturnParam
}

func (g getWishParam) Valid() error {
	return validator.New().Struct(g)
}
