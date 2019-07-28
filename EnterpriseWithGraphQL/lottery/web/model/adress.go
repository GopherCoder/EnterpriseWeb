package model

import (
	"EnterpriseWeb/EnterpriseWithGraphQL/lottery/pkg/database"
	"fmt"
)

type Address struct {
	Base    `xorm:"extends"`
	AdminID int64  `xorm:"index" json:"admin_id"`
	Detail  string `json:"detail"`
}

func (A Address) TableName() string {
	return fmt.Sprintf("%s_%s", PROJECT, "adresses")
}

func GetAddresses(adminID int64) ([]*Address, error) {
	var results []*Address
	if dbError := database.Engine.ID(adminID).Find(&results); dbError != nil {
		return results, dbError
	}
	return results, nil
}
