package model

import "time"

// 新抽奖
const (
	NORMAL = iota
	JOKE
	HIGH
	HOMEPAGE
	LUCYATONCE
)

var LotteryClass = make(map[int]string)

func init() {
	LotteryClass[NORMAL] = "普通抽奖"
	LotteryClass[JOKE] = "皮一下"
	LotteryClass[HIGH] = "高级抽奖"
	LotteryClass[HOMEPAGE] = "自助上首页"
	LotteryClass[LUCYATONCE] = "现场抽奖"

}

type Lottery struct {
	Base            `xorm:"extends"`
	ImageURL        string    `json:"image_url"`
	Deadline        time.Time `json:"deadline"`
	Levels          []Level   `json:"levels"`
	WinnerLotteryId int64     `xorm:"index" json:"winner_lottery_id"`
	Class           int       `json:"class"`
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
	Level     int    `json:"level"`
	LotteryId int64  `xorm:"index" json:"lottery_id"`
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
	Name        string `json:"name"`
	Description string `json:"description"`
	Number      int    `json:"number"`
}
