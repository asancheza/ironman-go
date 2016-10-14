package main

import (
	"fmt"
	"net/http"
)

func example_controller(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi example")
}
