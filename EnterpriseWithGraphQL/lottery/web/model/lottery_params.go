package model

import (
	"time"

	"gopkg.in/go-playground/validator.v9"
)

type LevelsParams struct {
	Name     string `json:"name" validate:"required"`
	ImageURL string `json:"image_url"`
	Number   int    `json:"number" validate:"min=1"`
	Class    int    `json:"class" validate:"minx=0,max=5"`
}

func (L LevelsParams) Valid() error {
	return validator.New().Struct(L)
}

type CreateLotteryParams struct {
	Levels                 []LevelsParams `json:"levels" validate:"min=1"`
	Deadline               time.Time      `json:"deadline"`
	LotteryClass           int            `json:"lottery_class" validate:"min=0,max=4"`
	Limit                  int            `json:"limit" validate:"min=0"`
	WinnerLotteryCondition int64          `json:"winner_lottery_condition" validate:"min=0,max=2"`
	AdminId                int64          `json:"admin_id" validate:"required"`
}

func (C CreateLotteryParams) Valid() error {
	return validator.New().Struct(C)
}
