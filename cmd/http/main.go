package main

import (
	"net/http"

	cmdlib "github.com/ryanchaiyakul/datagen/internal/cmd"
)

func main() {
	http.HandleFunc("/", cmdlib.Handler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		return
	}
}
