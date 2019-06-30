package model

type Admin struct {
	base        `xorm:"extends"`
	AccountName string `xorm:"varchar(12)" json:"account_name"`
}
