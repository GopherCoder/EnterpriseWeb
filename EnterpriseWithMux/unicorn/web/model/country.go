package model

type Country struct {
	base
	Name string `json:"name" gorm:"varchar(12)"`
}

func (c Country) TableName() string {
	return "country"
}
