package make_result

import (
	"encoding/json"
	"log"
	"net/http"
)

func Result(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	var result = make(map[string]interface{})
	result["code"] = code
	if code == http.StatusOK {
		result["data"] = data
	} else {
		result["error"] = data
	}
	err := json.NewEncoder(w).Encode(result)
	if err != nil {
		log.Println(err)
	}
}
