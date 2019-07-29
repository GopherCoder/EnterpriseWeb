package model

import "fmt"

type Admin2Lottery struct {
	AdminID   int64 `xorm:"index 'admin_id'" json:"admin_id"`
	LotteryID int64 `xorm:"index 'lottery_id'" json:"lottery_id"`
}

func (A Admin2Lottery) TableName() string {
	return fmt.Sprintf("%s_%s", PROJECT, "admin2lottery")
}
