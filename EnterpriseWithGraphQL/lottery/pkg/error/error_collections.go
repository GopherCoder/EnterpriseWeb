package error_for_lottery

import "net/http"

var (
	ErrorInsert = ErrorForLottery{Code: http.StatusBadRequest, Message: "Insert into db fail", Detail: "插入记录失败"}
)
