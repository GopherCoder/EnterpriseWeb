package model

const (
	// 一般渴望、比较渴望、很渴望、非常渴望
	DESIRE_ZERO = iota
	DESIRE_ONE
	DESIRE_TWO
	DESIRE_THREE
)

const (
	// 一般挑战、比较挑战、很挑战、非常挑战
	CHALLENGE_ZERO = iota
	CHALLENGE_ONE
	CHALLENGE_TWO
	CHALLENGE_THREE
)

const (
	// 一个月内、一年内、三年内、三年以上
	TIME_ONE_MONTH = iota
	TIME_ONE_YEAR
	TIME_THREE_YEAR
	TIME_OVER_THREE_YEAR
)

var (
	DESIRE_MAP    = make(map[int]string)
	CHALLENGE_MAP = make(map[int]string)
	TIME_MAP      = make(map[int]string)
)

func init() {
	DESIRE_MAP = map[int]string{
		0: "一般渴望",
		1: "比较渴望",
		2: "很渴望",
		3: "非常渴望",
	}
	CHALLENGE_MAP = map[int]string{
		0: "一般挑战",
		1: "比较挑战",
		2: "很挑战",
		3: "非常挑战",
	}
	TIME_MAP = map[int]string{
		0: "一个月内",
		1: "一年内",
		2: "三年内",
		3: "三年以上",
	}
}

type Wish struct {
	base           `xorm:"extends"`
	AdminId        int64
	Title          string `xorm:"varchar(24)" json:"title"`
	Hope           string `json:"hope"` // 提示语：这样就能
	DesireLevel    int    `xorm:"int(2)" json:"level_desire"`
	ChallengeLevel int    `xorm:"int(2)" json:"challenge_level"`
	TimeLevel      int    `xorm:"int(2)" json:"time_level"`
	TargetId       int64
}

func (w Wish) TableName() string {
	return "targetNotes_wish"
}

type WishSerializer struct {
	Id             int64  `json:"id"`
	AdminID        int64  `json:"admin_id"`
	Title          string `json:"title"`
	Hope           string `json:"hope"`
	DesireLevel    string `json:"desire_level"`
	ChallengeLevel string `json:"challenge_level"`
	TimeLevel      string `json:"time_level"`
	TargetId       int64  `json:"target_id"`
}

type WishSerializers []WishSerializer

var getLevel = func(level int) int {
	var resultLevel int
	if level <= 2 {
		resultLevel = 0
	} else if level > 2 && level <= 5 {
		resultLevel = 1
	} else if level > 5 && level <= 7 {
		resultLevel = 2
	} else {
		resultLevel = 3
	}
	return resultLevel
}

var getHope = func(hope string) string {
	if hope == "" {
		return "这样就能..."
	}
	return hope
}

func (w Wish) Serializer() WishSerializer {
	return WishSerializer{
		Id:             w.Id,
		AdminID:        w.AdminId,
		Title:          w.Title,
		Hope:           getHope(w.Hope),
		DesireLevel:    DESIRE_MAP[getLevel(w.DesireLevel)],
		ChallengeLevel: CHALLENGE_MAP[getLevel(w.ChallengeLevel)],
		TimeLevel:      TIME_MAP[getLevel(w.TimeLevel)],
		TargetId:       w.TargetId,
	}
}
