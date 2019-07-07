package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Concept struct {
	gorm.Model
	Key    string `json:"key" gorm:"type:"varchar(16)"`
	Detail string `json:"detail"`
}

type ConceptSerializer struct {
	Id        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Key       string    `json:"key"`
	Detail    string    `json:"detail"`
}

func (c Concept) Serializer() ConceptSerializer {
	return ConceptSerializer{
		Id:        c.ID,
		CreatedAt: c.CreatedAt.Truncate(time.Second),
		UpdatedAt: c.UpdatedAt.Truncate(time.Second),
		Key:       c.Key,
		Detail:    c.Detail,
	}
}
