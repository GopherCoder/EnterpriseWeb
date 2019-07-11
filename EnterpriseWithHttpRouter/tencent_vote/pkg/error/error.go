package error_tencent_votes

import (
	"fmt"
	"net/http"
)

type ErrorTenCentVotes struct {
	Code    int
	Message string
	Detail  string
}

func (e ErrorTenCentVotes) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s, Detail: %s", e.Code, e.Message, e.Detail)
}

var (
	ErrorConnectDatabase = ErrorTenCentVotes{Code: http.StatusBadRequest, Message: "connect database fail", Detail: "连接数据库失败"}
)
