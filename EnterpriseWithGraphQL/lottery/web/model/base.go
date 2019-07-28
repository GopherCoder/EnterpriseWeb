package model

import "time"

/*
	基础字段
*/

type Base struct {
	Id        int64      `xorm:"pk notnull autoincr" json:"id"`
	CreatedAt time.Time  `xorm:"created" json:"created_at"`
	UpdatedAt time.Time  `xorm:"updated" json:"updated_at"`
	DeletedAt *time.Time `xorm:"deleted" json:"deleted_at"`
}
