package model

import "time"

type base struct {
	Id        int64
	CreatedAt time.Time `xorm:"created 'created_at'"`
	UpdatedAt time.Time `xorm:"updated 'updated_at'"`
	DeletedAt time.Time `xorm:"deleted 'deleted_at'"`
}
