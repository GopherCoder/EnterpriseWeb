package model

import (
	"EnterpriseWeb/EnterpriseWithEcho/target_notes/pkg/database"
	"log"
	"time"
)

type Task struct {
	base        `xorm:"extends"`
	TargetId    int64
	Description string  `xorm:"varchar(32)" json:"description"`
	Title       string  `xorm:"varchar(32)" json:"title"`
	ThingIds    []int64 `xorm:"blob" json:"thing_ids"`
	Status      int     `xorm:"int(1)" json:"status"`
	OrderLevel  int     `xorm:"int(2)" json:"order_level"`
}

func (t Task) TableName() string {
	return "targetNotes_tasks"
}

type TaskSerializer struct {
	Id          int64              `json:"id"`
	CreatedAt   time.Time          `json:"created_at"`
	TargetName  string             `json:"target_name"`
	Description string             `json:"description"`
	Title       string             `json:"title"`
	Things      []ThingsSerializer `json:"things"`
	Status      string             `json:"status"`
}

func (t Task) Serializer() TaskSerializer {

	var target Target
	database.Engine.ID(t.TargetId).Get(&target)

	log.Println(target.TaskIds, "task_ids")

	var things []Things
	database.Engine.In("id", t.ThingIds).Find(&things)

	log.Println(t.ThingIds)

	var th []ThingsSerializer
	for _, i := range things {
		th = append(th, i.Serializer())
	}
	return TaskSerializer{
		Id:          t.Id,
		CreatedAt:   t.CreatedAt,
		TargetName:  target.Title,
		Description: t.Description,
		Title:       t.Title,
		Things:      th,
		Status:      status(t.Status),
	}
}

type Things struct {
	base        `xorm:"extends"`
	Description string `xorm:"text" json:"description"`
}

func (t Things) TableName() string {
	return "targetNotes_things"
}

type ThingsSerializer struct {
	Id          int64     `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	Description string    `json:"description"`
}

func (t Things) Serializer() ThingsSerializer {
	return ThingsSerializer{
		Id:          t.Id,
		CreatedAt:   t.CreatedAt,
		Description: t.Description,
	}
}
