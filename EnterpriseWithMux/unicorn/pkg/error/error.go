package error_for_unicorn

import (
	"fmt"
	"net/http"
)

type ErrorFor struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Detail  string `json:"detail"`
}

func (e ErrorFor) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s, Detail: %s", e.Code, e.Message, e.Detail)
}

var (
	RecordError = ErrorFor{Code: http.StatusBadRequest, Message: "Record not found", Detail: "记录未找到"}
)
