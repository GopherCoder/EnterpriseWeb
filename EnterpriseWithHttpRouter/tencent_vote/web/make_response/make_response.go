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
	writer.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(writer)
	enc.SetIndent("", "")
	err := enc.Encode(result)
	if err != nil {
		log.Println(err)
	}
}
