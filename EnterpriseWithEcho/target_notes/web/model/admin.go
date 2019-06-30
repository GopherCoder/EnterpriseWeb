package model

import "time"

type Admin struct {
	base        `xorm:"extends"`
	AccountName string `xorm:"varchar(12)" json:"account_name"`
	Phone       string `xorm:"varchar(11)" json:"phone"`
	Token       string `xorm:"text" json:"token"`
}

func (a Admin) TableName() string {
	return "targetNotes_admins"
}

type AdminSerializer struct {
	Id          int64     `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	AccountName string    `json:"account_name"`
	Phone       string    `json:"phone"`
	Token       string    `json:"token"`
}

func (a Admin) Serializer() AdminSerializer {
	return AdminSerializer{
		Id:          a.Id,
		CreatedAt:   a.CreatedAt,
		AccountName: a.AccountName,
		Phone:       a.Phone,
		Token:       a.Token,
	}
}
