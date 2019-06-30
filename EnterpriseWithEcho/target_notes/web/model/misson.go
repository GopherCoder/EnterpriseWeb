package model

import "time"

const (
	UNDONE = iota
	DONE
)

var MISSION_MAP = make(map[int]string)

func init() {
	MISSION_MAP = map[int]string{
		0: "未完成",
		1: "已完成",
	}
}

type Mission struct {
	base        `xorm:"extends"`
	TargetId    int64
	Description string   `xorm:"text" json:"description"`
	Things      []Things `xorm:"text" json:"things"`
	Status      int      `xorm:"int(1)" json:"status"`
}

type Things struct {
	base        `xorm:"extends"`
	Title       string `xorm:"varchar(12)" json:"title"`
	Description string `xorm:"text" json:"description"`
}

func (m Mission) TableName() string {
	return "targetNotes_missions"
}

type MissionSerializer struct {
	Id          int64     `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	TargetId    int64     `json:"target_id"`
	Description string    `json:"description"`
	Things      []string  `json:"things"`
	Status      string    `json:"status"`
}

func (m Mission) Serializer() MissionSerializer {
	things := func(m Mission) []string {
		var results []string
		for _, i := range m.Things {
			results = append(results, i.Title)
		}
		return results
	}
	return MissionSerializer{
		Id:          m.Id,
		CreatedAt:   m.CreatedAt,
		TargetId:    m.TargetId,
		Description: m.Description,
		Things:      things(m),
		Status:      MISSION_MAP[m.Status],
	}
}
