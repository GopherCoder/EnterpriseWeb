package vote

import (
	"fmt"
	"time"

	"gopkg.in/go-playground/validator.v9"
)

type CreateParam struct {
	Title       string   `json:"title" validate:"required"`
	Description string   `json:"description"`
	Choice      []string `json:"choice" validate:"min=2"`
	Date        string   `json:"date"`
	IsAnonymous bool     `json:"is_anonymous"`
	IsSingle    bool     `json:"is_single" validate:"required"`
}

func (c CreateParam) Valid() error {
	return validator.New().Struct(c)
}

func (c CreateParam) toTime() (time.Time, error) {
	if c.Date == "" {
		now := time.Now()
		return now.AddDate(1, 0, 0), nil // 默认一年后
	}
	date, err := time.ParseInLocation(time.RFC3339, c.Date, time.Local)
	if err != nil {
		return time.Time{}, fmt.Errorf("format time error : %s", err.Error())
	}
	return date, nil
}

type PatchParam struct {
	ChoiceIds []uint `json:"choice_ids" validate:"min=1"`
}

func (p PatchParam) Valid() error {
	return validator.New().Struct(p)
}
