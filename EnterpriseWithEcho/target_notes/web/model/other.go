package model

type Other struct {
	base    `xorm:"extends"`
	AdminId uint
	TaskIds []int64 `xorm:"blob" json:"task_ids"`
}

func (o Other) TableName() string {
	return "targetNotes_other"
}

type oneTask struct {
	Id    int64  `json:"id"`
	Title string `json:"title"`
}
type OtherSerializer struct {
	Id   int64     `json:"id"`
	Name string    `json:"name"`
	Task []oneTask `json:"task"`
}

func (o Other) SerializerWithOut() OtherSerializer {
	return OtherSerializer{}

}
