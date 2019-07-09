package main

import (
	"net/http"

	httplib "github.com/ryanchaiyakul/datagen/internal/cmd/http"
)

func main() {
	http.HandleFunc("/", httplib.Handler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		return
	}
}
