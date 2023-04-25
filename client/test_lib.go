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
	return httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(returnCode)
		res.Write(b)
	}))
}
