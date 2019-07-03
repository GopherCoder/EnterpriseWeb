package model

import (
	"EnterpriseWeb/EnterpriseWithEcho/target_notes/pkg/database"
	"time"
)

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

const (
	UNDONE = iota
	DONE
)

var MISSION_MAP = make(map[int]string)

func init() {
	TARGET_STATUS = map[int]string{
		0: "核心目标", // just one
		1: "目标完成",
		2: "设为FLAG", // EveryOne Can See
		3: "普通目标",
	}
	MISSION_MAP = map[int]string{
		0: "未完成",
		1: "已完成",
	}
}

type Target struct {
	base        `xorm:"extends"`
	AdminId     int64
	Title       string  `xorm:"varchar(12)" json:"title"`
	Description string  `xorm:"text" json:"description"`
	TaskIds     []int64 `xorm:"blob" json:"task_ids"`
	Status      int     `xorm:"varchar(12)"`
	Level       int     `xorm:"int(1)" json:"level"`
}

func (t Target) TableName() string {
	return "targetNotes_target"
}

type SmallSerializer struct {
	Id          int64     `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	AdminId     int64     `json:"admin_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	Level       string    `json:"level"`
}

type TargetSerializer struct {
	SmallSerializer
	TaskIds []int64 `json:"task_ids"`
}
type OneTask struct {
	Id    int64  `json:"id"`
	Title string `json:"title"`
}
type TargetSerializerWithTaskTitle struct {
	SmallSerializer
	Tasks []OneTask `json:"tasks"`
}

var status = func(s int) string {
	return MISSION_MAP[s]
}

var level = func(l int) string {
	return TARGET_STATUS[l]
}

func (t Target) Serializer() TargetSerializer {

	return TargetSerializer{
		SmallSerializer: SmallSerializer{
			Id:          t.Id,
			CreatedAt:   t.CreatedAt,
			AdminId:     t.AdminId,
			Title:       t.Title,
			Description: t.Description,

			Status: status(t.Status),
			Level:  level(t.Level),
		},
		TaskIds: t.TaskIds,
	}
}

func (t Target) SerializerWithTaskTitle() TargetSerializerWithTaskTitle {
	var tasks []Task
	database.Engine.In("id", t.TaskIds).Find(&tasks)
	var taskCollections []OneTask
	for _, i := range tasks {
		var one OneTask
		one.Id = i.Id
		one.Title = i.Title
		taskCollections = append(taskCollections, one)
	}
	return TargetSerializerWithTaskTitle{
		SmallSerializer: SmallSerializer{
			Id:          t.Id,
			CreatedAt:   t.CreatedAt,
			AdminId:     t.AdminId,
			Title:       t.Title,
			Description: t.Description,
			Status:      status(t.Status),
			Level:       level(t.Level),
		},
		Tasks: taskCollections,
	}
}
