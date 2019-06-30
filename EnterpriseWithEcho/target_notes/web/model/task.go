package model

import "time"

type Task struct {
	base         `xorm:"extends"`
	AdminId      int64
	MissionIds   []int64 `xorm:"blob" json:"mission"`              // 任务
	Acquaintance string  `xorm:"varchar(255)" json:"acquaintance"` // 感悟
}

func (t Task) TableName() string {
	return "targetNotes_tasks"
}

type TaskSerializer struct {
	Id           int64        `json:"id"`
	CreatedAt    time.Time    `json:"created_at"`
	WeekDay      time.Weekday `json:"weekday"`
	MissionIds   []int64      `json:"mission_ids"`
	Acquaintance string       `json:"acquaintance"`
}

func (t TaskSerializer) Serializer() TaskSerializer {
	return TaskSerializer{
		Id:           t.Id,
		CreatedAt:    t.CreatedAt,
		WeekDay:      t.CreatedAt.Weekday(),
		MissionIds:   t.MissionIds,
		Acquaintance: t.Acquaintance,
	}
}
