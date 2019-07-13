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
	ErrorMethod          = ErrorTenCentVotes{Code: http.StatusBadRequest, Message: "method is rejected", Detail: "方法不允许"}
	ErrorGetRecord       = ErrorTenCentVotes{Code: http.StatusBadRequest, Message: "record not found", Detail: "记录不存在"}
	ErrorPostRecord      = ErrorTenCentVotes{Code: http.StatusBadRequest, Message: "record create fail", Detail: "记录创建失败"}
	ErrorPatchRecord     = ErrorTenCentVotes{Code: http.StatusBadRequest, Message: "record update fail", Detail: "记录更新失败"}
	ErrorDeleteRecord    = ErrorTenCentVotes{Code: http.StatusBadRequest, Message: "record delete fail", Detail: "记录删除失败"}
)
