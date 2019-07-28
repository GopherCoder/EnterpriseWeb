package model

import (
	"EnterpriseWeb/EnterpriseWithGraphQL/lottery/pkg/database"
	"fmt"
	"time"
)

type Address struct {
	Base    `xorm:"extends"`
	AdminId int64  `xorm:"index" json:"admin_id"`
	Detail  string `json:"detail"`
}

func (A Address) TableName() string {
	return fmt.Sprintf("%s_%s", PROJECT, "addresses")
}

type AddressSerialize struct {
	Id        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	AdminId   int64     `json:"admin_id"`
	AdminName string    `json:"admin_name"`
	Detail    string    `json:"detail"`
}

func (A Address) Serializer() AddressSerialize {
	var admin Admin
	database.Engine.ID(A.AdminId).Get(&admin)
	return AddressSerialize{
		Id:        A.Id,
		CreatedAt: A.CreatedAt.Truncate(time.Second),
		UpdatedAt: A.UpdatedAt.Truncate(time.Second),
		AdminId:   A.AdminId,
		Detail:    A.Detail,
		AdminName: admin.Name,
	}
}
