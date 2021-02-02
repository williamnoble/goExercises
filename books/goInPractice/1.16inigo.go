package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloing)
	http.ListenAndServe("localhost:8080", nil)
}

func helloing(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Hello Earthlings, My Name is William")
}
