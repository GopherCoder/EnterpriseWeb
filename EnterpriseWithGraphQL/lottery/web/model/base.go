package model

import "time"

type Base struct {
	Id        int64     `xorm:"pk notnull autoincr"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
	DeletedAt time.Time `xorm:"deleted"`
}
