package httplib

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Handler is the HTTP Rest API function to be passed into net/http
func Handler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "405 method not allowed.", http.StatusMethodNotAllowed)
		return
	}

	ret := getData(r)
	switch ret.(type) {
	case error:
		w.Write([]byte(fmt.Sprint(ret)))
	default:
		if retEncoded, err := json.Marshal(ret); err == nil {
			w.Write(retEncoded)
		} else {
			w.Write([]byte(fmt.Sprintf("Handler : failed to parse : %v", ret)))
		}
	}
}
