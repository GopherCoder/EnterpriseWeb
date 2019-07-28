package model

import (
	"EnterpriseWeb/EnterpriseWithGraphQL/lottery/pkg/database"
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

// 用户参与的抽奖项目
type AdminTakePart struct {
	Base       `xorm:"extends"`
	AdminId    int64   `xorm:"index" json:"admin_id"`
	LotteryIds []int64 `json:"lottery_ids"`
}

func (A AdminTakePart) TableName() string {
	return fmt.Sprintf("%s_%s", PROJECT, "admin_take_part")
}

type AdminTakePartSerializer struct {
	Id         int64     `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	AdminId    int64     `json:"admin_id"`
	AdminName  string    `json:"admin_name"`
	LotteryIds []int64   `json:"lottery_ids"`
}

func (A AdminTakePart) Serializer() AdminTakePartSerializer {
	var admin Admin
	database.Engine.ID(A.AdminId).Get(&admin)
	return AdminTakePartSerializer{
		Id:         A.Id,
		CreatedAt:  A.CreatedAt.Truncate(time.Second),
		UpdatedAt:  A.UpdatedAt.Truncate(time.Second),
		AdminId:    A.AdminId,
		AdminName:  admin.Name,
		LotteryIds: A.LotteryIds,
	}
}
