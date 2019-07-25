package model

import "fmt"

type Admin struct {
	Base     `xorm:"extends"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

func (A Admin) TableName() string {
	return fmt.Sprintf("%s_%s", PROJECT, "admin")
}

type AdminTakePart struct {
	Base       `xorm:"extends"`
	AdminId    int64   `xorm:"index" json:"admin_id"`
	LotteryIds []int64 `json:"lottery_id"`
}
