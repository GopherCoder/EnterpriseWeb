package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Vote struct {
	gorm.Model
	Title       string `json:"title" gorm:"type:varchar(32)"`
	AdminId     uint   `json:"admin_id"`
	Description string `json:"description" gorm:"type:varchar(64)"`
	Choice      []Choice
	DeadLine    time.Time
	IsAnonymous bool
	IsSingle    bool
}
type VoteSerializer struct {
	Id          uint      `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string    `json:"title"`
	AdminId     uint      `json:"admin_id"`
	Description string    `json:"description"`
	Choice      []string  `json:"choice"`
	DeadLine    time.Time `json:"dead_line"`
	IsAnonymous bool      `json:"is_anonymous"`
	IsSingle    bool      `json:"is_single"`
}

func (v Vote) Serializer(tx *gorm.DB) VoteSerializer {

	choices := func(i uint) []string {
		var choices []Choice
		tx.Where("id = ?", i).Find(&choices)
		var result []string
		for _, i := range choices {
			result = append(result, i.Title)
		}
		return result
	}

	return VoteSerializer{
		Id:          v.ID,
		Title:       v.Title,
		CreatedAt:   v.CreatedAt.Truncate(time.Second),
		UpdatedAt:   v.UpdatedAt.Truncate(time.Second),
		AdminId:     v.AdminId,
		Description: v.Description,
		Choice:      choices(v.ID),
		DeadLine:    v.DeadLine.Truncate(time.Second),
		IsAnonymous: v.IsAnonymous,
		IsSingle:    v.IsSingle,
	}
}

type Choice struct {
	gorm.Model
	VoteId uint
	Title  string `gorm:"type:varchar(32)" json:"title"`
}

type ChoiceSerializer struct {
	Id          uint      `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	VoteTitle   string    `json:"vote_title"`
	ChoiceTitle string    `json:"choice_title"`
}

func (c Choice) Serializer(tx *gorm.DB) ChoiceSerializer {
	voteTile := func(i uint) string {
		var vote Vote
		tx.Where("id = ?", i).First(&vote)
		return vote.Title
	}
	return ChoiceSerializer{
		Id:          c.ID,
		CreatedAt:   c.CreatedAt.Truncate(time.Second),
		UpdatedAt:   c.UpdatedAt.Truncate(time.Second),
		VoteTitle:   voteTile(c.VoteId),
		ChoiceTitle: c.Title,
	}

}
