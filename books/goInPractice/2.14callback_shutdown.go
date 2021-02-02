package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/shutdown", shutdown)
	http.HandleFunc("/", homePage)
	_ = http.ListenAndServe("localhost:8080", nil)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	_, _ = fmt.Fprint(w, "The Homepage")
}

func shutdown(_ http.ResponseWriter, _ *http.Request) {
	os.Exit(0)
}
