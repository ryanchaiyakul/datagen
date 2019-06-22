package main

import (
	"net/http"

	httpmainlib "github.com/ryanchaiyakul/datagen/internal/cmd/http"
)

func main() {
	http.HandleFunc("/", httpmainlib.Handler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		return
	}
}
