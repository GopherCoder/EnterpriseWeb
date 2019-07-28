package error_for_lottery

import "fmt"

type ErrorForLottery struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Detail  string `json:"detail"`
}

func (E ErrorForLottery) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s, Detail: %s",
		E.Code, E.Message, E.Detail)
}
