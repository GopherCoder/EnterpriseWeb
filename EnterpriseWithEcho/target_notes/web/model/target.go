package model

import "time"

const (
	// 核心目标、目标完成、设Flag、普通目标
	FOCUSON = iota
	COMPLETE
	FLAG
	NORMAL
)

var (
	TARGET_STATUS = make(map[int]string)
)

func init() {
	TARGET_STATUS = map[int]string{
		0: "核心目标", // just one
		1: "目标完成",
		2: "设为FLAG", // EveryOne Can See
		3: "普通目标",
	}
}

type Target struct {
	base        `xorm:"extends"`
	AdminId     int64
	Title       string  `xorm:"varchar(12)" json:"title"`
	Description string  `xorm:"text" json:"description"`
	TaskIds     []int64 `xorm:"blob" json:"task_ids"`
	Status      int     `xorm:"varchar(12)"`
}

func (t Target) TableName() string {
	return "targetNotes_target"
}

type TargetSerializer struct {
	Id          int64
	CreatedAt   time.Time `json:"created_at"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	TargetIds   []int64   `json:"target_ids"`
	Status      string    `json:"status"`
}

func (t Target) Serializer() TargetSerializer {
	return TargetSerializer{
		Id:          t.Id,
		CreatedAt:   t.CreatedAt,
		Title:       t.Title,
		Description: t.Description,
		TargetIds:   t.TaskIds,
		Status:      TARGET_STATUS[t.Status],
	}
}

type Target2Task struct {
	TargetId int64 `xorm:"index"`
	TaskId   int64 `xorm:"index"`
}

func (t Target2Task) TableName() string {
	return "targetNotes_target2task"
}
