package make_request

import (
	"encoding/json"
	"net/http"
)

func BindJson(request *http.Request, param interface{}) error {
	if err := json.NewDecoder(request.Body).Decode(&param); err != nil {
		return err
	}
	return nil
}

func Query(r *http.Request, key string) string {
	return r.URL.Query().Get(key)
}

func QueryAndDefault(r *http.Request, key string, de string) string {
	value := r.URL.Query().Get(key)
	if value == "" {
		return de
	}
	return value
}

func Vars(r *http.Request) map[string]interface{} {
	var results = make(map[string]interface{})
	for key, i := range r.URL.Query() {
		results[key] = i[0]
	}
	return results
}
