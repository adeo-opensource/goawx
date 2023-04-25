package awx

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
)

func GetServer(returnCode int, obj map[string]interface{}) *httptest.Server {
	b, err := json.Marshal(obj)
	if err != nil {
		return nil
	}

	handler := func(res http.ResponseWriter, r *http.Request) {
		res.WriteHeader(returnCode)
		_, err := res.Write(b)
		if err != nil {
			return
		}
	}

	return httptest.NewServer(http.HandlerFunc(handler))
}
