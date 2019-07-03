package target

import "gopkg.in/go-playground/validator.v9"

type createTaskParam struct {
	Title string `json:"title" validate:"required"`
}

func (c createTaskParam) Valid() error {
	return validator.New().Struct(c)
}

type patchTaskParam struct {
	Data struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Date        string `json:"date"`
		TargetId    int    `json:"target_id"`
		Status      int    `json:"status"`
	} `json:"data"`
}

func (p patchTaskParam) Valid() error {
	return validator.New().Struct(p)
}

type patchOrderParam struct {
	TaskIds []int `json:"task_ids" validate:"required"`
}

func (p patchOrderParam) Valid() error {
	return validator.New().Struct(p)
}
