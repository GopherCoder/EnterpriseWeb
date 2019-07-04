package error_target_notes

import "fmt"

type ErrorTarget struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Detail  string `json:"detail"`
	Report  string `json:"report"`
}

func (e ErrorTarget) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s, Detail: %s, Report: %s", e.Code, e.Message, e.Detail, e.Report)
}

var (
	ParamErrorTarget    = ErrorTarget{Code: 400, Message: "param bind fail", Detail: "请求参数错误"}
	ValidErrorTarget    = ErrorTarget{Code: 400, Message: "param valid fail", Detail: "参数无效"}
	InsertErrorTarget   = ErrorTarget{Code: 400, Message: "insert one record fail", Detail: "数据库插入数据错误"}
	PassWordErrorTarget = ErrorTarget{Code: 400, Message: "password is not correct", Detail: "密码不匹配"}
	RecordErrorTarget   = ErrorTarget{Code: 400, Message: "record not found", Detail: "记录不存在"}
	UpdateErrorTarget   = ErrorTarget{Code: 400, Message: "update field fail", Detail: "更新数据失败"}
	DeleteErrorTarget   = ErrorTarget{Code: 400, Message: "delete fail", Detail: "删除记录失败"}
	RightErrorTarget    = ErrorTarget{Code: 400, Message: "no right of this router", Detail: "权限不允许"}
)
