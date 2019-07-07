package model

import "github.com/jinzhu/gorm"

type Country struct {
	gorm.Model
	Name string `json:"name" gorm:"varchar(12)"`
}

//func (c Country) TableName() string {
//	return "country"
//}
