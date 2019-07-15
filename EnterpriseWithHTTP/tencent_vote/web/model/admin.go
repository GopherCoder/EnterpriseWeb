package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Admin struct {
	gorm.Model
	Phone    string `gorm:"type:varchar(11);unique;not null" json:"phone"`
	Password string `gorm:"type:varchar(128)" json:"password"`
	Token    string `gorm:"type:varchar(32)" json:"token"`
}

type AdminSerializer struct {
	Id        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Phone     string    `json:"phone"`
	Password  string    `json:"password"`
	Token     string    `json:"token"`
}

func (a Admin) Serializer() AdminSerializer {
	return AdminSerializer{
		Id:        a.ID,
		CreatedAt: a.CreatedAt.Truncate(time.Second),
		UpdatedAt: a.UpdatedAt.Truncate(time.Second),
		Phone:     a.Phone,
		Password:  a.Password,
		Token:     a.Token,
	}
}
