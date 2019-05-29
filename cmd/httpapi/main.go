package main

import (
	"net/http"

	httpapilib "github.com/ryanchaiyakul/datagen/internal/http"
)

func main() {
	http.HandleFunc("/", httpapilib.Handler)
	if err := http.ListenAndServe(":80", nil); err != nil {
		return
	}
}
