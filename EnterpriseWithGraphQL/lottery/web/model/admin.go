package model

import (
	"fmt"
)

// 用户信息
type Admin struct {
	Base     `xorm:"extends"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Token    string `json:"token"`
	Name     string `json:"name"`
}

func (A Admin) TableName() string {
	return fmt.Sprintf("%s_%s", PROJECT, "admin")
}

// 用户参与的抽奖项目
type AdminTakePart struct {
	Base       `xorm:"extends"`
	AdminId    int64   `xorm:"index" json:"admin_id"`
	LotteryIds []int64 `json:"lottery_ids"`
}

func (A AdminTakePart) TableName() string {
	return fmt.Sprintf("%s_%s", PROJECT, "admin_take_part")
}
