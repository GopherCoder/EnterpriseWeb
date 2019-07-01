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
	ParamErrorTarget  = ErrorTarget{Code: 400, Message: "param bind fail", Detail: "请求参数错误"}
	ValidErrorTarget  = ErrorTarget{Code: 400, Message: "param valid fail", Detail: "参数无效"}
	InsertErrorTarget = ErrorTarget{Code: 400, Message: "insert one record fail", Detail: "数据库插入数据错误"}
)
