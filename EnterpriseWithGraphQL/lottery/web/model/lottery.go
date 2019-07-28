package model

import (
	"EnterpriseWeb/EnterpriseWithGraphQL/lottery/pkg/database"
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

var LotteryClass = make(map[int]string)

func init() {
	LotteryClass[NORMAL] = "普通抽奖"
	LotteryClass[JOKE] = "皮一下"
	LotteryClass[HIGH] = "高级抽奖"
	LotteryClass[HOMEPAGE] = "自助上首页"
	LotteryClass[LUCKY] = "现场抽奖"

}

type Lottery struct {
	Base            `xorm:"extends"`
	ImageURL        string    `json:"image_url"`
	Deadline        time.Time `json:"deadline"`
	Levels          []Level   `json:"levels"`
	WinnerLotteryId int64     `xorm:"index" json:"winner_lottery_id"`
	Class           int       `json:"class"`
	Number          int       `json:"number"`
	Limit           int       `json:"limit"`
	AdminID         int64     `xorm:"index" json:"admin_id"`
}

type LotterySerialize struct {
	Id                int64            `json:"id"`
	CreatedAt         time.Time        `json:"created_at"`
	UpdatedAt         time.Time        `json:"updated_at"`
	ImageURL          string           `json:"image_url"`
	DeadLine          string           `json:"dead_line"`
	Levels            []LevelSerialize `json:"levels"`
	WinnerLotteryName string           `json:"winner_lottery_name"`
	Class             string           `json:"class"`
	Number            int              `json:"number"`
	Limit             int              `json:"limit"`
	AdminName         string           `json:"admin_name"`
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

var Prize = make(map[int]string)

func init() {
	Prize[FIFTH] = "一等奖"
	Prize[SECOND] = "二等奖"
	Prize[THIRD] = "三等奖"
	Prize[FOURTH] = "四等奖"
	Prize[FIFTH] = "五等奖"
	Prize[SIXTH] = "六等奖"
}

type Level struct {
	Base      `xorm:"extends"`
	Name      string `json:"name"`
	Number    int    `json:"number"`
	Class     int    `json:"class"`
	LotteryId int64  `xorm:"index" json:"lottery_id"`
}

type LevelSerialize struct {
	Id        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Number    int       `json:"number"`
	Class     int       `json:"class"`
	LotteryId int64     `json:"lottery_id"`
}

// 开奖条件

const (
	TIMELEVEL = iota
	PERSONLEVEL
	NOWLEVEL
)

var WinnerCondition = make(map[int]string)

func init() {
	WinnerCondition[TIMELEVEL] = "按时间自动开奖"
	WinnerCondition[PERSONLEVEL] = "按人数自动开奖"
	WinnerCondition[NOWLEVEL] = "即开即中"

}

type WinnerLottery struct {
	Base        `xorm:"extends"`
	Description string `json:"description"`
	Class       int    `json:"index"`
}

func ListLottery(ownerID int64) ([]Lottery, error) {
	var results []Lottery

	if dbError := database.Engine.ID(ownerID).Find(&results); dbError != nil {
		return results, dbError
	}
	return results, nil
}

func OneLottery(id int64) (Lottery, error) {
	var result Lottery
	if has, err := database.Engine.ID(id).Get(&result); !has || err != nil {
		return result, err
	}
	return result, nil
}

func InvolvementsLottery(adminID int64) ([]Lottery, error) {
	var results []Lottery
	var adminTakePartIn AdminTakePart
	if has, dbError := database.Engine.ID(adminID).Get(&adminTakePartIn); dbError != nil || !has {
		return results, dbError
	}

	if dbError := database.Engine.In("lottery_ids", adminTakePartIn.LotteryIds).Find(&results); dbError != nil {
		return results, dbError
	}
	return results, nil
}
