package model

import (
	"fmt"
	"time"
)

// 用户信息
type Admin struct {
	Base     `xorm:"extends"`
	Phone    string `xorm:"unique" json:"phone"`
	Password string `json:"password"`
	Token    string `json:"token"`
	Name     string `xorm:"unique" json:"name"`
}

func (A Admin) TableName() string {
	return fmt.Sprintf("%s_%s", PROJECT, "admin")
}

type AdminSerializer struct {
	Id        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Phone     string    `json:"phone"`
	Token     string    `json:"token"`
	Name      string    `json:"name"`
}

func (A Admin) Serializer() AdminSerializer {
	return AdminSerializer{
		Id:        A.Id,
		CreatedAt: A.CreatedAt,
		UpdatedAt: A.UpdatedAt,
		Phone:     A.Phone,
		Token:     A.Token,
		Name:      A.Name,
	}
}
