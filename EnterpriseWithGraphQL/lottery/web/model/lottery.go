package model

import (
	"EnterpriseWeb/EnterpriseWithGraphQL/lottery/pkg/database"
	"fmt"
	"time"
)

// 新抽奖
const (
	NORMAL = iota
	JOKE
	HIGH
	HOMEPAGE
	LUCKY
)

var LotteryClass = map[int]string{}

func init() {
	LotteryClass = make(map[int]string)

	LotteryClass[NORMAL] = "普通抽奖"
	LotteryClass[JOKE] = "皮一下"
	LotteryClass[HIGH] = "高级抽奖"
	LotteryClass[HOMEPAGE] = "自助上首页"
	LotteryClass[LUCKY] = "现场抽奖"

}

type Lottery struct {
	Base            `xorm:"extends"`
	Deadline        time.Time `json:"deadline"`
	Levels          []int64   `json:"levels"`
	WinnerLotteryId int64     `xorm:"index 'winner_lottery_id'" json:"winner_lottery_id"`
	Class           int       `json:"class"`
	Number          int       `json:"number"`
	Limit           int       `json:"limit"`
	AdminID         int64     `xorm:"index 'admin_id'" json:"admin_id"`
}

func (L Lottery) TableName() string {
	return fmt.Sprintf("%s_%s", PROJECT, "lottery")
}

type LotterySerializer struct {
	Id                int64              `json:"id"`
	CreatedAt         time.Time          `json:"created_at"`
	UpdatedAt         time.Time          `json:"updated_at"`
	Deadline          string             `json:"deadline"`
	Levels            []*LevelSerializer `json:"levels"`
	WinnerLotteryName string             `json:"winner_lottery_name"`
	Class             int                `json:"class"`
	ClassString       string             `json:"class_string"`
	Number            int                `json:"number"`
	Limit             int                `json:"limit"`
	AdminName         string             `json:"admin_name"`
}

func (L Lottery) Serializer() *LotterySerializer {
	var levels []*LevelSerializer
	for _, i := range L.Levels {
		var one Level
		database.Engine.ID(i).Get(&one)
		levels = append(levels, one.Serializer())
	}
	var winnerLottery WinnerLottery
	if has, dbError := database.Engine.ID(L.WinnerLotteryId).Get(&winnerLottery); !has || dbError != nil {
		return nil
	}
	var admin Admin
	if has, dbError := database.Engine.ID(L.AdminID).Get(&admin); !has || dbError != nil {
		return nil
	}

	return &LotterySerializer{
		Id:                L.Id,
		CreatedAt:         L.CreatedAt.Truncate(time.Second),
		UpdatedAt:         L.UpdatedAt.Truncate(time.Second),
		Deadline:          L.Deadline.Format(time.RFC3339),
		Levels:            levels,
		WinnerLotteryName: WinnerCondition[winnerLottery.Class],
		ClassString:       LotteryClass[L.Class],
		Class:             L.Class,
		Number:            L.Number,
		Limit:             L.Limit,
		AdminName:         admin.Name,
	}
}

// 奖项
const (
	FIRST = iota
	SECOND
	THIRD
	FOURTH
	FIFTH
	SIXTH
)

var Prize = map[int]string{}

func init() {
	Prize = make(map[int]string)
	Prize[FIRST] = "一等奖"
	Prize[SECOND] = "二等奖"
	Prize[THIRD] = "三等奖"
	Prize[FOURTH] = "四等奖"
	Prize[FIFTH] = "五等奖"
	Prize[SIXTH] = "六等奖"
}

type Level struct {
	Base     `xorm:"extends"`
	Name     string `json:"name"`
	ImageURL string `xorm:"'image_url'"json:"image_url"`
	Number   int    `json:"number"`
	Class    int    `json:"class"`
}

func (L Level) TableName() string {
	return fmt.Sprintf("%s_%s", PROJECT, "level")
}

type LevelSerializer struct {
	Id          int64     `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	ImageURL    string    `json:"image_url"`
	Name        string    `json:"name"`
	Number      int       `json:"number"`
	Class       int       `json:"class"`
	ClassString string    `json:"class_string"`
}

func (L Level) Serializer() *LevelSerializer {
	return &LevelSerializer{
		Id:          L.Id,
		CreatedAt:   L.CreatedAt,
		UpdatedAt:   L.UpdatedAt,
		ImageURL:    L.ImageURL,
		Name:        L.Name,
		Number:      L.Number,
		Class:       L.Class,
		ClassString: Prize[L.Class],
	}
}

// 开奖条件

const (
	TIMELEVEL = iota
	PERSONLEVEL
	NOWLEVEL
)

var WinnerCondition = map[int]string{}
var DefaultWinnerLottery = []WinnerLottery{}

func init() {
	WinnerCondition = make(map[int]string)
	WinnerCondition[TIMELEVEL] = "按时间自动开奖"
	WinnerCondition[PERSONLEVEL] = "按人数自动开奖"
	WinnerCondition[NOWLEVEL] = "即开即中"
	DefaultWinnerLottery = []WinnerLottery{
		{
			Class:       TIMELEVEL,
			Description: WinnerCondition[TIMELEVEL],
		}, {
			Class:       PERSONLEVEL,
			Description: WinnerCondition[PERSONLEVEL],
		}, {
			Class:       NOWLEVEL,
			Description: WinnerCondition[NOWLEVEL],
		},
	}
}

type WinnerLottery struct {
	Base        `xorm:"extends"`
	Description string `json:"description"`
	Class       int    `json:"index"`
}

func (W WinnerLottery) TableName() string {
	return fmt.Sprintf("%s_%s", PROJECT, "winner_lottery")
}

type WinnerLotterySerializer struct {
	Id          int64     `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Description string    `json:"description"`
	Class       int       `json:"class"`
	ClassString string    `json:"class_string"`
}

func (W WinnerLottery) Serializer() *WinnerLotterySerializer {
	return &WinnerLotterySerializer{
		Id:          W.Id,
		CreatedAt:   W.CreatedAt,
		UpdatedAt:   W.UpdatedAt,
		Description: W.Description,
		Class:       W.Class,
		ClassString: WinnerCondition[W.Class],
	}
}
