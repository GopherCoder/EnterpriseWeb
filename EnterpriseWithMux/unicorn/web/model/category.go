package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Category struct {
	gorm.Model
	Name string `json:"name" gorm:"varchar(12)"`
}

//func (c Category) TableName() string {
//	return "category"
//}

type CategorySerializer struct {
	Id        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
}

func (c Category) Serializer() CategorySerializer {
	return CategorySerializer{
		Id:        c.ID,
		CreatedAt: c.CreatedAt.Truncate(time.Second),
		UpdatedAt: c.UpdatedAt.Truncate(time.Second),
		Name:      c.Name,
	}
}
