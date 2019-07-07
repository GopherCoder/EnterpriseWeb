package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Country struct {
	gorm.Model
	Name string `json:"name" gorm:"varchar(12)"`
}

//func (c Country) TableName() string {
//	return "country"
//}

type CountrySerializer struct {
	Id        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
}

func (c Country) Serializer() CountrySerializer {
	return CountrySerializer{
		Id:        c.ID,
		CreatedAt: c.CreatedAt.Truncate(time.Second),
		UpdatedAt: c.UpdatedAt.Truncate(time.Second),
		Name:      c.Name,
	}
}
