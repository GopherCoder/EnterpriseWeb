package make_response

import (
	"encoding/json"
	"log"
	"net/http"
)

func Result(writer http.ResponseWriter, request *http.Request, code int, data interface{}) {
	var result = make(map[string]interface{})
	result["code"] = code
	if code == http.StatusOK {
		result["data"] = data
	} else {
		result["error"] = data
	}

	if err := json.NewEncoder(writer).Encode(result); err != nil {
		log.Println("err : ", err.Error())
	}

}
