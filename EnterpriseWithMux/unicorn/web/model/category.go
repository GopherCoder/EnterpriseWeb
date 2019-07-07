package model

type Category struct {
	base
	Name string `json:"name" gorm:"varchar(12)"`
}

func (c Category) TableName() string {
	return "category"
}
