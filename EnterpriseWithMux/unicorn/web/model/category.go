package model

import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model
	Name string `json:"name" gorm:"varchar(12)"`
}

//func (c Category) TableName() string {
//	return "category"
//}
